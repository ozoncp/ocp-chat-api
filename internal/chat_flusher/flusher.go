package chat_flusher

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/ozoncp/ocp-chat-api/internal/utils"
)

//go:generate mockgen --source=./flusher.go -destination=../mocks/chat_repo/flushable_repo_mock.go -package=chat_repo

type FlushableChatRepo interface {
	AddBatch(ctx context.Context, chats []*chat.Chat) error
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

func (f *ChatFlusher) Flush(ctx context.Context, repo FlushableChatRepo, chats []*chat.Chat) error {
	logger := utils.LoggerFromCtxOrCreate(ctx)
	chunks := utils.SplitChatsListToChunks(f.chunkSize, chats...)
	logger.Debug().Int("num_chunks", len(chunks)).Msg("flush")
	for _, chunk := range chunks {
		logger.Debug().Msgf("msg: %v\n", chunk)
		if err := repo.AddBatch(ctx, chunk); err != nil {
			return errors.Wrap(err, "flush batch of chats to chat repo")
		}
	}
	return nil
}
