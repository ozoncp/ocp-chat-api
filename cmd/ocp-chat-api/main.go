package main

import (
	"fmt"
	"os"

	"github.com/ozoncp/ocp-chat-api/internal/chat_flusher"

	"github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/ozoncp/ocp-chat-api/internal/chat_repo"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var defaultLogger = log.Logger.With().Timestamp().Logger()

type Flusher interface {
	Flush(chats []*chat.Chat) error
}

func main() {
	if err := Run(); err != nil {
		defaultLogger.Fatal().Err(err).Msg("run application")
	}
}

//go:generate mockgen --source=./main.go -destination=../../internal/mocks/chat_flusher/flusher_mock.go -package=chat_flusher

func Run() error {
	defaultLogger.Info().Msg("Hi, Victor Akhlynin will write this project")

	for i := 0; i < 5; i++ {
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

	// our future persistent DB
	chatStorage := chat_repo.NewRepoInMemory()

	storageFlusherDeps := chat_flusher.Deps{
		ChunkSize:      1,
		ChatRepository: chatStorage,
	}

	// our i/o channel with chat objects
	chatQueue := chat_repo.NewRepoInMemory()

	chatQueueDeps := chat_flusher.Deps{
		ChunkSize:      1,
		ChatRepository: chatQueue,
	}

	chatStorageFlusher := chat_flusher.NewChatFlusher(storageFlusherDeps)
	chatQueueFlusher := chat_flusher.NewChatFlusher(chatQueueDeps)

	fmt.Printf("%+v, %+v", chatStorageFlusher, chatQueueFlusher)

	return nil
}
