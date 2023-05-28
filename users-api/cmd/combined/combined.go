package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/felixge/httpsnoop"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	pb "github.com/sysradium/petproject/users-api/api/users/v1"
	"github.com/sysradium/petproject/users-api/internal/app/server"
	"github.com/sysradium/petproject/users-api/internal/storage/ephemeral"
	"github.com/uptrace/opentelemetry-go-extra/otelplay"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)

	shutdown := otelplay.ConfigureOpentelemetry(ctx)
	defer shutdown()

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)
	hdlr := server.New(
		server.WithLogger(logger),
		server.WithStorage(ephemeral.New()),
	)
	pb.RegisterUsersServiceServer(grpcServer, hdlr)
	reflection.Register(grpcServer)

	if err := pb.RegisterUsersServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:8080",
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	); err != nil {
		log.Fatal(err)
	}

	httpSrv := http.Server{
		Handler: withLogger(mux),
	}

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	m := cmux.New(l)

	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	go httpSrv.Serve(httpL)
	go grpcServer.Serve(grpcL)
	m.Serve()
}

func withLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, writer, request)
		log.Printf("http[%d]-- %s -- %s\n", m.Code, m.Duration, request.URL.Path)
	})
}
