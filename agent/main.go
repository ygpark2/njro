package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ygpark2/njro/agent/client"
	"github.com/ygpark2/njro/agent/curator"
	"github.com/ygpark2/njro/agent/mcp"
	"github.com/ygpark2/njro/agent/moderator"
	"github.com/ygpark2/njro/agent/orchestrator"
	"github.com/ygpark2/njro/pkg/event"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(0)
	}

	command := os.Args[1]

	// 1. Initialize Backend gRPC Clients
	cfg := client.DefaultConfig()
	backendClients, err := client.NewBackendClients(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: backend client init error: %v\n", err)
	}
	defer backendClients.Close()

	// 2. Initialize Tool Registry
	registry := mcp.NewToolRegistry(backendClients)
	llm := orchestrator.NewDefaultLLMClient()

	switch command {
	case "mcp":
		// Run MCP JSON-RPC Server over stdio (Claude Desktop / Cursor integration)
		server := mcp.NewServer(registry)
		if err := server.RunStdio(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "MCP server exited with error: %v\n", err)
			os.Exit(1)
		}

	case "tools":
		// Print list of available MCP tools in pretty JSON
		tools := registry.ListTools()
		data, _ := json.MarshalIndent(tools, "", "  ")
		fmt.Println(string(data))

	case "chat":
		// Run AI Agent ReAct loop with user prompt
		prompt := "시스템 상태 및 게시판 목록을 확인해줘."
		if len(os.Args) > 2 {
			prompt = strings.Join(os.Args[2:], " ")
		}

		agent := orchestrator.NewAgent(llm, registry)
		fmt.Printf("🤖 [AI Agent] 작업 시작: %q\n", prompt)
		resp, err := agent.Run(context.Background(), prompt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "에이전트 실행 실패: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("\n=== [AI Agent 최종 응답] ===")
		fmt.Println(resp)

	case "moderate":
		// Run Community Moderator Agent on a post
		postID := "sample-post-id"
		title := "샘플 게시글 제목"
		content := "샘플 본문 내용입니다."
		authorEmail := "author@example.com"

		if len(os.Args) > 2 {
			postID = os.Args[2]
		}
		if len(os.Args) > 3 {
			title = os.Args[3]
		}
		if len(os.Args) > 4 {
			content = strings.Join(os.Args[4:], " ")
		}

		modAgent := moderator.NewModeratorAgent(llm, backendClients, nil, moderator.DefaultConfig())
		fmt.Printf("🛡️ [Moderator Agent] 게시글 심사 시작: PostID=%s, Title=%q\n", postID, title)
		res, err := modAgent.ModeratePost(context.Background(), postID, title, content, authorEmail)
		if err != nil {
			fmt.Fprintf(os.Stderr, "모더레이션 실패: %v\n", err)
			os.Exit(1)
		}

		data, _ := json.MarshalIndent(res, "", "  ")
		fmt.Println("\n=== [모더레이션 결과] ===")
		fmt.Println(string(data))

	case "curate":
		// Run Knowledge Curator Agent (RAG)
		query := "MSA 아키텍처 및 분산 트랜잭션 관련 글 요약"
		if len(os.Args) > 2 {
			query = strings.Join(os.Args[2:], " ")
		}

		curAgent := curator.NewCuratorAgent(llm, backendClients)
		fmt.Printf("📚 [Curator Agent] 지식 큐레이션 탐색 시작: %q\n", query)
		resp, err := curAgent.Curate(context.Background(), query)
		if err != nil {
			fmt.Fprintf(os.Stderr, "큐레이션 실패: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("\n=== [큐레이션 요약 답변] ===")
		fmt.Println(resp.Answer)
		if len(resp.References) > 0 {
			fmt.Printf("\n참고 문서 (%d건):\n", len(resp.References))
			for i, ref := range resp.References {
				fmt.Printf("[%d] ID: %s - %s\n", i+1, ref.DocID, ref.Snippet)
			}
		}

	case "watch":
		// Run event-driven background listener
		bus := event.NewInMemoryEventBus()
		defer bus.Close()

		modAgent := moderator.NewModeratorAgent(llm, backendClients, bus, moderator.DefaultConfig())
		unsub, err := modAgent.StartEventListener(context.Background())
		if err != nil {
			fmt.Fprintf(os.Stderr, "이벤트 리스너 등록 실패: %v\n", err)
			os.Exit(1)
		}
		defer unsub()

		fmt.Println("🎧 [EventBus Watcher] 실시간 이벤트 수신 대기 중... (Ctrl+C로 종료)")
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		fmt.Println("\n이벤트 리스너를 안전하게 종료합니다.")

	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println(`njro AI Agent & MCP Layer

Usage:
  agent mcp               Run as Model Context Protocol (MCP) server over stdio
  agent chat <prompt>     Run autonomous AI Agent with natural language prompt
  agent tools             List all available tools exposed to AI and MCP clients
  agent moderate <id>     Evaluate content safety with Moderator Agent
  agent curate <query>    Search community knowledge and synthesize answer with Curator Agent
  agent watch             Listen to real-time domain events for background moderation

Environment Variables:
  OPENAI_API_KEY          (Optional) OpenAI / Claude / Ollama API Key
  OPENAI_BASE_URL         (Optional) API Base URL (default: https://api.openai.com/v1)
  LLM_MODEL               (Optional) Model name (default: gpt-4o)
  BOARD_SERVICE_URL       (Default: localhost:8081)
  CONTENT_SERVICE_URL     (Default: localhost:8082)
  POST_SERVICE_URL        (Default: localhost:8083)
  COMMENT_SERVICE_URL     (Default: localhost:8084)
  ACCOUNT_SERVICE_URL     (Default: localhost:8085)
  EMAILER_SERVICE_URL     (Default: localhost:8086)
  SEARCH_SERVICE_URL      (Default: localhost:8087)`)
}
