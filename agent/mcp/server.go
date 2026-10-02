package mcp

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// JSONRPCRequest represents an incoming JSON-RPC 2.0 request.
type JSONRPCRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      any             `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

// JSONRPCResponse represents an outgoing JSON-RPC 2.0 response.
type JSONRPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      any             `json:"id,omitempty"`
	Result  any             `json:"result,omitempty"`
	Error   *JSONRPCError   `json:"error,omitempty"`
}

// JSONRPCError represents a JSON-RPC 2.0 error.
type JSONRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Server implements a Model Context Protocol (MCP) server.
type Server struct {
	registry *ToolRegistry
	logger   *log.Logger
}

// NewServer creates a new MCP server.
func NewServer(registry *ToolRegistry) *Server {
	return &Server{
		registry: registry,
		logger:   log.New(os.Stderr, "[MCP Server] ", log.LstdFlags),
	}
}

// HandleMessage processes a single JSON-RPC message and returns the response (if not a notification).
func (s *Server) HandleMessage(ctx context.Context, data []byte) (*JSONRPCResponse, error) {
	var req JSONRPCRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return &JSONRPCResponse{
			JSONRPC: "2.0",
			Error: &JSONRPCError{
				Code:    -32700,
				Message: fmt.Sprintf("Parse error: %v", err),
			},
		}, nil
	}

	// Notifications have no ID and do not return responses
	isNotification := req.ID == nil

	resp := &JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      req.ID,
	}

	switch req.Method {
	case "initialize":
		resp.Result = map[string]any{
			"protocolVersion": "2024-11-05",
			"capabilities": map[string]any{
				"tools": map[string]any{},
			},
			"serverInfo": map[string]any{
				"name":    "njro-agent-mcp",
				"version": "1.0.0",
			},
		}

	case "notifications/initialized":
		return nil, nil // Notification, no response

	case "ping":
		resp.Result = map[string]any{}

	case "tools/list":
		resp.Result = map[string]any{
			"tools": s.registry.ListTools(),
		}

	case "tools/call":
		var callParams struct {
			Name      string         `json:"name"`
			Arguments map[string]any `json:"arguments"`
		}
		if err := json.Unmarshal(req.Params, &callParams); err != nil {
			resp.Error = &JSONRPCError{
				Code:    -32602,
				Message: fmt.Sprintf("Invalid params: %v", err),
			}
			break
		}

		output, err := s.registry.Call(ctx, callParams.Name, callParams.Arguments)
		if err != nil {
			resp.Result = map[string]any{
				"content": []map[string]any{
					{
						"type": "text",
						"text": fmt.Sprintf("Error executing tool %s: %v", callParams.Name, err),
					},
				},
				"isError": true,
			}
		} else {
			resp.Result = map[string]any{
				"content": []map[string]any{
					{
						"type": "text",
						"text": output,
					},
				},
				"isError": false,
			}
		}

	default:
		if isNotification {
			return nil, nil
		}
		resp.Error = &JSONRPCError{
			Code:    -32601,
			Message: fmt.Sprintf("Method not found: %s", req.Method),
		}
	}

	if isNotification {
		return nil, nil
	}
	return resp, nil
}

// RunStdio starts listening for MCP JSON-RPC requests on reader and responding on writer.
func (s *Server) RunStdio(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	// Buffer up to 10MB per line for large payloads
	buf := make([]byte, 1024*1024)
	scanner.Buffer(buf, 10*1024*1024)

	s.logger.Println("MCP Server started in stdio mode (ready for Cursor / Claude Desktop)...")

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		resp, err := s.HandleMessage(context.Background(), line)
		if err != nil {
			s.logger.Printf("Error processing message: %v\n", err)
			continue
		}
		if resp == nil {
			continue // Notification
		}

		respBytes, err := json.Marshal(resp)
		if err != nil {
			s.logger.Printf("Failed to marshal response: %v\n", err)
			continue
		}

		if _, err := w.Write(append(respBytes, '\n')); err != nil {
			return fmt.Errorf("failed to write response: %w", err)
		}
	}

	return scanner.Err()
}
