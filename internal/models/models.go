package models

type Chat struct {
	link string
}

func (c *Chat) String() string {
	return "Chat. Link: " + c.link
}
