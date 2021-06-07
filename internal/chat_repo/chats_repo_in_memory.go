package chat_repo

import (
	"fmt"
	"strings"

	"github.com/ozoncp/ocp-chat-api/internal/chat"
	"github.com/pkg/errors"
)

var ErrChatNotFound = errors.New("no chat with this ID")

type ChatsRepoInMemory struct {
	chats []*chat.Chat
}

func NewChatsRepoInMemory() *ChatsRepoInMemory {
	return &ChatsRepoInMemory{}
}

func (c *ChatsRepoInMemory) GetAll() []*chat.Chat {
	return c.chats
}

func (c *ChatsRepoInMemory) Add(ch *chat.Chat) {
	c.chats = append(c.chats, ch)
}

func (c *ChatsRepoInMemory) AddBatch(mess []*chat.Chat) {
	c.chats = append(c.chats, mess...)
}

func (c *ChatsRepoInMemory) RemoveByID(chatID uint64) error {
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

func (c *ChatsRepoInMemory) DescribeByID(chatID uint64) (string, error) {
	for n, ch := range c.chats {
		if ch.ID == chatID {
			return fmt.Sprintf("order_num %d, message %+v", n, ch), nil
		}
	}
	return "", ErrChatNotFound
}

func (c *ChatsRepoInMemory) List() string {
	list := make([]string, 0, len(c.chats))
	for _, ch := range c.chats {
		list = append(list, ch.String())
	}
	return strings.Join(list, "\n")
}
