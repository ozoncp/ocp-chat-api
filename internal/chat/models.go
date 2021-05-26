package chat

import "fmt"

type Deps struct {
	Id          uint64
	ClassroomId uint64
	Link        string
}

type Chat struct {
	id          uint64
	classroomId uint64
	link        string
}

func New(deps *Deps) *Chat {
	return &Chat{
		deps.Id,
		deps.ClassroomId,
		deps.Link,
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
