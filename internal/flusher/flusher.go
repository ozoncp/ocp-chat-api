package flusher

import (
	"fmt"

	"github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/ozoncp/ocp-chat-api/internal/message"
	"github.com/ozoncp/ocp-chat-api/internal/utils"
)

type FlusherMessagesToChat struct {
	chunkSize   int
	messageRepo chat.MessageRepo
}

type Deps struct {
	ChunkSize         int
	MessageRepository chat.MessageRepo
}

func NewFlusherMessagesToChat(deps Deps) *FlusherMessagesToChat {
	return &FlusherMessagesToChat{
		chunkSize:   deps.ChunkSize,
		messageRepo: deps.MessageRepository,
	}
}

func (f *FlusherMessagesToChat) Flush(messages []*message.Message) {
	chunks := utils.SplitMessagesListToChunks(f.chunkSize, messages...)
	fmt.Printf("num of chunks: %d\n", len(chunks))
	for _, chunk := range chunks {
		fmt.Printf("msg: %v\n", chunk)
		f.messageRepo.AddMessagesBatch(chunk)
	}
}
