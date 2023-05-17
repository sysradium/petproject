package server

import (
	"github.com/sirupsen/logrus"
	"github.com/sysradium/petproject/users-api/internal/storage"
)

type Option func(*Server) error

func WithLogger(l logrus.FieldLogger) Option {
	return func(s *Server) error {
		s.logger = l
		return nil
	}
}

func WithStorage(st storage.Storage) Option {
	return func(s *Server) error {
		s.st = st
		return nil
	}
}
