package chat

import (
	"fmt"
)

type Deps struct {
	Id          uint64
	ClassroomId uint64
	Link        string
}

type Chat struct {
	ID          uint64
	ClassroomID uint64
	Link        string
}

func New(deps *Deps) *Chat {
	return &Chat{
		ID:          deps.Id,
		ClassroomID: deps.ClassroomId,
		Link:        deps.Link,
	}
}

func (c *Chat) String() string {
	return fmt.Sprintf("Chat. { ID: %v, classroomID: %v, Link: %v", c.ID, c.ClassroomID, c.Link)
}
