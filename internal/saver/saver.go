package saver

import (
	"context"
	"github.com/ozoncp/ocp-chat-api/internal/chat_flusher"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/ozoncp/ocp-chat-api/internal/chat"
)

type OverflowStrategy int

const (
	NotDefined OverflowStrategy = iota
	RemoveOldest
	RemoveRandom
)

//go:generate mockgen --source=./saver.go -destination=../mocks/saver/flusher_mock.go -package=saver

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

func (s *BufferingSaver) Save(ctx context.Context, chat *chat.Chat) error {
	s.chatsGuard.Lock()
	defer s.chatsGuard.Unlock()
	s.bufferChats = append(s.bufferChats, chat)
	return nil
}

func (s *BufferingSaver) Run(ctx context.Context) error {
	ticker := time.NewTicker(s.flushPeriod)
	for {
		select {
		case <-ctx.Done():
			if err := s.flusher.Flush(ctx, s.repo, s.bufferChats); err != nil {
				return errors.Wrap(err, "flush by ctx done")
			}
		case <-ticker.C:
			if err := s.flusher.Flush(ctx, s.repo, s.bufferChats); err != nil {
				return errors.Wrap(err, "flush by ticker")
			}
		}
	}
}
