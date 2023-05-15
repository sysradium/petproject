package server

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	pb "github.com/sysradium/petproject/users-api/proto/users/v1"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	pb.UnimplementedUsersServiceServer
	logger logrus.FieldLogger
}

func (s *Server) List(_ context.Context, _ *pb.ListRequest) (*pb.ListResponse, error) {

	return &pb.ListResponse{
		Users: []*pb.User{{
			Username:     "foo",
			Email:        "s@s.com",
			RegisteredAt: timestamppb.New(time.Now()),
		}, {
			Username: "someone",
			Email:    "at@gmail.com",
		}},
	}, nil
}

func New(opts ...Option) *Server {
	s := &Server{
		logger: logrus.New(),
	}

	for _, o := range opts {
		o(s)
	}

	return s
}
