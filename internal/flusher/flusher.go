package flusher

import (
	"fmt"

	"github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/ozoncp/ocp-chat-api/internal/utils"
)

type ChatRepo interface {
	GetAll() []*chat.Chat
	RemoveByID(messageID uint64) error
	DescribeByID(messageID uint64) (string, error)
	List() string
	Add(mess *chat.Chat)
	AddBatch(mess []*chat.Chat)
}

//go:generate mockgen --source=./flusher.go -destination=../mocks/chat_repo/chat_repo_mock.go -package=chat_repo

type FlusherChats struct {
	chunkSize int
	chatRepo  ChatRepo
}

type Deps struct {
	ChunkSize      int
	ChatRepository ChatRepo
}

func NewFlusherMessagesToChat(deps Deps) *FlusherChats {
	return &FlusherChats{
		chunkSize: deps.ChunkSize,
		chatRepo:  deps.ChatRepository,
	}
}

func (f *FlusherChats) Flush(messages []*chat.Chat) {
	chunks := utils.SplitMessagesListToChunks(f.chunkSize, messages...)
	fmt.Printf("num of chunks: %d\n", len(chunks))
	for _, chunk := range chunks {
		fmt.Printf("msg: %v\n", chunk)
		f.chatRepo.AddBatch(chunk)
	}
}
