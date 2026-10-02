package entpb_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ygpark2/njro/service/board/ent/enttest"
	"github.com/ygpark2/njro/service/board/ent/proto/entpb"
)

func TestBoardService_CRUD(t *testing.T) {
	// 1. In-memory SQLite client via enttest
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	// 2. Initialize generated gRPC BoardService
	svc := entpb.NewBoardService(client)
	ctx := context.Background()

	// 3. Test Create
	created, err := svc.Create(ctx, &entpb.CreateBoardRequest{
		Board: &entpb.Board{
			Title:       "개발자 자유게시판",
			MobileTitle: "자유게시판",
			Order:       1,
			Search:      true,
			Description: "개발자들을 위한 자유로운 소통 공간",
			Notices:     []string{"공지사항 1", "공지사항 2"},
		},
	})
	require.NoError(t, err)
	assert.NotEmpty(t, created.GetId())
	assert.Equal(t, "개발자 자유게시판", created.GetTitle())
	assert.Equal(t, "자유게시판", created.GetMobileTitle())
	assert.Equal(t, uint32(1), created.GetOrder())
	assert.True(t, created.GetSearch())
	assert.Len(t, created.GetNotices(), 2)

	// 4. Test Get
	fetched, err := svc.Get(ctx, &entpb.GetBoardRequest{
		Id: created.GetId(),
	})
	require.NoError(t, err)
	assert.Equal(t, created.GetId(), fetched.GetId())
	assert.Equal(t, "개발자 자유게시판", fetched.GetTitle())

	// 5. Test Update
	updated, err := svc.Update(ctx, &entpb.UpdateBoardRequest{
		Board: &entpb.Board{
			Id:          created.GetId(),
			Title:       "공식 개발자 자유게시판 (수정됨)",
			MobileTitle: "자유게시판",
		},
	})
	require.NoError(t, err)
	assert.Equal(t, "공식 개발자 자유게시판 (수정됨)", updated.GetTitle())

	// 6. Test List
	listed, err := svc.List(ctx, &entpb.ListBoardRequest{
		PageSize: 10,
	})
	require.NoError(t, err)
	assert.Len(t, listed.GetBoardList(), 1)
	assert.Equal(t, "공식 개발자 자유게시판 (수정됨)", listed.GetBoardList()[0].GetTitle())

	// 7. Test Delete
	_, err = svc.Delete(ctx, &entpb.DeleteBoardRequest{
		Id: created.GetId(),
	})
	require.NoError(t, err)

	// Verify Deleted
	_, err = svc.Get(ctx, &entpb.GetBoardRequest{
		Id: created.GetId(),
	})
	require.Error(t, err)
}
