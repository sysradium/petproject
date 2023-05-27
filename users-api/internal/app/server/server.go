package server

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	pb "github.com/sysradium/petproject/users-api/api/users/v1"
	"github.com/sysradium/petproject/users-api/internal/storage"
	"github.com/sysradium/petproject/users-api/internal/storage/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedUsersServiceServer
	logger logrus.FieldLogger
	st     storage.Storage
}

func (s *Server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	u, err := s.st.Create(ctx, &models.User{
		Email: req.User.GetEmail(),
	})

	if err != nil {
		s.logger.Errorf("unable to create: %+v", err)
		return nil, err
	}

	return &pb.CreateResponse{
		UserId: u.Id.String(),
	}, nil
}

func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	rsp, err := s.st.Get(ctx, uuid.MustParse(req.Id))
	if err == nil {
		return &pb.GetResponse{
			User: rsp.ToProto(),
		}, nil
	}

	if errors.Is(err, storage.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return nil, err
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
