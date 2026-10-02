package curator

import (
	"context"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/ygpark2/njro/agent/client"
	"github.com/ygpark2/njro/agent/orchestrator"
	searchpb "github.com/ygpark2/njro/service/search/proto"
)

// Reference represents a cited post or document used to build the answer.
type Reference struct {
	DocID   string `json:"doc_id"`
	Snippet string `json:"snippet"`
}

// CuratorResponse represents the synthesized knowledge answer from the community.
type CuratorResponse struct {
	Query       string      `json:"query"`
	Answer      string      `json:"answer"`
	References  []Reference `json:"references"`
	SourceCount int         `json:"source_count"`
}

// CuratorAgent discovers knowledge across community services and generates synthesized answers (RAG).
type CuratorAgent struct {
	llm     orchestrator.LLMClient
	clients *client.BackendClients
}

// NewCuratorAgent creates a new CuratorAgent instance.
func NewCuratorAgent(llm orchestrator.LLMClient, clients *client.BackendClients) *CuratorAgent {
	return &CuratorAgent{
		llm:     llm,
		clients: clients,
	}
}

// Curate queries the search service, gathers relevant community posts, and synthesizes a structured answer.
func (c *CuratorAgent) Curate(ctx context.Context, query string) (*CuratorResponse, error) {
	log.Info().Str("query", query).Msg("CuratorAgent: processing curation request")

	refs := make([]Reference, 0)
	var contextBuilder strings.Builder

	// 1. Search community via SearchService (Bleve index)
	if c.clients != nil && c.clients.Search != nil {
		searchResp, err := c.clients.Search.Search(ctx, &searchpb.SearchRequest{
			Keyword: query,
		})
		if err != nil {
			log.Warn().Err(err).Msg("CuratorAgent: search service call failed, proceeding with fallback")
		} else if searchResp != nil && len(searchResp.Documents) > 0 {
			for _, doc := range searchResp.Documents {
				snippet := doc.Text
				if len(snippet) > 200 {
					snippet = snippet[:200] + "..."
				}
				refs = append(refs, Reference{
					DocID:   doc.Id,
					Snippet: snippet,
				})
				contextBuilder.WriteString(fmt.Sprintf("- [문서 ID: %s]: %s\n", doc.Id, doc.Text))
			}
		}
	}

	// 2. Synthesize with LLM
	var answer string
	if c.llm != nil {
		var prompt string
		if contextBuilder.Len() > 0 {
			prompt = fmt.Sprintf(
				"당신은 커뮤니티 지식 큐레이션 AI 어시스턴트입니다.\n아래 수집된 커뮤니티 게시글 내용들을 기반으로 사용자의 질문에 대해 명확하고 일목요연하게 핵심을 요약하여 답변해 주세요. 관련 문서 번호가 있다면 함께 인용해 주세요.\n\n[수집된 커뮤니티 문서]\n%s\n\n[사용자 질문]\n%s",
				contextBuilder.String(), query,
			)
		} else {
			prompt = fmt.Sprintf(
				"당신은 커뮤니티 지식 큐레이션 AI 어시스턴트입니다. 커뮤니티에 직접 관련된 저장 문서는 없으나, 사용자의 질문에 대해 전문가 관점에서 친절하게 핵심 지식을 설명해 주세요.\n\n[사용자 질문]\n%s",
				query,
			)
		}

		resp, err := c.llm.Chat(ctx, []orchestrator.Message{
			{Role: "user", Content: prompt},
		}, nil)
		if err != nil {
			return nil, fmt.Errorf("llm synthesis failed: %w", err)
		}
		if resp != nil {
			answer = resp.Content
		}
	} else {
		// Fallback when LLM is unavailable
		if len(refs) > 0 {
			answer = fmt.Sprintf("질문 %q에 대해 총 %d개의 관련 커뮤니티 게시글을 찾았습니다:\n", query, len(refs))
			for i, r := range refs {
				answer += fmt.Sprintf("%d. [문서 %s] %s\n", i+1, r.DocID, r.Snippet)
			}
		} else {
			answer = fmt.Sprintf("질문 %q에 대한 관련 커뮤니티 문서를 찾지 못했습니다.", query)
		}
	}

	return &CuratorResponse{
		Query:       query,
		Answer:      answer,
		References:  refs,
		SourceCount: len(refs),
	}, nil
}
