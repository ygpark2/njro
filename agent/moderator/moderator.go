package moderator

import (
	"context"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/ygpark2/njro/agent/client"
	"github.com/ygpark2/njro/agent/orchestrator"
	"github.com/ygpark2/njro/pkg/event"
	emailerpb "github.com/ygpark2/njro/service/emailer/proto/emailer"
	postpb "github.com/ygpark2/njro/service/post/ent/proto/entpb"
)

// ModerationResult represents the outcome of a content review.
type ModerationResult struct {
	PostID      string `json:"post_id"`
	IsHarmful   bool   `json:"is_harmful"`
	Category    string `json:"category,omitempty"` // SPAM, ABUSE, AD, PROFANITY
	Reason      string `json:"reason"`
	ActionTaken string `json:"action_taken"` // NONE, BLINDED, NOTIFIED
}

// Config controls the behavior of ModeratorAgent.
type Config struct {
	AutoBlindSpam bool
	NotifyAuthor  bool
	SenderEmail   string
}

// DefaultConfig returns reasonable default settings.
func DefaultConfig() Config {
	return Config{
		AutoBlindSpam: true,
		NotifyAuthor:  true,
		SenderEmail:   "moderator@community.njro",
	}
}

// ModeratorAgent monitors posts and comments, evaluates content safety, and takes automated moderation actions.
type ModeratorAgent struct {
	llm      orchestrator.LLMClient
	clients  *client.BackendClients
	eventBus event.EventBus
	cfg      Config
}

// NewModeratorAgent creates a new ModeratorAgent.
func NewModeratorAgent(llm orchestrator.LLMClient, clients *client.BackendClients, eventBus event.EventBus, cfg Config) *ModeratorAgent {
	return &ModeratorAgent{
		llm:      llm,
		clients:  clients,
		eventBus: eventBus,
		cfg:      cfg,
	}
}

// ModeratePost analyzes a post and executes automated moderation actions if violated.
func (m *ModeratorAgent) ModeratePost(ctx context.Context, postID string, title, content, authorEmail string) (*ModerationResult, error) {
	log.Info().Str("post_id", postID).Str("title", title).Msg("ModeratorAgent: reviewing post content")

	// 1. Analyze with LLM or safety heuristics
	isHarmful, category, reason := m.evaluateContent(ctx, title, content)

	res := &ModerationResult{
		PostID:      postID,
		IsHarmful:   isHarmful,
		Category:    category,
		Reason:      reason,
		ActionTaken: "NONE",
	}

	if !isHarmful {
		log.Info().Str("post_id", postID).Msg("ModeratorAgent: post approved (clean)")
		return res, nil
	}

	log.Warn().Str("post_id", postID).Str("category", category).Str("reason", reason).
		Msg("ModeratorAgent: harmful content detected, taking mitigation action")

	// 2. Action: Delete / Hide the post via Post gRPC client
	if m.cfg.AutoBlindSpam && m.clients != nil && m.clients.Post != nil {
		if rawID, err := client.ParseUUID(postID); err == nil {
			_, err := m.clients.Post.Delete(ctx, &postpb.DeletePostRequest{
				Id: rawID,
			})
			if err != nil {
				log.Error().Err(err).Str("post_id", postID).Msg("failed to delete harmful post")
			} else {
				res.ActionTaken = "BLINDED"
			}
		}
	}

	// 3. Action: Send warning notification to author via Emailer gRPC client
	if m.cfg.NotifyAuthor && authorEmail != "" && m.clients != nil && m.clients.Emailer != nil {
		subject := fmt.Sprintf("[Community Notice] Your post #%s has been hidden due to community guidelines", postID)
		body := fmt.Sprintf(
			"Hello,\n\nYour post titled %q was flagged by our automated AI moderation system.\nReason: %s (%s)\n\nIf you believe this was an error, please contact support.\n\nCommunity Operations Team",
			title, reason, category,
		)
		_, err := m.clients.Emailer.SendEmail(ctx, &emailerpb.SendEmailRequest{
			To:      authorEmail,
			From:    m.cfg.SenderEmail,
			Subject: subject,
			Body:    body,
		})
		if err != nil {
			log.Error().Err(err).Str("to", authorEmail).Msg("failed to send moderation notification email")
		} else {
			if res.ActionTaken == "BLINDED" {
				res.ActionTaken = "BLINDED_AND_NOTIFIED"
			} else {
				res.ActionTaken = "NOTIFIED"
			}
		}
	}

	return res, nil
}

// StartEventListener subscribes to PostCreated domain events from EventBus for real-time auto moderation.
func (m *ModeratorAgent) StartEventListener(ctx context.Context) (func(), error) {
	if m.eventBus == nil {
		return nil, fmt.Errorf("eventBus is nil")
	}

	unsubscribe := m.eventBus.Subscribe(event.TypePostCreated, func(c context.Context, evt event.Event) error {
		payload, ok := evt.Payload().(event.PostCreatedPayload)
		if !ok {
			return nil
		}

		// Perform asynchronous moderation
		go func(p event.PostCreatedPayload) {
			_, err := m.ModeratePost(context.Background(), p.PostID, p.Title, p.Content, p.UserEmail)
			if err != nil {
				log.Error().Err(err).Str("post_id", p.PostID).Msg("error during event-driven post moderation")
			}
		}(payload)

		return nil
	})

	log.Info().Msg("ModeratorAgent: registered event listener for post.created")
	return unsubscribe, nil
}

// evaluateContent uses LLM or heuristic rules to evaluate text safety.
func (m *ModeratorAgent) evaluateContent(ctx context.Context, title, content string) (bool, string, string) {
	fullText := strings.ToLower(title + " " + content)

	// Quick heuristic check for known spam/malicious keywords
	badKeywords := []string{"불법도박", "카지노", "대출문의", "성인광고", "viagra", "crypto scam"}
	for _, kw := range badKeywords {
		if strings.Contains(fullText, kw) {
			return true, "SPAM", fmt.Sprintf("contains prohibited advertisement keyword: %q", kw)
		}
	}

	// If LLM client is available, prompt for nuanced semantic judgment
	if m.llm != nil {
		prompt := fmt.Sprintf(
			"Analyze the following community post for spam, severe toxicity, personal information leaks, or fraud. If violated, respond with VIOLATION: <category> - <reason>. Otherwise respond with SAFE.\n\nTitle: %s\nContent: %s",
			title, content,
		)
		resp, err := m.llm.Chat(ctx, []orchestrator.Message{
			{Role: "user", Content: prompt},
		}, nil)
		if err == nil && resp != nil && strings.HasPrefix(strings.TrimSpace(resp.Content), "VIOLATION:") {
			parts := strings.SplitN(resp.Content, "-", 2)
			category := strings.TrimPrefix(parts[0], "VIOLATION:")
			category = strings.TrimSpace(category)
			reason := "Flagged by AI moderator"
			if len(parts) > 1 {
				reason = strings.TrimSpace(parts[1])
			}
			return true, category, reason
		}
	}

	return false, "", ""
}
