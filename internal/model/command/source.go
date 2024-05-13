package command

import (
	"computer-club-manager/internal/lib/time"
	"fmt"
)

type SourceMessageType int

const (
	Enter SourceMessageType = iota
	Sit
	Wait
	Leave
)

type SourceMessage struct {
	Time time.Timestamp
	Type SourceMessageType
	Args []string
}

func (m SourceMessage) ToString() string {
	args := m.Args[0]

	if m.Type == 1 {
		args += " " + m.Args[1]
	}

	return fmt.Sprintf("%s %d %s", m.Time.ToString(), m.Type+1, args)
}
