package main

import (
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/server"

	"github.com/asim/go-micro/v3"

	// tags "github.com/ygpark2/services/tags/proto"

	"github.com/ygpark2/mboard/shared/config"
	"github.com/ygpark2/mboard/shared/constants"

	logWrapper "github.com/ygpark2/mboard/shared/wrapper/log"

	"github.com/ygpark2/mboard/service/board/registry"

	boardPB "github.com/ygpark2/mboard/service/board/proto/board"
)

func main() {
	cfg := config.GetConfig()

	// Initialize Features
	var clientWrappers []client.Wrapper
	var handlerWrappers []server.HandlerWrapper
	var subscriberWrappers []server.SubscriberWrapper

	// Wrappers are invoked in the order as they added
	if cfg.Features.Reqlogs.Enabled {
		clientWrappers = append(clientWrappers, logWrapper.NewClientWrapper())
		handlerWrappers = append(handlerWrappers, logWrapper.NewHandlerWrapper())
		subscriberWrappers = append(subscriberWrappers, logWrapper.NewSubscriberWrapper())
	}

	logger.Debug("++++++++++++++++++++++ start NewContainer ++++++++++++++++++++++++++++++")
	// Initialize DI Container
	ctn := registry.NewContainer()

	logger.Debug("++++++++++++++++++++++ start auth start ++++++++++++++++++++++++++++++")
	// setupAuthForService("admin", "micro")
	logger.Debug("++++++++++++++++++++++ end auth start ++++++++++++++++++++++++++++++")

	// Create the service
	service := micro.NewService(
		micro.Name(constants.BOARD_SERVICE),
		micro.Version(config.Version),

		// Adding some optional lifecycle actions
		micro.BeforeStart(func() (err error) {
			logger.Debug("called BeforeStart")
			return
		}),

		micro.BeforeStop(func() (err error) {
			logger.Debug("called BeforeStop")
			return
		}),

		// micro.WrapHandler(ctn.BoardHandler),
	)

	service.Init(
		micro.WrapHandler(handlerWrappers...),
		micro.WrapSubscriber(subscriberWrappers...),
	)

	// Publisher publish to "mkit.service.emailer"
	// publisher := service.NewEvent(constants.EMAILER_SERVICE)
	// greeterSrv Client to call "mkit.service.greeter"
	// greeterSrvClient := greeterPB.NewGreeterService(constants.GREETER_SERVICE, srv.Client())

	// Register Handler
	/*
		srv.Handle(ctn.BoardHandler)
		srv.Handle(&handler.Posts{
			Tags: tags.NewTagsService("tags", srv.Client()),
		})
	*/

	boardPB.RegisterBoardServiceHandler(service.Server(), ctn.BoardHandler)

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
