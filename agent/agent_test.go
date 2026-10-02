package main_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ygpark2/njro/agent/client"
	"github.com/ygpark2/njro/agent/mcp"
	"github.com/ygpark2/njro/agent/orchestrator"
)

func setupTestRegistry(t *testing.T) *mcp.ToolRegistry {
	cfg := client.DefaultConfig()
	backendClients, err := client.NewBackendClients(cfg)
	require.NoError(t, err)
	t.Cleanup(func() {
		backendClients.Close()
	})
	return mcp.NewToolRegistry(backendClients)
}

func TestToolRegistry_ToolDefinitions(t *testing.T) {
	registry := setupTestRegistry(t)
	tools := registry.ListTools()

	assert.GreaterOrEqual(t, len(tools), 10)

	expectedTools := []string{
		"search_community",
		"list_boards",
		"get_board",
		"create_board",
		"list_posts",
		"get_post",
		"create_post",
		"get_content",
		"send_email",
		"get_user",
	}

	toolMap := make(map[string]mcp.ToolDefinition)
	for _, tool := range tools {
		toolMap[tool.Name] = tool
		assert.NotEmpty(t, tool.Description)
		assert.NotNil(t, tool.InputSchema)
	}

	for _, name := range expectedTools {
		_, exists := toolMap[name]
		assert.True(t, exists, "Expected tool %s to be registered", name)
	}
}

func TestMCPServer_InitializeAndList(t *testing.T) {
	registry := setupTestRegistry(t)
	server := mcp.NewServer(registry)
	ctx := context.Background()

	// 1. Test initialize
	initReq := []byte(`{
		"jsonrpc": "2.0",
		"id": 1,
		"method": "initialize",
		"params": {
			"protocolVersion": "2024-11-05",
			"capabilities": {},
			"clientInfo": {"name": "test-client", "version": "1.0.0"}
		}
	}`)
	initResp, err := server.HandleMessage(ctx, initReq)
	require.NoError(t, err)
	require.NotNil(t, initResp)
	assert.Equal(t, float64(1), initResp.ID)

	resMap, ok := initResp.Result.(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "2024-11-05", resMap["protocolVersion"])

	// 2. Test ping
	pingReq := []byte(`{"jsonrpc": "2.0", "id": 2, "method": "ping"}`)
	pingResp, err := server.HandleMessage(ctx, pingReq)
	require.NoError(t, err)
	assert.NotNil(t, pingResp)
	assert.Equal(t, float64(2), pingResp.ID)

	// 3. Test tools/list
	listReq := []byte(`{"jsonrpc": "2.0", "id": 3, "method": "tools/list"}`)
	listResp, err := server.HandleMessage(ctx, listReq)
	require.NoError(t, err)
	assert.NotNil(t, listResp)
	assert.Nil(t, listResp.Error)

	listData, ok := listResp.Result.(map[string]any)
	require.True(t, ok)
	toolsSlice, ok := listData["tools"].([]mcp.ToolDefinition)
	require.True(t, ok)
	assert.GreaterOrEqual(t, len(toolsSlice), 10)

	// 4. Test notification (no response expected)
	notifyReq := []byte(`{"jsonrpc": "2.0", "method": "notifications/initialized"}`)
	notifyResp, err := server.HandleMessage(ctx, notifyReq)
	require.NoError(t, err)
	assert.Nil(t, notifyResp)

	// 5. Test unknown tool call
	callUnknown := []byte(`{
		"jsonrpc": "2.0",
		"id": 4,
		"method": "tools/call",
		"params": {
			"name": "non_existent_tool",
			"arguments": {}
		}
	}`)
	callResp, err := server.HandleMessage(ctx, callUnknown)
	require.NoError(t, err)
	assert.NotNil(t, callResp)
	callResult, ok := callResp.Result.(map[string]any)
	require.True(t, ok)
	assert.True(t, callResult["isError"].(bool))
}

func TestAgentOrchestrator_OfflineExecution(t *testing.T) {
	registry := setupTestRegistry(t)
	mockLLM := &orchestrator.MockLLMClient{}
	agent := orchestrator.NewAgent(mockLLM, registry)

	ctx := context.Background()

	// Test general fallback prompt
	resp, err := agent.Run(ctx, "안녕 너는 누구니?")
	require.NoError(t, err)
	assert.Contains(t, resp, "njro 마이크로서비스 백엔드를 관리하는 AI 에이전트")
}

func TestUUIDHelpers(t *testing.T) {
	testUUID := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	raw, err := client.ParseUUID(testUUID)
	require.NoError(t, err)
	assert.Len(t, raw, 16)

	str := client.UUIDToString(raw)
	assert.Equal(t, testUUID, str)

	// Invalid UUID handling
	_, err = client.ParseUUID("invalid-uuid")
	assert.Error(t, err)

	invalidStr := client.UUIDToString([]byte("too-short"))
	assert.Empty(t, invalidStr)
}
