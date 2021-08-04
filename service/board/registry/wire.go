//+build wireinject

package registry

import (
	"github.com/google/wire"

	"github.com/asim/go-micro/v3"

	"github.com/ygpark2/njro/service/board/handler"
	"github.com/ygpark2/njro/shared/config"
)

// NewContainer - create new Container
func NewContainer() *Container {

	// providerSet := wire.NewSet()

	// providerSet.add()

	// wire.NewSet()

	panic(wire.Build(
		// wire.Struct(new(*micro.Event)),
		wire.Struct(new(micro.Event)),
		// wire.Value(service.Event{}),
		// wire.Value(cfg),
		config.GetConfig,
		BuildBoardRepository,
		handler.NewBoardHandler,
		wire.Struct(new(Container), "BoardRepository", "BoardHandler"),
		// wire.Struct(new(Container), "boardHandler"),
	))
	/*
		return &Container{
			ctn: wire.Build(providerSet),
		}, nil
	*/
}
