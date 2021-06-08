package chat_flusher

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/ozoncp/ocp-chat-api/internal/utils"
)

//go:generate mockgen --source=./flusher.go -destination=../mocks/chat_repo/flushable_repo_mock.go -package=chat_repo

type FlushableChatRepo interface {
	AddBatch(chats []*chat.Chat) error
}

type Deps struct {
	ChunkSize int
}

type ChatFlusher struct {
	chunkSize int
}

func NewChatFlusher(deps Deps) *ChatFlusher {
	return &ChatFlusher{
		chunkSize: deps.ChunkSize,
	}
}

func (f *ChatFlusher) Flush(repo FlushableChatRepo, chats []*chat.Chat) error {
	chunks := utils.SplitMessagesListToChunks(f.chunkSize, chats...)
	fmt.Printf("num of chunks: %d\n", len(chunks))
	for _, chunk := range chunks {
		fmt.Printf("msg: %v\n", chunk)
		if err := repo.AddBatch(chunk); err != nil {
			return errors.Wrap(err, "flush batch of chats to chat repo")
		}
	}
	return nil
}
