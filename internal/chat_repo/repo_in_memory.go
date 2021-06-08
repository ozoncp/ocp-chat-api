package chat_repo

import (
	"fmt"
	"strings"

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

func (c *RepoInMemory) GetAll() ([]*chat.Chat, error) {
	return c.chats, nil
}

func (c *RepoInMemory) Add(ch *chat.Chat) error {
	c.chats = append(c.chats, ch)
	return nil
}

func (c *RepoInMemory) AddBatch(mess []*chat.Chat) error {
	c.chats = append(c.chats, mess...)
	return nil
}

func (c *RepoInMemory) RemoveByID(chatID uint64) error {
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

func (c *RepoInMemory) DescribeByID(chatID uint64) (string, error) {
	for n, ch := range c.chats {
		if ch.ID == chatID {
			return fmt.Sprintf("order_num %d, message %+v", n, ch), nil
		}
	}
	return "", ErrChatNotFound
}

func (c *RepoInMemory) List() (string, error) {
	list := make([]string, 0, len(c.chats))
	for _, ch := range c.chats {
		list = append(list, ch.String())
	}
	return strings.Join(list, "\n"), nil
}