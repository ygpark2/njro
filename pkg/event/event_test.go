package event

import (
	"context"
	"sync"
	"testing"
)

func TestInMemoryEventBus(t *testing.T) {
	bus := NewInMemoryEventBus()
	defer bus.Close()

	var mu sync.Mutex
	received := make([]string, 0)

	bus.Subscribe(TypePostCreated, func(ctx context.Context, evt Event) error {
		mu.Lock()
		defer mu.Unlock()
		payload := evt.Payload().(PostCreatedPayload)
		received = append(received, payload.Title)
		return nil
	})

	evt := NewBaseEvent(TypePostCreated, PostCreatedPayload{
		PostID:  "post-1",
		Title:   "First Post",
		Content: "Hello World",
	})

	err := bus.Publish(context.Background(), evt)
	if err != nil {
		t.Fatalf("unexpected error publishing: %v", err)
	}

	mu.Lock()
	if len(received) != 1 || received[0] != "First Post" {
		t.Errorf("expected ['First Post'], got %v", received)
	}
	mu.Unlock()
}

func TestEventBusMultipleSubscribers(t *testing.T) {
	bus := NewInMemoryEventBus()
	defer bus.Close()

	var counter int
	var mu sync.Mutex

	h1 := func(ctx context.Context, evt Event) error {
		mu.Lock()
		counter++
		mu.Unlock()
		return nil
	}
	h2 := func(ctx context.Context, evt Event) error {
		mu.Lock()
		counter += 10
		mu.Unlock()
		return nil
	}

	bus.Subscribe(TypePostCreated, h1)
	bus.Subscribe(TypePostCreated, h2)

	evt := NewBaseEvent(TypePostCreated, PostCreatedPayload{PostID: "post-10"})
	if err := bus.Publish(context.Background(), evt); err != nil {
		t.Fatalf("publish error: %v", err)
	}

	mu.Lock()
	if counter != 11 {
		t.Errorf("expected counter 11, got %d", counter)
	}
	mu.Unlock()
}
