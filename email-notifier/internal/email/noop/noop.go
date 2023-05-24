package noop

import (
	"github.com/sysradium/petproject/email-notifier/internal/email"
)

type NoopSender struct {
}

func (n NoopSender) Send(_ email.Message) error {
	return nil
}

func New() *NoopSender {
	return &NoopSender{}
}
