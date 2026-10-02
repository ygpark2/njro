package orchestrator

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ygpark2/njro/agent/mcp"
)

// Message represents a chat message.
type Message struct {
	Role       string     `json:"role"`
	Content    string     `json:"content"`
	Name       string     `json:"name,omitempty"`
	ToolCalls  []ToolCall `json:"tool_calls,omitempty"`
	ToolCallID string     `json:"tool_call_id,omitempty"`
}

// ToolCall represents a tool execution request by the LLM.
type ToolCall struct {
	ID       string       `json:"id"`
	Type     string       `json:"type"`
	Function FunctionCall `json:"function"`
}

// FunctionCall represents the function name and JSON arguments string.
type FunctionCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

// LLMClient represents an abstraction for calling an LLM.
type LLMClient interface {
	Chat(ctx context.Context, messages []Message, tools []mcp.ToolDefinition) (*Message, error)
}

// NewDefaultLLMClient returns an OpenAI-compatible client if OPENAI_API_KEY or LLM_API_KEY is present,
// otherwise returns the smart MockLLMClient.
func NewDefaultLLMClient() LLMClient {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		apiKey = os.Getenv("LLM_API_KEY")
	}

	if apiKey != "" {
		baseURL := os.Getenv("OPENAI_BASE_URL")
		if baseURL == "" {
			baseURL = "https://api.openai.com/v1"
		}
		model := os.Getenv("LLM_MODEL")
		if model == "" {
			model = "gpt-4o"
		}
		return &OpenAILLMClient{
			APIKey:  apiKey,
			BaseURL: strings.TrimRight(baseURL, "/"),
			Model:   model,
			Client:  &http.Client{Timeout: 60 * time.Second},
		}
	}

	return &MockLLMClient{}
}

// OpenAILLMClient implements LLMClient using standard OpenAI-compatible Chat Completions API.
type OpenAILLMClient struct {
	APIKey  string
	BaseURL string
	Model   string
	Client  *http.Client
}

func (c *OpenAILLMClient) Chat(ctx context.Context, messages []Message, tools []mcp.ToolDefinition) (*Message, error) {
	type apiFunction struct {
		Name        string         `json:"name"`
		Description string         `json:"description"`
		Parameters  map[string]any `json:"parameters"`
	}
	type apiTool struct {
		Type     string      `json:"type"`
		Function apiFunction `json:"function"`
	}

	apiTools := make([]apiTool, 0, len(tools))
	for _, t := range tools {
		apiTools = append(apiTools, apiTool{
			Type: "function",
			Function: apiFunction{
				Name:        t.Name,
				Description: t.Description,
				Parameters:  t.InputSchema,
			},
		})
	}

	reqBody := map[string]any{
		"model":    c.Model,
		"messages": messages,
	}
	if len(apiTools) > 0 {
		reqBody["tools"] = apiTools
	}

	payload, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseURL+"/chat/completions", bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("LLM API returned status %d: %s", resp.StatusCode, string(body))
	}

	var res struct {
		Choices []struct {
			Message Message `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	if len(res.Choices) == 0 {
		return nil, fmt.Errorf("no choice returned by LLM")
	}

	return &res.Choices[0].Message, nil
}

// MockLLMClient provides offline rule-based natural language intent recognition for testing.
type MockLLMClient struct{}

func (m *MockLLMClient) Chat(ctx context.Context, messages []Message, tools []mcp.ToolDefinition) (*Message, error) {
	lastMsg := messages[len(messages)-1]

	// If the last message was a tool result, synthesize a final answer
	if lastMsg.Role == "tool" {
		return &Message{
			Role:    "assistant",
			Content: fmt.Sprintf("도구 실행 결과는 다음과 같습니다:\n\n%s", lastMsg.Content),
		}, nil
	}

	userText := strings.ToLower(lastMsg.Content)

	// Rule 1: Search community
	if strings.Contains(userText, "검색") || strings.Contains(userText, "search") || strings.Contains(userText, "찾아") {
		query := "공지"
		words := strings.Fields(lastMsg.Content)
		for _, w := range words {
			if !strings.Contains(w, "검색") && !strings.Contains(w, "search") && !strings.Contains(w, "찾아") {
				query = w
				break
			}
		}
		argsJSON, _ := json.Marshal(map[string]any{"query": query})
		return &Message{
			Role: "assistant",
			ToolCalls: []ToolCall{
				{
					ID:   "call_search_1",
					Type: "function",
					Function: FunctionCall{
						Name:      "search_community",
						Arguments: string(argsJSON),
					},
				},
			},
		}, nil
	}

	// Rule 2: List boards
	if strings.Contains(userText, "게시판") || strings.Contains(userText, "board") {
		argsJSON, _ := json.Marshal(map[string]any{"page_size": 10})
		return &Message{
			Role: "assistant",
			ToolCalls: []ToolCall{
				{
					ID:   "call_board_1",
					Type: "function",
					Function: FunctionCall{
						Name:      "list_boards",
						Arguments: string(argsJSON),
					},
				},
			},
		}, nil
	}

	// Rule 3: List posts
	if strings.Contains(userText, "게시글") || strings.Contains(userText, "글") || strings.Contains(userText, "post") {
		argsJSON, _ := json.Marshal(map[string]any{"page_size": 10})
		return &Message{
			Role: "assistant",
			ToolCalls: []ToolCall{
				{
					ID:   "call_post_1",
					Type: "function",
					Function: FunctionCall{
						Name:      "list_posts",
						Arguments: string(argsJSON),
					},
				},
			},
		}, nil
	}

	// Rule 4: Send email
	if strings.Contains(userText, "메일") || strings.Contains(userText, "email") {
		argsJSON, _ := json.Marshal(map[string]any{
			"to":      "admin@example.com",
			"subject": "[알림] 시스템 공지사항",
			"body":    "에이전트가 자동 생성한 알림 메시지입니다.",
		})
		return &Message{
			Role: "assistant",
			ToolCalls: []ToolCall{
				{
					ID:   "call_email_1",
					Type: "function",
					Function: FunctionCall{
						Name:      "send_email",
						Arguments: string(argsJSON),
					},
				},
			},
		}, nil
	}

	// Fallback general response
	return &Message{
		Role:    "assistant",
		Content: "안녕하세요! 저는 njro 마이크로서비스 백엔드를 관리하는 AI 에이전트입니다. 게시판 조회, 게시글 검색, 메일 발송 등을 자연어로 요청해 주세요.",
	}, nil
}
