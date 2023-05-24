package providers

import "github.com/ThreeDotsLabs/watermill"

func ProvideLogger() watermill.LoggerAdapter {
	return watermill.NewStdLogger(false, false)
}
