package sender

import "computer-club-manager/internal/model/command"

type MessageSender interface {
	SendOutgoingMessage(message command.OutgoingMessage) error
	SendSourceMessage(message command.SourceMessage) error
	Send(message interface{}) error
}
