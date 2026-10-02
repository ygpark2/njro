package event

import (
	"context"
	"sync"
	"time"
)

// Event represents a domain event within the system.
type Event interface {
	Type() string
	Payload() any
	Timestamp() time.Time
}

// EventHandler processes an event.
type EventHandler func(ctx context.Context, evt Event) error

// EventBus is the interface for publishing and subscribing to domain events.
type EventBus interface {
	Publish(ctx context.Context, evt Event) error
	Subscribe(eventType string, handler EventHandler) (unsubscribe func())
	Close()
}

// BaseEvent provides common fields for domain events.
type BaseEvent struct {
	eventType string
	payload   any
	timestamp time.Time
}

func NewBaseEvent(eventType string, payload any) BaseEvent {
	return BaseEvent{
		eventType: eventType,
		payload:   payload,
		timestamp: time.Now(),
	}
}

func (e BaseEvent) Type() string        { return e.eventType }
func (e BaseEvent) Payload() any        { return e.payload }
func (e BaseEvent) Timestamp() time.Time { return e.timestamp }

// Common domain event type constants
const (
	TypePostCreated    = "post.created"
	TypePostUpdated    = "post.updated"
	TypeCommentCreated = "comment.created"
	TypeUserRegistered = "user.registered"
)

// PostCreatedPayload represents data when a post is created.
type PostCreatedPayload struct {
	PostID    string `json:"post_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	AuthorID  string `json:"author_id"`
	UserEmail string `json:"user_email"`
}

// CommentCreatedPayload represents data when a comment is created.
type CommentCreatedPayload struct {
	CommentID string `json:"comment_id"`
	PostID    string `json:"post_id"`
	Content   string `json:"content"`
	AuthorID  string `json:"author_id"`
	UserEmail string `json:"user_email"`
}

// InMemoryEventBus is an in-process, concurrent event bus implementation.
type InMemoryEventBus struct {
	mu          sync.RWMutex
	subscribers map[string][]EventHandler
	closed      bool
}

// NewInMemoryEventBus creates a new InMemoryEventBus instance.
func NewInMemoryEventBus() *InMemoryEventBus {
	return &InMemoryEventBus{
		subscribers: make(map[string][]EventHandler),
	}
}

// Subscribe registers a handler for the given event type.
func (b *InMemoryEventBus) Subscribe(eventType string, handler EventHandler) func() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return func() {}
	}

	b.subscribers[eventType] = append(b.subscribers[eventType], handler)

	// Return unsubscribe func
	return func() {
		b.mu.Lock()
		defer b.mu.Unlock()
		handlers := b.subscribers[eventType]
		for i, h := range handlers {
			// Compare func pointers
			if &h == &handler {
				b.subscribers[eventType] = append(handlers[:i], handlers[i+1:]...)
				break
			}
		}
	}
}

// Publish sends the event to all registered handlers for its type.
func (b *InMemoryEventBus) Publish(ctx context.Context, evt Event) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if b.closed {
		return nil
	}

	handlers, ok := b.subscribers[evt.Type()]
	if !ok || len(handlers) == 0 {
		return nil
	}

	// Dispatch to handlers
	for _, h := range handlers {
		// Execute handler with context
		if err := h(ctx, evt); err != nil {
			return err
		}
	}
	return nil
}

// Close closes the event bus.
func (b *InMemoryEventBus) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.closed = true
	b.subscribers = make(map[string][]EventHandler)
}
