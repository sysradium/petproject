package handlers

import "github.com/ThreeDotsLabs/watermill/message"

type EventHandlers map[string]func(msg *message.Message) error

type Handler[T any] interface {
	Handle(e T) error
	NewEvent() interface{}
}
