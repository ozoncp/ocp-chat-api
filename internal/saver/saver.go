package saver

import (
	"context"
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
	Flush(ctx context.Context, repo BatchRepo, chats []*chat.Chat) error
}

//go:generate mockgen --source=./saver.go -destination=../mocks/saver/repo_mock.go -package=saver

type BatchRepo interface {
	AddBatch(ctx context.Context, chats []*chat.Chat) error
}

type Saver struct {
	capacity uint64
	flusher  Flusher

	flushPeriod time.Duration
	repo        BatchRepo
	chats       []*chat.Chat
	chatsGuard  sync.Mutex
}

type Deps struct {
	Capacity    uint64
	FlusherHere Flusher
	Repository  BatchRepo
	FlushPeriod time.Duration
	Strategy    OverflowStrategy
}

func New(capacity uint64, flusher Flusher) *Saver {
	return &Saver{
		capacity: capacity,
		flusher:  flusher,
	}
}

func (s *Saver) Save(ctx context.Context, chat *chat.Chat) error {
	s.chatsGuard.Lock()
	defer s.chatsGuard.Unlock()
	s.chats = append(s.chats, chat)
	return nil
}

func (s *Saver) Run(ctx context.Context) error {
	ticker := time.NewTicker(s.flushPeriod)
	for {
		select {
		case <-ctx.Done():
			if err := s.flusher.Flush(ctx, s.repo, s.chats); err != nil {
				return errors.Wrap(err, "flush by ctx done")
			}
		case <-ticker.C:
			if err := s.flusher.Flush(ctx, s.repo, s.chats); err != nil {
				return errors.Wrap(err, "flush by ticker")
			}
		}
	}
}
