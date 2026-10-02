package handler_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ygpark2/njro/service/search/handler"
	searchPB "github.com/ygpark2/njro/service/search/proto"
)

func TestSearchGRPCServer(t *testing.T) {
	srv := handler.NewSearchGRPCServer()
	ctx := context.Background()

	// 1. Index documents
	_, err := srv.Index(ctx, &searchPB.IndexRequest{
		Document: &searchPB.Document{
			Id:   "doc-1",
			Text: "Go Micro를 이용한 AI 에이전트 개발 가이드",
		},
	})
	require.NoError(t, err)

	_, err = srv.Index(ctx, &searchPB.IndexRequest{
		Document: &searchPB.Document{
			Id:   "doc-2",
			Text: "Ent와 Protobuf 기반의 현대적 마이크로서비스 설계",
		},
	})
	require.NoError(t, err)

	// 2. Search keyword "Ent"
	res, err := srv.Search(ctx, &searchPB.SearchRequest{
		Keyword: "Ent",
	})
	require.NoError(t, err)
	assert.Len(t, res.GetDocuments(), 1)
	assert.Equal(t, "doc-2", res.GetDocuments()[0].GetId())

	// 3. Search keyword "마이크로서비스"
	res, err = srv.Search(ctx, &searchPB.SearchRequest{
		Keyword: "마이크로서비스",
	})
	require.NoError(t, err)
	assert.Len(t, res.GetDocuments(), 1)

	// 4. Search keyword not found
	res, err = srv.Search(ctx, &searchPB.SearchRequest{
		Keyword: "없는키워드",
	})
	require.NoError(t, err)
	assert.Empty(t, res.GetDocuments())
}
