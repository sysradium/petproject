package handler

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"google.golang.org/protobuf/proto"
)

func New(h Handler) func(msg *message.Message) error {
	return func(msg *message.Message) error {
		e := h.NewEvent().(proto.Message)
		proto.Unmarshal(msg.Payload, e)
		return h.Handle(e)
	}
}

type Handler interface {
	Handle(e interface{}) error
	NewEvent() interface{}
}
