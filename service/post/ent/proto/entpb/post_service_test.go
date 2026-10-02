package entpb_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ygpark2/njro/service/post/ent/enttest"
	"github.com/ygpark2/njro/service/post/ent/proto/entpb"
)

func TestPostService_CRUD(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	svc := entpb.NewPostService(client)
	ctx := context.Background()

	// 1. Create
	created, err := svc.Create(ctx, &entpb.CreatePostRequest{
		Post: &entpb.Post{
			BoardId:  "board-abc",
			Title:    "Go Micro에서 Ent로의 전환기",
			Slug:     "go-micro-to-ent",
			Category: "tech",
			Content:  "Ent와 entproto를 사용하여 마이크로서비스를 개편했습니다.",
			Email:    "test@example.com",
			Writer:   "youngpark",
			Tags:     []string{"golang", "ent", "grpc"},
		},
	})
	require.NoError(t, err)
	assert.NotEmpty(t, created.GetId())
	assert.Equal(t, "Go Micro에서 Ent로의 전환기", created.GetTitle())
	assert.Equal(t, "board-abc", created.GetBoardId())
	assert.Len(t, created.GetTags(), 3)

	// 2. Get
	fetched, err := svc.Get(ctx, &entpb.GetPostRequest{
		Id: created.GetId(),
	})
	require.NoError(t, err)
	assert.Equal(t, created.GetId(), fetched.GetId())
	assert.Equal(t, "youngpark", fetched.GetWriter())

	// 3. Update
	updated, err := svc.Update(ctx, &entpb.UpdatePostRequest{
		Post: &entpb.Post{
			Id:    created.GetId(),
			Title: "Go Micro에서 Ent로의 전환기 (최종 완료)",
		},
	})
	require.NoError(t, err)
	assert.Equal(t, "Go Micro에서 Ent로의 전환기 (최종 완료)", updated.GetTitle())

	// 4. List
	listed, err := svc.List(ctx, &entpb.ListPostRequest{
		PageSize: 10,
	})
	require.NoError(t, err)
	assert.Len(t, listed.GetPostList(), 1)

	// 5. Delete
	_, err = svc.Delete(ctx, &entpb.DeletePostRequest{
		Id: created.GetId(),
	})
	require.NoError(t, err)

	// Verify Deleted
	_, err = svc.Get(ctx, &entpb.GetPostRequest{
		Id: created.GetId(),
	})
	require.Error(t, err)
}
