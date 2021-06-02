package message

import (
	"fmt"
	"time"
)

type Message struct {
	Timestamp time.Time
	text      string
	userID    string
	ID        string
}

func (m *Message) String() string {
	return fmt.Sprintf("%s[%v]: %v", m.ID, m.userID, m.text)
}
