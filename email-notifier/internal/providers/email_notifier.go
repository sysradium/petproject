package providers

import (
	"github.com/sysradium/petproject/email-notifier/internal/email"
	"github.com/sysradium/petproject/email-notifier/internal/email/noop"
)

func ProvideEmailNotifier() email.Sender {
	return noop.New()
}
