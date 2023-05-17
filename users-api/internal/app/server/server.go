package server

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/sysradium/petproject/users-api/internal/storage"
	"github.com/sysradium/petproject/users-api/internal/storage/models"
	pb "github.com/sysradium/petproject/users-api/proto/users/v1"
)

type Server struct {
	pb.UnimplementedUsersServiceServer
	logger logrus.FieldLogger
	st     storage.Storage
}

func (s *Server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	userId, err := s.st.Create(ctx, &models.User{
		Email: req.User.GetEmail(),
	})

	if err != nil {
		s.logger.Errorf("unable to create: %+v", err)
		return nil, err
	}

	return &pb.CreateResponse{
		UserId: userId,
	}, nil
}

func (s *Server) List(ctx context.Context, _ *pb.ListRequest) (*pb.ListResponse, error) {

	users, err := s.st.List(ctx)
	if err != nil {
		s.logger.Errorf("can not get users: %+v", err)
		return nil, err
	}

	rspUsers := make([]*pb.User, len(users))

	for i, u := range users {
		rspUsers[i] = u.ToProto()
	}

	return &pb.ListResponse{
		Users: rspUsers,
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
