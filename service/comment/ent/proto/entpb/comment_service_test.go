package entpb_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ygpark2/njro/service/comment/ent/enttest"
	"github.com/ygpark2/njro/service/comment/ent/proto/entpb"
)

func TestCommentService_CRUD(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	svc := entpb.NewCommentService(client)
	ctx := context.Background()

	// 1. Create
	created, err := svc.Create(ctx, &entpb.CreateCommentRequest{
		Comment: &entpb.Comment{
			BoardId:  "board-1",
			PostId:   "post-1",
			Userid:   "user-1",
			Username: "길동홍",
			Nickname: "홍길동",
			Email:    "hong@test.com",
			Content:  "좋은 글 잘 읽었습니다!",
		},
	})
	require.NoError(t, err)
	assert.NotEmpty(t, created.GetId())
	assert.Equal(t, "좋은 글 잘 읽었습니다!", created.GetContent())
	assert.Equal(t, "홍길동", created.GetNickname())

	// 2. Get
	fetched, err := svc.Get(ctx, &entpb.GetCommentRequest{
		Id: created.GetId(),
	})
	require.NoError(t, err)
	assert.Equal(t, created.GetId(), fetched.GetId())

	// 3. Update
	updated, err := svc.Update(ctx, &entpb.UpdateCommentRequest{
		Comment: &entpb.Comment{
			Id:      created.GetId(),
			Content: "수정된 댓글 내용입니다.",
		},
	})
	require.NoError(t, err)
	assert.Equal(t, "수정된 댓글 내용입니다.", updated.GetContent())

	// 4. List
	listed, err := svc.List(ctx, &entpb.ListCommentRequest{
		PageSize: 10,
	})
	require.NoError(t, err)
	assert.Len(t, listed.GetCommentList(), 1)

	// 5. Delete
	_, err = svc.Delete(ctx, &entpb.DeleteCommentRequest{
		Id: created.GetId(),
	})
	require.NoError(t, err)

	// Verify Deleted
	_, err = svc.Get(ctx, &entpb.GetCommentRequest{
		Id: created.GetId(),
	})
	require.Error(t, err)
}
