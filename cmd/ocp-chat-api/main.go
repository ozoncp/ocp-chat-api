package main

import (
	"fmt"
	"os"

	chat2 "github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/ozoncp/ocp-chat-api/internal/message_repo"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var defaultLogger = log.Logger.With().Timestamp().Logger()

func main() {
	if err := Run(); err != nil {
		defaultLogger.Fatal().Err(err).Msg("run application")
	}
}

func Run() error {
	defaultLogger.Info().Msg("Hi, Victor Akhlynin will write this project")

	for i := 0; i < 100; i++ {
		f, err := os.Open("go.mod")
		if err != nil {
			return errors.Wrap(err, "open file")
		}
		defaultLogger.Info().Msg("open successful")
		defer func() {
			if err := f.Close(); err != nil {
				defaultLogger.Error().Err(err).Msg("close file bad")
			}
		}()
	}

	messageRepo := message_repo.NewMessageRepoInMemory()

	chatDeps := &chat2.Deps{
		Id:          1,
		ClassroomId: 1,
		Link:        "http://chatmychatwhereisit.com",
		Messages:    messageRepo,
	}

	chat := chat2.New(chatDeps)

	fmt.Printf("%+v it's chat", chat)

	return nil
}
