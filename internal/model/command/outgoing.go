package command

import (
	"computer-club-manager/internal/lib/time"
	"fmt"
)

type OutgoingMessageType int

const (
	Exit OutgoingMessageType = iota
	Sat
	Error
)

type OutgoingMessage struct {
	Time    time.Timestamp
	Type    OutgoingMessageType
	Message string
}

func (m OutgoingMessage) ToString() string {
	return fmt.Sprintf("%s %d %s", m.Time.ToString(), m.Type+11, m.Message)
}
