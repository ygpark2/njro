package entpb_test

import (
	"context"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/ygpark2/njro/service/account/ent/enttest"
	"github.com/ygpark2/njro/service/account/ent/proto/entpb"
	"github.com/ygpark2/njro/service/account/ent/user"
)

func TestUserService_CRUD_And_PasswordHook(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	svc := entpb.NewUserService(client)
	ctx := context.Background()

	rawPassword := "my-secret-password-1234"

	// 1. Create User via gRPC Service
	created, err := svc.Create(ctx, &entpb.CreateUserRequest{
		User: &entpb.User{
			Username:  "youngpark",
			FirstName: "Young",
			LastName:  "Park",
			Email:     "youngpark@example.com",
			Password:  wrapperspb.String(rawPassword),
		},
	})
	require.NoError(t, err)
	assert.NotEmpty(t, created.GetId())
	assert.Equal(t, "youngpark", created.GetUsername())

	// 2. Verify Password Hook: Check database entity directly
	dbUser, err := client.User.Query().Where(user.Username("youngpark")).Only(ctx)
	require.NoError(t, err)

	// DB에 저장된 비밀번호는 평문이 아니라 bcrypt 해시여야 함
	assert.NotEqual(t, rawPassword, dbUser.Password)
	assert.True(t, strings.HasPrefix(dbUser.Password, "$2a$"), "Password should be bcrypt hashed")

	// bcrypt 검증 통과 여부 확인
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(rawPassword))
	assert.NoError(t, err, "Password hash must match raw password")

	// 3. Get User via gRPC
	fetched, err := svc.Get(ctx, &entpb.GetUserRequest{
		Id: created.GetId(),
	})
	require.NoError(t, err)
	assert.Equal(t, created.GetId(), fetched.GetId())
	assert.Equal(t, "youngpark", fetched.GetUsername())

	// 4. Update User Profile
	updated, err := svc.Update(ctx, &entpb.UpdateUserRequest{
		User: &entpb.User{
			Id:        created.GetId(),
			FirstName: "YoungUpdated",
			LastName:  "ParkUpdated",
		},
	})
	require.NoError(t, err)
	assert.Equal(t, "YoungUpdated", updated.GetFirstName())

	// 5. Delete User
	_, err = svc.Delete(ctx, &entpb.DeleteUserRequest{
		Id: created.GetId(),
	})
	require.NoError(t, err)
}
