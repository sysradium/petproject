package server

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	pb "github.com/sysradium/petproject/users-api/proto/server/v1"
)

type Server struct {
	pb.UnimplementedGreeterServiceServer
	logger logrus.FieldLogger
}

func (s *Server) PullMessages(_ *pb.PullMessagesRequest, ch pb.GreeterService_PullMessagesServer) error {
	t := time.NewTicker(time.Second)
	defer t.Stop()

	i := 0
	for range t.C {
		s.logger.Debugf("sending message: %d", i)
		msg := &pb.GreeterServicePullMessagesResponse{Msg: fmt.Sprintf("hi: %d", i)}
		if err := ch.Send(msg); err != nil {
			return err
		}
		i++
	}

	return nil
}

func (s *Server) SayHello(ctx context.Context, in *pb.GreeterServiceSayHelloRequest) (*pb.GreeterServiceSayHelloResponse, error) {
	return &pb.GreeterServiceSayHelloResponse{
		Rsp: in.Name,
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
