package utils

import (
	"context"
	"github.com/rs/zerolog"
	"os"
)

func LoggerFromCtxOrCreate(ctx context.Context) *zerolog.Logger {
	logger := zerolog.Ctx(ctx)
	if logger != nil {
		l := zerolog.New(os.Stderr).With().Timestamp().Logger()
		logger = &l
	}
	return logger
}
