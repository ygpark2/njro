package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	"github.com/ygpark2/njro/service/tags/handler"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("tags"),
	)

	// Register Handler
	srv.Handle(new(handler.Tags))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
