package providers

import "github.com/ThreeDotsLabs/watermill"

func NewLogger() watermill.LoggerAdapter {
	return watermill.NewStdLogger(false, false)
}
