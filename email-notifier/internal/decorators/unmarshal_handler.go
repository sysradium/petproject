package decorators

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/sysradium/petproject/email-notifier/internal/handlers"
	"google.golang.org/protobuf/proto"
)

func Unmarshal[T any](
	h handlers.Handler[T],
) func(msg *message.Message) error {

	return func(msg *message.Message) error {
		e := h.NewEvent().(proto.Message)
		proto.Unmarshal(msg.Payload, e)
		return h.Handle(e.(T))
	}
}
