package providers

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcConnString string

func NewGrpcClient(addr GrpcConnString) (*grpc.ClientConn, func(), error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	}

	conn, err := grpc.Dial(string(addr), opts...)
	cleanup := func() {
		conn.Close()
	}
	return conn, cleanup, err
}
