package session

import (
	"computer-club-manager/internal/lib/time"
	"computer-club-manager/internal/model/entity"
)

type Repository interface {
	Create(entity.Session) error
	CloseSession(string, time.Timestamp) error
	IsSessionOpenForComputer(int) (bool, error)
	CloseAllSessions(time.Timestamp) ([]string, error)
	GetOpenedByClient(string) (entity.Session, error)
	CountOpenedSessions() int
	ListAllClosed() []entity.Session
	GetAllClosedSessionsForComputer(int) []entity.Session
}
