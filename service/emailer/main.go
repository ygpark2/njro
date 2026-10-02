package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	emailerPB "github.com/ygpark2/njro/service/emailer/proto/emailer"
	"github.com/ygpark2/njro/service/emailer/service"
)

func main() {
	port := 8086

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to listen on port %d", port)
	}

	grpcServer := grpc.NewServer()
	emailerSvc := service.NewEmailerGRPCServer()
	emailerPB.RegisterEmailerServiceServer(grpcServer, emailerSvc)
	reflection.Register(grpcServer)

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info().Msgf("Starting Emailer gRPC Service on port :%d", port)
		if err := grpcServer.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			log.Fatal().Err(err).Msg("failed to serve gRPC")
		}
	}()

	<-stopCh
	log.Info().Msg("Shutting down Emailer gRPC server gracefully...")
	grpcServer.GracefulStop()
	log.Info().Msg("Emailer server stopped.")
}
