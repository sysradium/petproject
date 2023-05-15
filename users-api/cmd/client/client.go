package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/sysradium/petproject/users-api/proto/users/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:9090",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("unable to dial server: %v", err)
	}

	defer conn.Close()
	c := pb.NewUsersServiceClient(
		conn,
	)

	rsp, err := c.List(context.Background(), &pb.ListRequest{})
	if err != nil {
		log.Fatalf("unable to get users: %+v", err)
	}
	fmt.Println(rsp.Users)
}
