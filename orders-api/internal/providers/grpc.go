package providers

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcConnString string

func NewGrpcClient(addr GrpcConnString) (*grpc.ClientConn, func(), error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(string(addr), opts...)
	cleanup := func() {
		conn.Close()
	}
	return conn, cleanup, err
}
