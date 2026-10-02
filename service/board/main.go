package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ygpark2/njro/service/board/ent"
	"github.com/ygpark2/njro/service/board/ent/proto/entpb"
	"github.com/ygpark2/njro/pkg/config"
)

func main() {
	cfg := config.GetConfig()
	port := 8081

	// 1. Initialize Ent Client
	dialect := "sqlite3"
	dsn := "file:board.db?cache=shared&_fk=1"
	if cfg.Database != nil && cfg.Database.Host != "" {
		if dbDsn, err := cfg.Database.DSN(); err == nil && dbDsn != "" {
			dsn = dbDsn
			dialect = "postgres"
		}
	}

	client, err := ent.Open(dialect, dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("failed opening connection to database")
	}
	defer client.Close()

	// 2. Auto-migration
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("failed creating schema resources")
	}

	// 3. Setup gRPC Listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to listen on port %d", port)
	}

	grpcServer := grpc.NewServer()
	boardSvc := entpb.NewBoardService(client)
	entpb.RegisterBoardServiceServer(grpcServer, boardSvc)
	reflection.Register(grpcServer)

	// 4. Graceful Shutdown
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info().Msgf("Starting Board gRPC Service on port :%d (dialect: %s)", port, dialect)
		if err := grpcServer.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			log.Fatal().Err(err).Msg("failed to serve gRPC")
		}
	}()

	<-stopCh
	log.Info().Msg("Shutting down Board gRPC server gracefully...")
	grpcServer.GracefulStop()
	log.Info().Msg("Board server stopped.")
}
