package main

import (
	"context"
	"io"
	"log"

	pb "github.com/sysradium/petproject/users-api/proto/server/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:3000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("unable to dial server: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterServiceClient(
		conn,
	)

	stream, err := c.PullMessages(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("unable to open stream: %+v", err)
	}

	for {
		rsp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		log.Printf("rsp: %v", rsp.Msg)
	}

}
