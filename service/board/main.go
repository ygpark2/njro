package main

import (
	/*
		"github.com/asim/go-micro/v3/client"
		"github.com/asim/go-micro/v3/logger"
		"github.com/asim/go-micro/v3/server"

		"github.com/asim/go-micro/v3"
	*/

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	// tags "github.com/ygpark2/services/tags/proto"

	"github.com/ygpark2/njro/shared/config"
	"github.com/ygpark2/njro/shared/constants"

	"github.com/ygpark2/njro/service/board/registry"

	boardPB "github.com/ygpark2/njro/service/board/proto/board"
)

func main() {
	cfg := config.GetConfig()

	// Initialize Features
	// var clientWrappers []client.Wrapper
	// var handlerWrappers []server.HandlerWrapper
	// var subscriberWrappers []server.SubscriberWrapper

	// Wrappers are invoked in the order as they added
	if cfg.Features.Reqlogs.Enabled {
		// clientWrappers = append(clientWrappers, logWrapper.NewClientWrapper())
		// handlerWrappers = append(handlerWrappers, logWrapper.NewHandlerWrapper())
		// subscriberWrappers = append(subscriberWrappers, logWrapper.NewSubscriberWrapper())
	}

	logger.Debug("++++++++++++++++++++++ start auth start ++++++++++++++++++++++++++++++")
	// setupAuthForService("admin", "micro")
	logger.Debug("++++++++++++++++++++++ end auth start ++++++++++++++++++++++++++++++")

	// Create the service
	srv := service.New(
		service.Name(constants.BOARD_SERVICE),
		service.Version(config.Version),

		// Adding some optional lifecycle actions
		service.BeforeStart(func() (err error) {
			logger.Debug("called BeforeStart")
			return
		}),

		service.BeforeStop(func() (err error) {
			logger.Debug("called BeforeStop")
			return
		}),

		// micro.WrapHandler(ctn.BoardHandler),
	)

	srv.Init(
	// micro.WrapHandler(handlerWrappers...),
	// micro.WrapSubscriber(subscriberWrappers...),
	)

	// Publisher publish to "mkit.service.emailer"
	publisher := service.NewEvent(constants.EMAILER_SERVICE)
	// greeterSrv Client to call "mkit.service.greeter"
	// greeterSrvClient := greeterPB.NewGreeterService(constants.GREETER_SERVICE, srv.Client())

	logger.Debug("++++++++++++++++++++++ start NewContainer ++++++++++++++++++++++++++++++")
	// Initialize DI Container
	ctn := registry.NewContainer(publisher)

	// Register Handler
	/*
		srv.Handle(ctn.BoardHandler)
		srv.Handle(&handler.Posts{
			Tags: tags.NewTagsService("tags", srv.Client()),
		})
	*/

	boardPB.RegisterBoardServiceHandler(srv.Server(), ctn.BoardHandler)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
