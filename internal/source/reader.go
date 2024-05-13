package source

import "computer-club-manager/internal/model/command"

type MessageSource interface {
	GetMessage() (command.SourceMessage, error)
}
