package impl

import (
	"computer-club-manager/internal/lib/time"
	"computer-club-manager/internal/model"
	"computer-club-manager/internal/model/entity"
	"computer-club-manager/internal/repository/session"
)

var _ session.Repository = (*SessionRepository)(nil)

type SessionRepository struct {
	computerBusy              []bool
	openSessionsForClient     map[string]entity.Session
	closedSessionsForComputer [][]entity.Session
}

func NewSessionRepository(numberOfComputers int) *SessionRepository {
	return &SessionRepository{
		computerBusy:              make([]bool, numberOfComputers),
		openSessionsForClient:     make(map[string]entity.Session),
		closedSessionsForComputer: make([][]entity.Session, numberOfComputers),
	}
}

func (s *SessionRepository) Create(session entity.Session) error {
	if s.computerBusy[session.Computer.Id] != false {
		return model.ComputerAlreadyOccupiedError
	}

	s.computerBusy[session.Computer.Id] = true
	s.openSessionsForClient[session.Client.Id] = session

	return nil
}

func (s *SessionRepository) GetOpenedByClient(clientName string) (entity.Session, error) {
	curSession, isOpened := s.openSessionsForClient[clientName]
	if !isOpened {
		return entity.Session{}, model.SessionDoesNotExistsError
	}

	return curSession, nil
}

func (s *SessionRepository) CountOpenedSessions() int {
	return len(s.openSessionsForClient)
}

func (s *SessionRepository) CloseSession(clientName string, endTime time.Timestamp) error {
	curSession, isOpened := s.openSessionsForClient[clientName]
	if !isOpened {
		return model.SessionDoesNotExistsError
	}
	delete(s.openSessionsForClient, clientName)

	curSession.EndedAt = endTime

	s.computerBusy[curSession.Computer.Id] = false

	s.closedSessionsForComputer[curSession.Computer.Id] =
		append(s.closedSessionsForComputer[curSession.Computer.Id], curSession)

	return nil
}

func (s *SessionRepository) IsSessionOpenForComputer(id int) (bool, error) {
	if id < 0 || id >= len(s.computerBusy) {
		return false, model.IncorrectComputerNameError
	}

	return s.computerBusy[id], nil
}

func (s *SessionRepository) CloseAllSessions(endTime time.Timestamp) ([]string, error) {
	var closedFor []string

	for key, _ := range s.openSessionsForClient {
		err := s.CloseSession(key, endTime)
		if err != nil {
			return []string{}, err
		}
		closedFor = append(closedFor, key)
	}
	return closedFor, nil
}

func (s *SessionRepository) ListAllClosed() []entity.Session {
	var sessionList []entity.Session
	for _, sessions := range s.closedSessionsForComputer {
		for _, curSession := range sessions {
			sessionList = append(sessionList, curSession)
		}
	}

	return sessionList
}

func (s *SessionRepository) GetAllClosedSessionsForComputer(id int) []entity.Session {
	//id--
	if id < 0 || id >= len(s.closedSessionsForComputer) {
		return []entity.Session{}
	}

	return s.closedSessionsForComputer[id]
}
