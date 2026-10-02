package entpb_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ygpark2/njro/service/content/ent/enttest"
	"github.com/ygpark2/njro/service/content/ent/proto/entpb"
)

func TestContentService_CRUD(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	svc := entpb.NewContentService(client)
	ctx := context.Background()

	// 1. Create
	created, err := svc.Create(ctx, &entpb.CreateContentRequest{
		Content: &entpb.Content{
			BoardId: "board-123",
			PostId:  "post-456",
			Content: "이것은 게시글 본문 내용입니다.",
		},
	})
	require.NoError(t, err)
	assert.NotEmpty(t, created.GetId())
	assert.Equal(t, "board-123", created.GetBoardId())
	assert.Equal(t, "이것은 게시글 본문 내용입니다.", created.GetContent())

	// 2. Get
	fetched, err := svc.Get(ctx, &entpb.GetContentRequest{
		Id: created.GetId(),
	})
	require.NoError(t, err)
	assert.Equal(t, created.GetId(), fetched.GetId())

	// 3. Update
	updated, err := svc.Update(ctx, &entpb.UpdateContentRequest{
		Content: &entpb.Content{
			Id:      created.GetId(),
			Content: "수정된 게시글 본문 내용입니다.",
		},
	})
	require.NoError(t, err)
	assert.Equal(t, "수정된 게시글 본문 내용입니다.", updated.GetContent())

	// 4. Delete
	_, err = svc.Delete(ctx, &entpb.DeleteContentRequest{
		Id: created.GetId(),
	})
	require.NoError(t, err)
}
