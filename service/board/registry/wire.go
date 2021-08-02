//+build wireinject

package registry

import (
	"github.com/google/wire"

	"github.com/micro/micro/v3/service"

	"github.com/ygpark2/mboard/service/board/handler"
	"github.com/ygpark2/mboard/shared/config"
)

// NewContainer - create new Container
func NewContainer() *Container {

	// providerSet := wire.NewSet()

	// providerSet.add()

	// wire.NewSet()

	panic(wire.Build(
		wire.Struct(new(service.Event)),
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
