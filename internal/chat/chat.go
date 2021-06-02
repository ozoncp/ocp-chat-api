package chat

import (
	"fmt"
	"os"

	"github.com/ozoncp/ocp-chat-api/internal/message"
)

type MessageRepo interface {
	GetMessages() []message.Message
	RemoveMessageById(messageID string) error
	DescribeMessageById(messageID string) (string, error)
	ListMessages() string
	AddMessage(mess *message.Message)
}

type Deps struct {
	Id          uint64
	ClassroomId uint64
	Link        string
	Messages    MessageRepo
}

type Chat struct {
	id          uint64
	classroomId uint64
	link        string
	messages    MessageRepo
}

func New(deps *Deps) *Chat {
	return &Chat{
		id:          deps.Id,
		classroomId: deps.ClassroomId,
		link:        deps.Link,
	}
}

func (c *Chat) ID() uint64 {
	return c.id
}

func (c *Chat) String() string {
	return fmt.Sprintf("Chat. { ID: %v, classroomID: %v, link: %v", c.id, c.classroomId, c.link)
}

func (c *Chat) SetClassroomID(classroomID uint64) {
	c.classroomId = classroomID
}

func (c *Chat) ClassroomID() uint64 {
	return c.classroomId
}

func (c *Chat) SetLink(link string) {
	c.link = link
}

func (c *Chat) Link() string {
	return c.link
}

func (c *Chat) GetMessages() []message.Message {
	return c.messages.GetMessages()
}

func (c *Chat) AddMessage(mess *message.Message) {
	c.messages.AddMessage(mess)
	_, _ = fmt.Fprintf(os.Stderr, "add mess %v \n", mess)
}

func (c *Chat) RemoveMessageById(messageID string) error {
	return c.messages.RemoveMessageById(messageID)
}

func (c *Chat) DescribeMessageById(messageID string) (string, error) {
	return c.messages.DescribeMessageById(messageID)
}

func (c *Chat) ListMessages() string {
	return c.messages.ListMessages()
}
