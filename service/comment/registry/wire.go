//+build wireinject

package registry

import (
	"github.com/google/wire"

	"github.com/micro/micro/v3"
	// "github.com/asim/go-micro/v3/client"

	"github.com/ygpark2/njro/service/comment/handler"
	"github.com/ygpark2/njro/shared/config"
	// "github.com/ygpark2/njro/shared/constants"
)

// NewContainer - create new Container
func NewContainer(publisher *service.Event) *Container {

	panic(wire.Build(
		config.GetConfig,
		BuildCommentRepository,
		handler.NewCommentHandler,
		wire.Struct(new(Container), "CommentRepository", "CommentHandler"),
	))
	/*
		return &Container{
			ctn: wire.Build(providerSet),
		}, nil
	*/
}
