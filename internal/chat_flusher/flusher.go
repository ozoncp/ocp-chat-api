package chat_flusher

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/ozoncp/ocp-chat-api/internal/utils"
)

type ChatRepo interface {
	GetAll() ([]*chat.Chat, error)
	RemoveByID(messageID uint64) error
	DescribeByID(messageID uint64) (string, error)
	List() (string, error)
	Add(mess *chat.Chat) error
	AddBatch(mess []*chat.Chat) error
}

//go:generate mockgen --source=./flusher.go -destination=../mocks/chat_repo/repo_mock.go -package=chat_repo

type Deps struct {
	ChunkSize      int
	ChatRepository ChatRepo
}

type ChatFlusher struct {
	chunkSize int
	chatRepo  ChatRepo
}

func NewChatFlusher(deps Deps) *ChatFlusher {
	return &ChatFlusher{
		chunkSize: deps.ChunkSize,
		chatRepo:  deps.ChatRepository,
	}
}

func (f *ChatFlusher) Flush(chats []*chat.Chat) error {
	chunks := utils.SplitMessagesListToChunks(f.chunkSize, chats...)
	fmt.Printf("num of chunks: %d\n", len(chunks))
	for _, chunk := range chunks {
		fmt.Printf("msg: %v\n", chunk)
		if err := f.chatRepo.AddBatch(chunk); err != nil {
			return errors.Wrap(err, "flush batch of chats to chat repo")
		}
	}
	return nil
}
