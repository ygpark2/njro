package main

import (
	"github.com/micro/micro/v3/service/client"
	"github.com/micro/micro/v3/service/server"
	"github.com/rs/zerolog/log"

	"github.com/ygpark2/mboard/service/greeter/handler"
	greeterPB "github.com/ygpark2/mboard/service/greeter/proto/greeter"
	healthPB "github.com/ygpark2/mboard/service/greeter/proto/health"
	"github.com/ygpark2/mboard/shared/config"
	"github.com/ygpark2/mboard/shared/constants"
	// myMicro "github.com/ygpark2/mboard/shared/util/micro"
	logWrapper "github.com/ygpark2/mboard/shared/wrapper/log"
	// transWrapper "github.com/ygpark2/mboard/shared/wrapper/transaction"
	validatorWrapper "github.com/ygpark2/mboard/shared/wrapper/validator"
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
	//if cfg.Features.Translogs.Enabled {
	//    topic := cfg.Features.Translogs.Topic
	//    publisher := micro.NewEvent(topic, client.DefaultClient) // service.Client())
	//    handlerWrappers = append(handlerWrappers, transWrapper.NewHandlerWrapper(publisher))
	//    subscriberWrappers = append(subscriberWrappers, transWrapper.NewSubscriberWrapper(publisher))
	//}
	if cfg.Features.Validator.Enabled {
		handlerWrappers = append(handlerWrappers, validatorWrapper.NewHandlerWrapper())
		subscriberWrappers = append(subscriberWrappers, validatorWrapper.NewSubscriberWrapper())
	}

	srv := service.New(
		service.Name(constants.GREETER_SERVICE),
		service.Version(config.Version),
		// myMicro.WithTLS(),
		// Wrappers are applied in reverse order so the last is executed first.
		service.WrapClient(clientWrappers...),
		// Adding some optional lifecycle actions
		service.BeforeStart(func() (err error) {
			log.Debug().Msg("called BeforeStart")
			return
		}),
		service.BeforeStop(func() (err error) {
			log.Debug().Msg("called BeforeStop")
			return
		}),
	)

	if cfg.Features.Translogs.Enabled {
		/*
		topic := cfg.Features.Translogs.Topic
		publisher := service.NewEvent(topic, service.Client())
		handlerWrappers = append(handlerWrappers, transWrapper.NewHandlerWrapper(publisher))
		subscriberWrappers = append(subscriberWrappers, transWrapper.NewSubscriberWrapper(publisher))
		*/
	}

	service.Init(
		service.WrapHandler(handlerWrappers...),
		service.WrapSubscriber(subscriberWrappers...),
	)

	_ = healthPB.RegisterHealthHandler(service.Server(), handler.NewHealthHandler())
	_ = greeterPB.RegisterGreeterServiceHandler(service.Server(), handler.NewGreeterHandler())

	println(config.GetBuildInfo())

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
