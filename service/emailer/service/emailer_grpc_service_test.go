package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	emailerPB "github.com/ygpark2/njro/service/emailer/proto/emailer"
	"github.com/ygpark2/njro/service/emailer/service"
)

func TestEmailerGRPCServer_SendEmail(t *testing.T) {
	server := service.NewEmailerGRPCServer()
	ctx := context.Background()

	// 1. Success case
	res, err := server.SendEmail(ctx, &emailerPB.SendEmailRequest{
		To:      "user@example.com",
		From:    "no-reply@njro.com",
		Subject: "회원가입 환영 메일",
		Body:    "환영합니다!",
	})
	require.NoError(t, err)
	assert.True(t, res.GetSuccess())
	assert.Contains(t, res.GetMessage(), "user@example.com")

	// 2. Failure case (empty recipient)
	_, err = server.SendEmail(ctx, &emailerPB.SendEmailRequest{
		To: "",
	})
	assert.Error(t, err)
}
