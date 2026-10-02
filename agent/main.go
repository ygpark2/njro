package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ygpark2/njro/agent/client"
	"github.com/ygpark2/njro/agent/mcp"
	"github.com/ygpark2/njro/agent/orchestrator"
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

		llm := orchestrator.NewDefaultLLMClient()
		agent := orchestrator.NewAgent(llm, registry)

		fmt.Printf("🤖 [AI Agent] 작업 시작: %q\n", prompt)
		resp, err := agent.Run(context.Background(), prompt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "에이전트 실행 실패: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("\n=== [AI Agent 최종 응답] ===")
		fmt.Println(resp)

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
