//go:build wireinject
// +build wireinject

package main

import (
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/wire"
	"github.com/sysradium/petproject/orders-api/internal/adapters/ephemeral"
	"github.com/sysradium/petproject/orders-api/internal/app"
	"github.com/sysradium/petproject/orders-api/internal/app/server"
	"github.com/sysradium/petproject/orders-api/internal/domain/order"
	"github.com/sysradium/petproject/orders-api/internal/ports"
	"github.com/sysradium/petproject/orders-api/internal/providers"
	pbUsers "github.com/sysradium/petproject/users-api/api/users/v1"
	"google.golang.org/grpc"
)

func Initialize(
	addr providers.GrpcConnString,
	kafkaAddr providers.KafkaAddress,
) (*server.Server, func(), error) {
	wire.Build(
		providers.NewGrpcClient,
		providers.NewLogger,
		wire.Bind(new(grpc.ClientConnInterface), new(*grpc.ClientConn)),
		wire.Bind(new(order.Repository), new(*ephemeral.Ephemeral)),
		wire.Bind(new(app.Publisher), new(*cqrs.EventBus)),
		wire.Bind(new(server.Runner), new(*message.Router)),
		providers.NewRouter,
		providers.NewEventHandlers,
		pbUsers.NewUsersServiceClient,
		ports.NewHttpServer,
		providers.NewCQRSFacade,
		ephemeral.New,
		app.NewApplication,
		providers.NewEcho,
		server.New,
	)
	return &server.Server{}, func() {}, nil
}
