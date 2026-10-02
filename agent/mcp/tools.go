package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ygpark2/njro/agent/client"
	accountpb "github.com/ygpark2/njro/service/account/ent/proto/entpb"
	boardpb "github.com/ygpark2/njro/service/board/ent/proto/entpb"
	contentpb "github.com/ygpark2/njro/service/content/ent/proto/entpb"
	emailerpb "github.com/ygpark2/njro/service/emailer/proto/emailer"
	postpb "github.com/ygpark2/njro/service/post/ent/proto/entpb"
	searchpb "github.com/ygpark2/njro/service/search/proto"
)

// ToolDefinition defines an MCP tool.
type ToolDefinition struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	InputSchema map[string]any `json:"inputSchema"`
}

// ToolHandler is a function that executes an MCP tool call.
type ToolHandler func(ctx context.Context, args map[string]any) (string, error)

// ToolRegistry manages all registered MCP tools and their handlers.
type ToolRegistry struct {
	tools    []ToolDefinition
	handlers map[string]ToolHandler
	clients  *client.BackendClients
}

// NewToolRegistry initializes tools with backend gRPC clients.
func NewToolRegistry(clients *client.BackendClients) *ToolRegistry {
	r := &ToolRegistry{
		tools:    make([]ToolDefinition, 0),
		handlers: make(map[string]ToolHandler),
		clients:  clients,
	}

	r.registerAll()
	return r
}

// ListTools returns the list of available MCP tool definitions.
func (r *ToolRegistry) ListTools() []ToolDefinition {
	return r.tools
}

// Call executes the specified tool by name with arguments.
func (r *ToolRegistry) Call(ctx context.Context, name string, args map[string]any) (string, error) {
	handler, ok := r.handlers[name]
	if !ok {
		return "", fmt.Errorf("unknown tool: %s", name)
	}
	return handler(ctx, args)
}

func (r *ToolRegistry) addTool(def ToolDefinition, handler ToolHandler) {
	r.tools = append(r.tools, def)
	r.handlers[def.Name] = handler
}

func (r *ToolRegistry) registerAll() {
	// 1. search_community
	r.addTool(ToolDefinition{
		Name:        "search_community",
		Description: "Search community posts and articles by keyword via Bleve/Elasticsearch.",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"query": map[string]any{
					"type":        "string",
					"description": "Search keyword or query string",
				},
				"limit": map[string]any{
					"type":        "integer",
					"description": "Max number of results to return (default: 10)",
				},
			},
			"required": []string{"query"},
		},
	}, r.handleSearchCommunity)

	// 2. list_boards
	r.addTool(ToolDefinition{
		Name:        "list_boards",
		Description: "List all bulletin boards with their titles, descriptions, and IDs.",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"page_size": map[string]any{
					"type":        "integer",
					"description": "Number of boards per page (default: 10)",
				},
			},
		},
	}, r.handleListBoards)

	// 3. get_board
	r.addTool(ToolDefinition{
		Name:        "get_board",
		Description: "Get detailed information of a specific board by its UUID.",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"board_id": map[string]any{
					"type":        "string",
					"description": "Board UUID string",
				},
			},
			"required": []string{"board_id"},
		},
	}, r.handleGetBoard)

	// 4. create_board
	r.addTool(ToolDefinition{
		Name:        "create_board",
		Description: "Create a new bulletin board in the community.",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"title": map[string]any{
					"type":        "string",
					"description": "Board title",
				},
				"description": map[string]any{
					"type":        "string",
					"description": "Board description or purpose",
				},
			},
			"required": []string{"title"},
		},
	}, r.handleCreateBoard)

	// 5. list_posts
	r.addTool(ToolDefinition{
		Name:        "list_posts",
		Description: "List recent community posts.",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"page_size": map[string]any{
					"type":        "integer",
					"description": "Number of posts to return (default: 10)",
				},
			},
		},
	}, r.handleListPosts)

	// 6. get_post
	r.addTool(ToolDefinition{
		Name:        "get_post",
		Description: "Get post metadata (title, author, created time) by post UUID.",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"post_id": map[string]any{
					"type":        "string",
					"description": "Post UUID string",
				},
			},
			"required": []string{"post_id"},
		},
	}, r.handleGetPost)

	// 7. create_post
	r.addTool(ToolDefinition{
		Name:        "create_post",
		Description: "Create a new post in a specific bulletin board.",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"board_id": map[string]any{
					"type":        "string",
					"description": "Target board UUID string",
				},
				"title": map[string]any{
					"type":        "string",
					"description": "Post title",
				},
				"writer": map[string]any{
					"type":        "string",
					"description": "Author name or user ID",
				},
			},
			"required": []string{"board_id", "title", "writer"},
		},
	}, r.handleCreatePost)

	// 8. get_content
	r.addTool(ToolDefinition{
		Name:        "get_content",
		Description: "Retrieve full body text content of a post or comment.",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"content_id": map[string]any{
					"type":        "string",
					"description": "Content UUID string",
				},
			},
			"required": []string{"content_id"},
		},
	}, r.handleGetContent)

	// 9. send_email
	r.addTool(ToolDefinition{
		Name:        "send_email",
		Description: "Send an email alert or notification to a user or administrator.",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"to": map[string]any{
					"type":        "string",
					"description": "Recipient email address",
				},
				"subject": map[string]any{
					"type":        "string",
					"description": "Email subject",
				},
				"body": map[string]any{
					"type":        "string",
					"description": "Email body content",
				},
			},
			"required": []string{"to", "subject", "body"},
		},
	}, r.handleSendEmail)

	// 10. get_user
	r.addTool(ToolDefinition{
		Name:        "get_user",
		Description: "Get user account profile (email, role, created time) by user UUID.",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"user_id": map[string]any{
					"type":        "string",
					"description": "User UUID string",
				},
			},
			"required": []string{"user_id"},
		},
	}, r.handleGetUser)
}

// Handlers implementation

func (r *ToolRegistry) handleSearchCommunity(ctx context.Context, args map[string]any) (string, error) {
	query, _ := args["query"].(string)

	res, err := r.clients.Search.Search(ctx, &searchpb.SearchRequest{
		Keyword: query,
	})
	if err != nil {
		return "", fmt.Errorf("search failed: %w", err)
	}

	data, _ := json.MarshalIndent(res, "", "  ")
	return string(data), nil
}

func (r *ToolRegistry) handleListBoards(ctx context.Context, args map[string]any) (string, error) {
	pageSize := int32(10)
	if p, ok := args["page_size"].(float64); ok && p > 0 {
		pageSize = int32(p)
	}

	res, err := r.clients.Board.List(ctx, &boardpb.ListBoardRequest{
		PageSize: pageSize,
	})
	if err != nil {
		return "", fmt.Errorf("list boards failed: %w", err)
	}

	type boardSummary struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Order       uint32 `json:"order"`
	}

	boards := make([]boardSummary, 0, len(res.GetBoardList()))
	for _, b := range res.GetBoardList() {
		boards = append(boards, boardSummary{
			ID:          client.UUIDToString(b.GetId()),
			Title:       b.GetTitle(),
			Description: b.GetDescription(),
			Order:       b.GetOrder(),
		})
	}

	data, _ := json.MarshalIndent(boards, "", "  ")
	return string(data), nil
}

func (r *ToolRegistry) handleGetBoard(ctx context.Context, args map[string]any) (string, error) {
	boardID, _ := args["board_id"].(string)
	rawID, err := client.ParseUUID(boardID)
	if err != nil {
		return "", fmt.Errorf("invalid board_id UUID format: %w", err)
	}

	b, err := r.clients.Board.Get(ctx, &boardpb.GetBoardRequest{
		Id: rawID,
	})
	if err != nil {
		return "", fmt.Errorf("get board failed: %w", err)
	}

	data, _ := json.MarshalIndent(map[string]any{
		"id":          client.UUIDToString(b.GetId()),
		"title":       b.GetTitle(),
		"description": b.GetDescription(),
		"notices":     b.GetNotices(),
	}, "", "  ")
	return string(data), nil
}

func (r *ToolRegistry) handleCreateBoard(ctx context.Context, args map[string]any) (string, error) {
	title, _ := args["title"].(string)
	description, _ := args["description"].(string)

	b, err := r.clients.Board.Create(ctx, &boardpb.CreateBoardRequest{
		Board: &boardpb.Board{
			Title:       title,
			MobileTitle: title,
			Description: description,
			Search:      true,
		},
	})
	if err != nil {
		return "", fmt.Errorf("create board failed: %w", err)
	}

	return fmt.Sprintf("Board created successfully! ID: %s, Title: %s", client.UUIDToString(b.GetId()), b.GetTitle()), nil
}

func (r *ToolRegistry) handleListPosts(ctx context.Context, args map[string]any) (string, error) {
	pageSize := int32(10)
	if p, ok := args["page_size"].(float64); ok && p > 0 {
		pageSize = int32(p)
	}

	res, err := r.clients.Post.List(ctx, &postpb.ListPostRequest{
		PageSize: pageSize,
	})
	if err != nil {
		return "", fmt.Errorf("list posts failed: %w", err)
	}

	type postSummary struct {
		ID     string `json:"id"`
		Title  string `json:"title"`
		Writer string `json:"writer"`
	}

	posts := make([]postSummary, 0, len(res.GetPostList()))
	for _, p := range res.GetPostList() {
		posts = append(posts, postSummary{
			ID:     client.UUIDToString(p.GetId()),
			Title:  p.GetTitle(),
			Writer: p.GetWriter(),
		})
	}

	data, _ := json.MarshalIndent(posts, "", "  ")
	return string(data), nil
}

func (r *ToolRegistry) handleGetPost(ctx context.Context, args map[string]any) (string, error) {
	postID, _ := args["post_id"].(string)
	rawID, err := client.ParseUUID(postID)
	if err != nil {
		return "", fmt.Errorf("invalid post_id UUID format: %w", err)
	}

	p, err := r.clients.Post.Get(ctx, &postpb.GetPostRequest{
		Id: rawID,
	})
	if err != nil {
		return "", fmt.Errorf("get post failed: %w", err)
	}

	data, _ := json.MarshalIndent(map[string]any{
		"id":       client.UUIDToString(p.GetId()),
		"title":    p.GetTitle(),
		"writer":   p.GetWriter(),
		"board_id": p.GetBoardId(),
	}, "", "  ")
	return string(data), nil
}

func (r *ToolRegistry) handleCreatePost(ctx context.Context, args map[string]any) (string, error) {
	title, _ := args["title"].(string)
	boardIDStr, _ := args["board_id"].(string)
	writer, _ := args["writer"].(string)

	p, err := r.clients.Post.Create(ctx, &postpb.CreatePostRequest{
		Post: &postpb.Post{
			Title:   title,
			BoardId: boardIDStr,
			Writer:  writer,
		},
	})
	if err != nil {
		return "", fmt.Errorf("create post failed: %w", err)
	}

	return fmt.Sprintf("Post created successfully! ID: %s, Title: %s", client.UUIDToString(p.GetId()), p.GetTitle()), nil
}

func (r *ToolRegistry) handleGetContent(ctx context.Context, args map[string]any) (string, error) {
	contentID, _ := args["content_id"].(string)
	rawID, err := client.ParseUUID(contentID)
	if err != nil {
		return "", fmt.Errorf("invalid content_id UUID format: %w", err)
	}

	c, err := r.clients.Content.Get(ctx, &contentpb.GetContentRequest{
		Id: rawID,
	})
	if err != nil {
		return "", fmt.Errorf("get content failed: %w", err)
	}

	return c.GetContent(), nil
}

func (r *ToolRegistry) handleSendEmail(ctx context.Context, args map[string]any) (string, error) {
	to, _ := args["to"].(string)
	subject, _ := args["subject"].(string)
	body, _ := args["body"].(string)

	res, err := r.clients.Emailer.SendEmail(ctx, &emailerpb.SendEmailRequest{
		To:      to,
		Subject: subject,
		Body:    body,
	})
	if err != nil {
		return "", fmt.Errorf("send email failed: %w", err)
	}

	return fmt.Sprintf("Email sent successfully: %s", res.GetMessage()), nil
}

func (r *ToolRegistry) handleGetUser(ctx context.Context, args map[string]any) (string, error) {
	userID, _ := args["user_id"].(string)
	rawID, err := client.ParseUUID(userID)
	if err != nil {
		return "", fmt.Errorf("invalid user_id UUID format: %w", err)
	}

	u, err := r.clients.Account.Get(ctx, &accountpb.GetUserRequest{
		Id: rawID,
	})
	if err != nil {
		return "", fmt.Errorf("get user failed: %w", err)
	}

	data, _ := json.MarshalIndent(map[string]any{
		"id":    client.UUIDToString(u.GetId()),
		"email": u.GetEmail(),
	}, "", "  ")
	return string(data), nil
}
