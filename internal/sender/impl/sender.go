package impl

import (
	"bufio"
	"computer-club-manager/internal/model/command"
	"fmt"
)

type BufferMessageSender struct {
	writer *bufio.Writer
}

func NewBufferMessageSender(writer *bufio.Writer) BufferMessageSender {
	return BufferMessageSender{writer: writer}
}

func (s BufferMessageSender) Send(message interface{}) error {

	_, err := fmt.Fprintln(s.writer, message)
	if err != nil {
		return err
	}

	return nil
}

func (s BufferMessageSender) SendOutgoingMessage(message command.OutgoingMessage) error {
	_, err := fmt.Fprintln(s.writer, message.ToString())
	return err
}

func (s BufferMessageSender) SendSourceMessage(message command.SourceMessage) error {
	_, err := fmt.Fprintln(s.writer, message.ToString())
	return err
}
