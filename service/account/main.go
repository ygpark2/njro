package main

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/server"

	"github.com/rs/zerolog/log"

	"github.com/ygpark2/mboard/shared/config"
	"github.com/ygpark2/mboard/shared/constants"

	logWrapper "github.com/ygpark2/mboard/shared/wrapper/log"
	validatorWrapper "github.com/ygpark2/mboard/shared/wrapper/validator"

	"github.com/ygpark2/mboard/service/account/handler"
	"github.com/ygpark2/mboard/service/account/registry"
	"github.com/ygpark2/mboard/service/account/repository"

	profilePB "github.com/ygpark2/mboard/service/account/proto/profile"
	userPB "github.com/ygpark2/mboard/service/account/proto/user"
	greeterPB "github.com/ygpark2/mboard/service/greeter/proto/greeter"
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

	/*
		if cfg.Features.Translogs.Enabled {
			topic := cfg.Features.Translogs.Topic
			publisher := micro.NewEvent(topic, client.DefaultClient) // service.Client())
			handlerWrappers = append(handlerWrappers, transWrapper.NewHandlerWrapper(publisher))
			subscriberWrappers = append(subscriberWrappers, transWrapper.NewSubscriberWrapper(publisher))
		}
	*/

	if cfg.Features.Validator.Enabled {
		handlerWrappers = append(handlerWrappers, validatorWrapper.NewHandlerWrapper())
		subscriberWrappers = append(subscriberWrappers, validatorWrapper.NewSubscriberWrapper())
	}

	service := micro.NewService(
		micro.Name(constants.ACCOUNT_SERVICE),
		micro.Version(config.Version),
		// myMicro.WithTLS(),
		// Wrappers are applied in reverse order so the last is executed first.
		// service.WrapClient(clientWrappers...),
		// Adding some optional lifecycle actions
		micro.BeforeStart(func() (err error) {
			log.Debug().Msg("called BeforeStart")
			return
		}),
		micro.BeforeStop(func() (err error) {
			log.Debug().Msg("called BeforeStop")
			return
		}),
	)

	service.Init(
		micro.WrapHandler(handlerWrappers...),
		micro.WrapSubscriber(subscriberWrappers...),
	)

	// Initialize DI Container
	ctn, err := registry.NewContainer(cfg)
	defer ctn.Clean()
	if err != nil {
		log.Fatal().Msgf("failed to build container: %v", err)
	}

	// Publisher publish to "mkit.service.emailer"
	publisher := micro.NewEvent(constants.EMAILER_SERVICE, service.Client())

	// greeterSrv Client to call "mkit.service.greeter"
	greeterSrvClient := greeterPB.NewGreeterService(constants.GREETER_SERVICE, service.Client())

	// // Handlers
	userHandler := handler.NewUserHandler(ctn.Resolve("user-repository").(repository.UserRepository), publisher, greeterSrvClient)
	profileHandler := ctn.Resolve("profile-handler").(profilePB.ProfileServiceHandler)

	// Register Handlers
	userPB.RegisterUserServiceHandler(service.Server(), userHandler)
	profilePB.RegisterProfileServiceHandler(service.Server(), profileHandler)

	println(config.GetBuildInfo())

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
