package main

import (
	"fmt"
	"os"

	"github.com/ozoncp/ocp-chat-api/internal/chat_service"

	"github.com/ozoncp/ocp-chat-api/internal/chat_flusher"
	"github.com/ozoncp/ocp-chat-api/internal/chat_repo"
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
		ChunkSize: 1,
	}

	// our i/o channel with chat objects
	chatQueue := chat_repo.NewRepoInMemory()

	chatQueueDeps := chat_flusher.Deps{
		ChunkSize: 1,
	}

	chatStorageFlusher := chat_flusher.NewChatFlusher(storageFlusherDeps)
	chatQueueFlusher := chat_flusher.NewChatFlusher(chatQueueDeps)

	serviceDeps := &chat_service.Deps{
		StorageRepo:    chatStorage,
		StorageFlusher: chatStorageFlusher,
		QueueRepo:      chatQueue,
		QueueFlusher:   chatQueueFlusher,
	}

	service := chat_service.New(serviceDeps)
	fmt.Printf("%+v", service)

	return nil
}
