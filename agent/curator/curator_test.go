package curator

import (
	"context"
	"strings"
	"testing"

	"github.com/ygpark2/njro/agent/mcp"
	"github.com/ygpark2/njro/agent/orchestrator"
)

type mockLLMClient struct {
	answer string
}

func (m *mockLLMClient) Chat(ctx context.Context, messages []orchestrator.Message, tools []mcp.ToolDefinition) (*orchestrator.Message, error) {
	return &orchestrator.Message{
		Role:    "assistant",
		Content: m.answer,
	}, nil
}

func TestCuratorAgent_WithoutClients(t *testing.T) {
	mockLLM := &mockLLMClient{
		answer: "분산 트랜잭션은 2PC 또는 Saga 패턴을 사용하여 서비스 간 데이터 일관성을 유지할 수 있습니다.",
	}

	agent := NewCuratorAgent(mockLLM, nil)
	resp, err := agent.Curate(context.Background(), "MSA 분산 트랜잭션 구현 방법")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Query != "MSA 분산 트랜잭션 구현 방법" {
		t.Errorf("expected query preserved, got %s", resp.Query)
	}
	if !strings.Contains(resp.Answer, "Saga 패턴") {
		t.Errorf("expected answer to contain Saga 패턴, got %s", resp.Answer)
	}
}

func TestCuratorAgent_FallbackWithoutLLM(t *testing.T) {
	agent := NewCuratorAgent(nil, nil)
	resp, err := agent.Curate(context.Background(), "알 수 없는 질문")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Answer == "" {
		t.Errorf("expected non-empty fallback answer")
	}
}
