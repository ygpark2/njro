package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	"github.com/ygpark2/njro/shared/config"
	"github.com/ygpark2/njro/shared/constants"

	"github.com/ygpark2/njro/service/board/registry"

	commentPB "github.com/ygpark2/njro/service/comment/proto/comment"
)

func main() {
	// Create the service
	srv := service.New(
		service.Name(constants.COMMENT_SERVICE),
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

	commentPB.RegisterBoardServiceHandler(srv.Server(), ctn.CommentHandler)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
