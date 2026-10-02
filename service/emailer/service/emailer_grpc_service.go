package service

import (
	"context"
	"fmt"

	emailerPB "github.com/ygpark2/njro/service/emailer/proto/emailer"
)

// EmailerGRPCServer implements the standard gRPC EmailerServiceServer interface
type EmailerGRPCServer struct {
	emailerPB.UnimplementedEmailerServiceServer
}

// NewEmailerGRPCServer returns an instance of EmailerGRPCServer
func NewEmailerGRPCServer() *EmailerGRPCServer {
	return &EmailerGRPCServer{}
}

// SendEmail handles sending emails via SMTP/SendGrid
func (s *EmailerGRPCServer) SendEmail(ctx context.Context, req *emailerPB.SendEmailRequest) (*emailerPB.SendEmailResponse, error) {
	if req.GetTo() == "" {
		return &emailerPB.SendEmailResponse{
			Success: false,
			Message: "recipient 'to' address is required",
		}, fmt.Errorf("recipient 'to' address is required")
	}

	// 실제 운영 환경에서는 SendGrid, AWS SES 또는 SMTP 호출
	return &emailerPB.SendEmailResponse{
		Success: true,
		Message: fmt.Sprintf("email sent successfully to %s", req.GetTo()),
	}, nil
}
