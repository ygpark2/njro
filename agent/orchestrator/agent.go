package orchestrator

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ygpark2/njro/agent/mcp"
)

// Agent manages the conversation and autonomous tool execution loop.
type Agent struct {
	llm      LLMClient
	registry *mcp.ToolRegistry
	logger   *log.Logger
}

// NewAgent creates a new Agent instance.
func NewAgent(llm LLMClient, registry *mcp.ToolRegistry) *Agent {
	return &Agent{
		llm:      llm,
		registry: registry,
		logger:   log.New(os.Stderr, "[AI Agent] ", log.LstdFlags),
	}
}

// Run executes a natural language task using the ReAct / Function Calling loop.
func (a *Agent) Run(ctx context.Context, userPrompt string) (string, error) {
	messages := []Message{
		{
			Role: "system",
			Content: `당신은 njro 고성능 마이크로서비스 백엔드 시스템을 총괄하는 수석 AI 운영 에이전트입니다.
사용자의 요청에 따라 게시판, 게시글, 댓글, 사용자, 검색, 이메일러 등의 도구(Tool)를 능동적으로 호출하여 작업을 완수하세요.
도구 실행 결과를 종합하여 사용자에게 친절하고 명확한 한국어로 최종 답변을 작성하세요.`,
		},
		{
			Role:    "user",
			Content: userPrompt,
		},
	}

	tools := a.registry.ListTools()
	const maxIterations = 6

	for iter := 0; iter < maxIterations; iter++ {
		resp, err := a.llm.Chat(ctx, messages, tools)
		if err != nil {
			return "", fmt.Errorf("LLM chat error at step %d: %w", iter, err)
		}

		messages = append(messages, *resp)

		// If no tools were called, we have our final synthesized response
		if len(resp.ToolCalls) == 0 {
			return resp.Content, nil
		}

		// Execute all requested tool calls
		for _, call := range resp.ToolCalls {
			a.logger.Printf("Executing Tool: %s with args: %s\n", call.Function.Name, call.Function.Arguments)

			var args map[string]any
			if err := json.Unmarshal([]byte(call.Function.Arguments), &args); err != nil {
				args = make(map[string]any)
			}

			output, execErr := a.registry.Call(ctx, call.Function.Name, args)
			toolContent := output
			if execErr != nil {
				toolContent = fmt.Sprintf("Error executing %s: %v", call.Function.Name, execErr)
			}

			messages = append(messages, Message{
				Role:       "tool",
				Name:       call.Function.Name,
				Content:    toolContent,
				ToolCallID: call.ID,
			})
		}
	}

	return "", fmt.Errorf("agent exceeded maximum iteration limit (%d)", maxIterations)
}
