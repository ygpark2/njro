package main

import (
	"time"

	"github.com/asim/go-micro/v3/auth"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/server"
	"github.com/google/uuid"

	"github.com/asim/go-micro/v3"

	// tags "github.com/ygpark2/services/tags/proto"

	"github.com/ygpark2/mboard/shared/config"
	"github.com/ygpark2/mboard/shared/constants"

	logWrapper "github.com/ygpark2/mboard/shared/wrapper/log"

	"github.com/ygpark2/mboard/service/board/registry"
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

		micro.WrapHandler(ctn.BoardHandler),
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

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}

// setupAuthForService generates auth credentials for the service
func setupAuthForService(accID string, accSecret string) error {
	// opts := auth.DefaultAuth.Options()

	// extract the account creds from options, these can be set by flags
	// accID := ID
	// accSecret := Secret

	// if no credentials were provided, self generate an account
	if len(accID) == 0 || len(accSecret) == 0 {
		opts := []auth.GenerateOption{
			auth.WithType("service"),
			auth.WithScopes("service"),
		}

		acc, err := auth.Generate(uuid.New().String(), opts...)
		if err != nil {
			return err
		}
		if logger.V(logger.DebugLevel, logger.DefaultLogger) {
			logger.Debugf("Auth [%v] Generated an auth account", auth.DefaultAuth.String())
		}

		accID = acc.ID
		accSecret = acc.Secret
	}

	// generate the first token
	token, err := auth.Token(
		auth.WithCredentials(accID, accSecret),
		auth.WithExpiry(time.Minute*10),
	)
	if err != nil {
		return err
	}

	// set the credentials and token in auth options
	auth.DefaultAuth.Init(
		auth.ClientToken(token),
		auth.Credentials(accID, accSecret),
	)
	return nil
}
