package impl

import (
	"bufio"
	"computer-club-manager/internal/lib/time"
	"computer-club-manager/internal/model/command"
	"fmt"
)

type BufferMessageReader struct {
	reader *bufio.Reader
}

func NewBufferMessageReader(read *bufio.Reader) *BufferMessageReader {
	return &BufferMessageReader{reader: read}
}

func (r *BufferMessageReader) GetMessage() (command.SourceMessage, error) {
	var curTimeString string

	_, err := fmt.Fscan(r.reader, &curTimeString)
	if err != nil {
		return command.SourceMessage{}, err
	}

	var type_ int
	_, err = fmt.Fscan(r.reader, &type_)
	if err != nil {
		return command.SourceMessage{}, err
	}
	type_--

	numberOfElements := [4]int{1, 2, 1, 1}

	args := make([]string, numberOfElements[type_])

	for i, _ := range args {
		var curArg string
		_, err = fmt.Fscan(r.reader, &curArg)
		if err != nil {
			return command.SourceMessage{}, err
		}
		args[i] = curArg
	}

	curTime, err := time.NewTimestampFromString(curTimeString)
	if err != nil {
		return command.SourceMessage{}, err
	}

	return command.SourceMessage{
		Time: curTime,
		Type: command.SourceMessageType(type_),
		Args: args,
	}, err
}
