package message_repo

import (
	"fmt"
	"strings"

	"github.com/ozoncp/ocp-chat-api/internal/message"
	"github.com/pkg/errors"
)

var ErrMessageNotFound = errors.New("no message with this ID")

type MessageRepoInMemory struct {
	messages []message.Message
}

func NewMessageRepoInMemory() *MessageRepoInMemory {
	return &MessageRepoInMemory{}
}

func (c *MessageRepoInMemory) GetMessages() []message.Message {
	return c.messages
}

func (c *MessageRepoInMemory) RemoveMessageById(messageID string) error {
	for n, mess := range c.messages {
		if mess.ID == messageID {
			if n == len(c.messages)-1 {
				c.messages = c.messages[:n]
			}
			c.messages = append(c.messages[:n], c.messages[n+1:]...)
			return nil
		}
	}
	return ErrMessageNotFound
}

func (c *MessageRepoInMemory) DescribeMessageById(messageID string) (string, error) {
	for n, mess := range c.messages {
		if mess.ID == messageID {
			return fmt.Sprintf("order_num %d, message %+v", n, mess), nil
		}
	}
	return "", ErrMessageNotFound
}

func (c *MessageRepoInMemory) ListMessages() string {
	list := make([]string, 0, len(c.messages))
	for _, mess := range c.messages {
		list = append(list, mess.String())
	}
	return strings.Join(list, "\n")
}
