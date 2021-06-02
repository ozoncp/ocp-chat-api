package flusher

import (
	"github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/ozoncp/ocp-chat-api/internal/message"
	"github.com/ozoncp/ocp-chat-api/internal/utils"
)

type Flusher struct {
	chunkSize   int
	messageRepo chat.MessageRepo
}

func (f *Flusher) Flush(messages []message.Message) {
	chunks := utils.SplitMessagesListToChunks(f.chunkSize, messages...)
	for _, chunk := range chunks {
		for _, m := range chunk {
			f.messageRepo.AddMessage(&m)
		}
	}
}
