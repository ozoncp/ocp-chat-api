package chat_repo

import (
	"context"
	"github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/pkg/errors"
)

var ErrChatNotFound = errors.New("no chat with this ID")

type RepoInMemory struct {
	chats []*chat.Chat
}

func NewRepoInMemory() *RepoInMemory {
	return &RepoInMemory{}
}

func (c *RepoInMemory) GetAll(ctx context.Context) ([]*chat.Chat, error) {
	return c.chats, nil
}

func (c *RepoInMemory) Insert(ctx context.Context, classroomID uint64, link string) (*chat.Chat, error) {
	ch := &chat.Chat{
		ID:          uint64(len(c.chats)),
		ClassroomID: classroomID,
		Link:        link,
	}
	c.chats = append(c.chats, ch)
	return ch, nil
}

func (c *RepoInMemory) Remove(ctx context.Context, chatID uint64) error {
	for n, ch := range c.chats {
		if ch.ID == chatID {
			if n == len(c.chats)-1 {
				c.chats = c.chats[:n]
			}
			c.chats = append(c.chats[:n], c.chats[n+1:]...)
			return nil
		}
	}
	return ErrChatNotFound
}

func (c *RepoInMemory) Describe(ctx context.Context, chatID uint64) (*chat.Chat, error) {
	for _, ch := range c.chats {
		if ch.ID == chatID {
			return ch, nil
		}
	}
	return nil, ErrChatNotFound
}

func (c *RepoInMemory) AddBatch(ctx context.Context, chats []*chat.Chat) error {
	for _, ch := range chats {
		if _, err := c.Insert(ctx, ch.ClassroomID, ch.Link); err != nil {
			return errors.Wrap(err, "insert to inmemory db")
		}
	}
	return nil
}
