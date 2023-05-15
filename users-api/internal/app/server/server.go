package server

import (
	"context"

	"github.com/sirupsen/logrus"
	pb "github.com/sysradium/petproject/users-api/proto/users/v1"
)

type Server struct {
	pb.UnimplementedUsersServiceServer
	logger logrus.FieldLogger
}

func (s *Server) List(_ context.Context, _ *pb.ListRequest) (*pb.ListResponse, error) {
	panic("not implemented") // TODO: Implement
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
