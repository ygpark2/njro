package handler

import (
	"context"

	"github.com/micro/micro/v3/service/logger"

	contentPB "github.com/ygpark2/njro/service/content/proto/content"
)

type ContentHandler struct{}

// Call is a single request handler called via client.Call or the generated client code
func (c *ContentHandler) Save(ctx context.Context, req *contentPB.Request, rsp *contentPB.Response) error {
	logger.Info("Not yet implemented")
	return nil
}
