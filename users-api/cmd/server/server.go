package main

import (
	"log"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/sirupsen/logrus"
	"github.com/sysradium/petproject/users-api/internal/app/server"
	"github.com/sysradium/petproject/users-api/internal/storage/ephemeral"
	pb "github.com/sysradium/petproject/users-api/proto/users/v1"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("unable to start server: %v", err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)

	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}

	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(InterceptorLogger(logger), opts...),
		),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(InterceptorLogger(logger), opts...),
		),
	)

	pb.RegisterUsersServiceServer(srv, server.New(
		server.WithLogger(logger),
		server.WithStorage(ephemeral.New()),
	))

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("unable to start server: %v", err)
	}
}
