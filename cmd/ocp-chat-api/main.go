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

	chatRepo := chat_repo.NewRepoInMemory()

	chatDeps1 := &chat.Deps{
		Id:          1,
		ClassroomId: 11,
		Link:        "http://chat1.com",
	}
	c1 := chat.New(chatDeps1)

	chatDeps2 := &chat.Deps{
		Id:          2,
		ClassroomId: 22,
		Link:        "http://chat2.com",
	}
	c2 := chat.New(chatDeps2)

	chatDep3 := &chat.Deps{
		Id:          3,
		ClassroomId: 33,
		Link:        "http://chat3.com",
	}
	c3 := chat.New(chatDep3)

	chatList := []*chat.Chat{c1, c2, c3}

	flusherDeps := chat_flusher.Deps{
		ChunkSize:      1,
		ChatRepository: chatRepo,
	}

	//nolint:gosimple // not finished application, ok
	var myFlusher Flusher
	myFlusher = chat_flusher.NewChatFlusher(flusherDeps)
	if err := myFlusher.Flush(chatList); err != nil {
		return errors.Wrap(err, "flush chats to chat list")
	}
	fmt.Printf("%+v finished", myFlusher)

	return nil
}
