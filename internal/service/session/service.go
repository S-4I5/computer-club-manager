package session

import (
	"computer-club-manager/internal/lib/time"
	"computer-club-manager/internal/model/entity"
)

type Service interface {
	Create(entity.Session) error
	CloseSession(string, time.Timestamp) error
	CloseAllSessions(time.Timestamp) ([]string, error)
	CalculateProfitForComputer(int) (int, error)
	CalculateTimeOfUsageForComputer(int) (time.Timestamp, error)
	SetClientWaitingIfNoEmptySpot(clientId string) error
	DeleteClientWithCloseSession(id string, endTime time.Timestamp) error
	TryToCreateFromWaitList(startTime time.Timestamp) (entity.Session, error)
}
