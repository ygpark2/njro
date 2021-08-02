package main

import (
	"github.com/micro/micro/v3/service/client"
	"github.com/micro/micro/v3/service/server"
	"github.com/rs/zerolog/log"

	transactionPB "github.com/ygpark2/mboard/service/recorder/proto/transaction"
	"github.com/ygpark2/mboard/service/recorder/registry"
	"github.com/ygpark2/mboard/shared/config"
	"github.com/ygpark2/mboard/shared/constants"
	// myMicro "github.com/ygpark2/mboard/shared/util/micro"
	logWrapper "github.com/ygpark2/mboard/shared/wrapper/log"
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
	if cfg.Features.Validator.Enabled {
		handlerWrappers = append(handlerWrappers, validatorWrapper.NewHandlerWrapper())
		subscriberWrappers = append(subscriberWrappers, validatorWrapper.NewSubscriberWrapper())
	}

	srv := service.NewService(
		service.Name(constants.RECORDER_SERVICE),
		service.Version(config.Version),
		// myMicro.WithTLS(),
		// Wrappers are applied in reverse order so the last is executed first.
		service.WrapClient(clientWrappers...),
		service.WrapHandler(handlerWrappers...),
		service.WrapSubscriber(subscriberWrappers...),
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

	srv.Init()

	// Initialize DI Container
	ctn, err := registry.NewContainer(cfg)
	defer ctn.Clean()
	if err != nil {
		log.Fatal().Msgf("failed to build container: %v", err)
	}

	transactionSubscriber := ctn.Resolve("transaction-subscriber") //.(subscriber.TransactionSubscriber)
	recorderTopic := cfg.Features.Translogs.Topic

	// Register Struct as Subscriber
	_ = service.RegisterSubscriber(recorderTopic, service.Server(), transactionSubscriber)

	// register subscriber with queue, each message is delivered to a unique subscriber
	// _ = micro.RegisterSubscriber(recorderTopic, service.Server(), transactionSubscriber, server.SubscriberQueue("queue.pubsub"))

	transactionHandler := ctn.Resolve("transaction-handler").(transactionPB.TransactionServiceHandler)
	transactionPB.RegisterTransactionServiceHandler(service.Server(), transactionHandler)

	println(config.GetBuildInfo())

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
