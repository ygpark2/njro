package client

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	accountpb "github.com/ygpark2/njro/service/account/ent/proto/entpb"
	boardpb "github.com/ygpark2/njro/service/board/ent/proto/entpb"
	commentpb "github.com/ygpark2/njro/service/comment/ent/proto/entpb"
	contentpb "github.com/ygpark2/njro/service/content/ent/proto/entpb"
	emailerpb "github.com/ygpark2/njro/service/emailer/proto/emailer"
	postpb "github.com/ygpark2/njro/service/post/ent/proto/entpb"
	searchpb "github.com/ygpark2/njro/service/search/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Config holds service target addresses.
type Config struct {
	BoardTarget   string
	ContentTarget string
	PostTarget    string
	CommentTarget string
	AccountTarget string
	EmailerTarget string
	SearchTarget  string
}

// DefaultConfig loads targets from environment variables or returns defaults.
func DefaultConfig() Config {
	getEnv := func(key, fallback string) string {
		if val := os.Getenv(key); val != "" {
			return val
		}
		return fallback
	}

	return Config{
		BoardTarget:   getEnv("BOARD_SERVICE_URL", "localhost:8081"),
		ContentTarget: getEnv("CONTENT_SERVICE_URL", "localhost:8082"),
		PostTarget:    getEnv("POST_SERVICE_URL", "localhost:8083"),
		CommentTarget: getEnv("COMMENT_SERVICE_URL", "localhost:8084"),
		AccountTarget: getEnv("ACCOUNT_SERVICE_URL", "localhost:8085"),
		EmailerTarget: getEnv("EMAILER_SERVICE_URL", "localhost:8086"),
		SearchTarget:  getEnv("SEARCH_SERVICE_URL", "localhost:8087"),
	}
}

// BackendClients holds initialized gRPC client connections to all services.
type BackendClients struct {
	Board   boardpb.BoardServiceClient
	Content contentpb.ContentServiceClient
	Post    postpb.PostServiceClient
	Comment commentpb.CommentServiceClient
	Account accountpb.UserServiceClient
	Emailer emailerpb.EmailerServiceClient
	Search  searchpb.SearchServiceClient

	conns []*grpc.ClientConn
}

// NewBackendClients initializes gRPC connections to backend services.
func NewBackendClients(cfg Config) (*BackendClients, error) {
	clients := &BackendClients{
		conns: make([]*grpc.ClientConn, 0, 7),
	}

	dial := func(target string) (*grpc.ClientConn, error) {
		conn, err := grpc.NewClient(
			target,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create gRPC client for %s: %w", target, err)
		}
		clients.conns = append(clients.conns, conn)
		return conn, nil
	}

	boardConn, err := dial(cfg.BoardTarget)
	if err != nil {
		clients.Close()
		return nil, err
	}
	clients.Board = boardpb.NewBoardServiceClient(boardConn)

	contentConn, err := dial(cfg.ContentTarget)
	if err != nil {
		clients.Close()
		return nil, err
	}
	clients.Content = contentpb.NewContentServiceClient(contentConn)

	postConn, err := dial(cfg.PostTarget)
	if err != nil {
		clients.Close()
		return nil, err
	}
	clients.Post = postpb.NewPostServiceClient(postConn)

	commentConn, err := dial(cfg.CommentTarget)
	if err != nil {
		clients.Close()
		return nil, err
	}
	clients.Comment = commentpb.NewCommentServiceClient(commentConn)

	accountConn, err := dial(cfg.AccountTarget)
	if err != nil {
		clients.Close()
		return nil, err
	}
	clients.Account = accountpb.NewUserServiceClient(accountConn)

	emailerConn, err := dial(cfg.EmailerTarget)
	if err != nil {
		clients.Close()
		return nil, err
	}
	clients.Emailer = emailerpb.NewEmailerServiceClient(emailerConn)

	searchConn, err := dial(cfg.SearchTarget)
	if err != nil {
		clients.Close()
		return nil, err
	}
	clients.Search = searchpb.NewSearchServiceClient(searchConn)

	return clients, nil
}

// Close closes all active gRPC connections.
func (c *BackendClients) Close() {
	for _, conn := range c.conns {
		if conn != nil {
			_ = conn.Close()
		}
	}
}

// ParseUUID converts a string UUID into bytes slice for Ent proto.
func ParseUUID(idStr string) ([]byte, error) {
	u, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}
	return u[:], nil
}

// UUIDToString converts bytes slice from Ent proto into a canonical UUID string.
func UUIDToString(b []byte) string {
	if len(b) != 16 {
		return ""
	}
	u, err := uuid.FromBytes(b)
	if err != nil {
		return ""
	}
	return u.String()
}

// ContextWithTimeout returns a context with default 5-second timeout.
func ContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}
