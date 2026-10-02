package moderator

import (
	"context"
	"testing"
	"time"

	"github.com/ygpark2/njro/pkg/event"
)

func TestModeratorAgent_CleanContent(t *testing.T) {
	agent := NewModeratorAgent(nil, nil, nil, DefaultConfig())

	res, err := agent.ModeratePost(context.Background(), "00000000-0000-0000-0000-000000000101", "안녕하세요 좋은 아침입니다", "오늘도 활기찬 하루 되세요!", "user@example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.IsHarmful {
		t.Errorf("expected clean content, got harmful")
	}
	if res.ActionTaken != "NONE" {
		t.Errorf("expected action NONE, got %s", res.ActionTaken)
	}
}

func TestModeratorAgent_HeuristicSpamDetection(t *testing.T) {
	agent := NewModeratorAgent(nil, nil, nil, DefaultConfig())

	res, err := agent.ModeratePost(context.Background(), "00000000-0000-0000-0000-000000000102", "최고의 승률 카지노 게임", "지금 바로 가입시 무료머니 지급 불법도박 사이트 바로가기", "spammer@spam.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !res.IsHarmful {
		t.Fatalf("expected harmful to be true")
	}
	if res.Category != "SPAM" {
		t.Errorf("expected category SPAM, got %s", res.Category)
	}
}

func TestModeratorAgent_EventListener(t *testing.T) {
	bus := event.NewInMemoryEventBus()
	defer bus.Close()

	agent := NewModeratorAgent(nil, nil, bus, DefaultConfig())

	unsub, err := agent.StartEventListener(context.Background())
	if err != nil {
		t.Fatalf("failed to start event listener: %v", err)
	}
	defer unsub()

	// Publish a clean post event
	evt := event.NewBaseEvent(event.TypePostCreated, event.PostCreatedPayload{
		PostID:    "00000000-0000-0000-0000-000000000201",
		Title:     "테스트 게시글",
		Content:   "정상적인 내용입니다.",
		UserEmail: "test@example.com",
	})

	if err := bus.Publish(context.Background(), evt); err != nil {
		t.Fatalf("publish error: %v", err)
	}

	// Give goroutine a moment to complete
	time.Sleep(50 * time.Millisecond)
}
