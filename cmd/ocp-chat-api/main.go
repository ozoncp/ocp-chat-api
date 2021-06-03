package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ozoncp/ocp-chat-api/internal/flusher"
	"github.com/ozoncp/ocp-chat-api/internal/message"

	chat2 "github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/ozoncp/ocp-chat-api/internal/message_repo"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var defaultLogger = log.Logger.With().Timestamp().Logger()

type Flusher interface {
	Flush(messages []message.Message)
}

func main() {
	if err := Run(); err != nil {
		defaultLogger.Fatal().Err(err).Msg("run application")
	}
}

//go:generate mockgen --source=./main.go -destination=../../internal/mocks/flusher/flusher_mock.go -package=flusher

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

	fmt.Printf("%+v it's chat\n", chat)

	m1 := message.Message{
		Timestamp: time.Time{},
		ID:        "asdfsdf",
	}

	m2 := message.Message{
		Timestamp: time.Time{},
		ID:        "r2r2fda",
	}

	messageList := []message.Message{m1, m2}

	flusherDeps := flusher.Deps{
		ChunkSize:         1,
		MessageRepository: messageRepo,
	}

	//nolint:gosimple // not finished application, ok
	var myFlusher Flusher
	myFlusher = flusher.NewFlusherMessagesToChat(flusherDeps)
	myFlusher.Flush(messageList)
	fmt.Printf("%+v finished", myFlusher)

	return nil
}
