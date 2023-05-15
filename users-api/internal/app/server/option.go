package server

import "github.com/sirupsen/logrus"

type Option func(*Server) error

func WithLogger(l logrus.FieldLogger) Option {
	return func(s *Server) error {
		s.logger = l
		return nil
	}
}
