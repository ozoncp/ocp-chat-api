package saver

import (
	"context"
	"sync"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/ozoncp/ocp-chat-api/internal/chat_flusher"
	"github.com/ozoncp/ocp-chat-api/internal/utils"

	"github.com/pkg/errors"

	"github.com/ozoncp/ocp-chat-api/internal/chat"
)

type OverflowStrategy int

const (
	NotDefined OverflowStrategy = iota
	RemoveOldest
	RemoveRandom
)

//go:generate mockgen --source=./saver.go -destination=../mocks/chat_flusher/flusher_mock.go -package=chat_flusher

type Flusher interface {
	Flush(ctx context.Context, repo chat_flusher.FlushableChatRepo, chats []*chat.Chat) error
}

type BatchRepo interface {
	AddBatch(ctx context.Context, chats []*chat.Chat) error
}

type BufferingSaver struct {
	capacity uint64
	flusher  Flusher

	flushPeriod      time.Duration
	repo             chat_flusher.FlushableChatRepo
	bufferChats      []*chat.Chat
	chatsGuard       sync.Mutex
	overflowStrategy OverflowStrategy
}

type Deps struct {
	Capacity    uint64
	FlusherHere Flusher
	Repository  chat_flusher.FlushableChatRepo
	FlushPeriod time.Duration
	Strategy    OverflowStrategy
}

func New(deps *Deps) *BufferingSaver {
	return &BufferingSaver{
		capacity:         deps.Capacity,
		flusher:          deps.FlusherHere,
		repo:             deps.Repository,
		flushPeriod:      deps.FlushPeriod,
		overflowStrategy: deps.Strategy,
	}
}

func (s *BufferingSaver) Save(ctx context.Context, chat ...*chat.Chat) error {
	logger := utils.LoggerFromCtxOrCreate(ctx)
	logger.Info().Msg("saver request create multiple chats")
	secondLevelSpan, ctx := opentracing.StartSpanFromContext(ctx, "save works here")
	defer secondLevelSpan.Finish()
	s.chatsGuard.Lock()
	defer s.chatsGuard.Unlock()
	s.bufferChats = append(s.bufferChats, chat...)
	return nil
}

func (s *BufferingSaver) Run(ctx context.Context) error {
	logger := utils.LoggerFromCtxOrCreate(ctx)
	ticker := time.NewTicker(s.flushPeriod)
	for {
		select {
		case <-ctx.Done():
			logger.Debug().Msgf("flusher flushes last time, bufferchats len : %v", len(s.bufferChats))
			if err := s.flusher.Flush(ctx, s.repo, s.bufferChats); err != nil {
				return errors.Wrap(err, "flush by ctx done")
			}
			return errors.Wrap(ctx.Err(), "finish buffering saver by context done")
		case <-ticker.C:
			logger.Debug().Msgf("flusher flushes, bufferchats len : %v", len(s.bufferChats))

			firstLevelSpan, ctx := opentracing.StartSpanFromContext(ctx, "BufferingSaver tick")
			defer firstLevelSpan.Finish()

			if err := s.flusher.Flush(ctx, s.repo, s.bufferChats); err != nil {
				logger.Err(err).Msg("flush by ticker")
			}
			s.bufferChats = []*chat.Chat{}
		}
	}
}
