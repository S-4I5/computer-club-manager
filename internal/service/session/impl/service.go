package impl

import (
	"computer-club-manager/internal/lib/time"
	"computer-club-manager/internal/model"
	"computer-club-manager/internal/model/entity"
	"computer-club-manager/internal/repository/session"
	"computer-club-manager/internal/service/client"
	"computer-club-manager/internal/service/computer"
)

type SessionService struct {
	repository      session.Repository
	computerService computer.Service
	clientService   client.Service
}

func NewSessionService(
	repository session.Repository,
	computerService computer.Service,
	clientService client.Service,
) *SessionService {
	return &SessionService{
		repository:      repository,
		computerService: computerService,
		clientService:   clientService,
	}
}

func (s *SessionService) Create(session entity.Session) error {
	if !s.computerService.IsComputerExists(session.Computer.Id) {
		return model.ComputerDoesNotExistsError
	}

	if !s.clientService.IsClientExists(session.Client.Id) {
		return model.ClientDoesNotExistsError
	}

	_, err := s.repository.GetOpenedByClient(session.Client.Id)
	if err == nil {
		err := s.CloseSession(session.Client.Id, session.StartedAt)
		if err != nil {
			return err
		}
	}

	return s.repository.Create(session)
}

func (s *SessionService) SetClientWaitingIfNoEmptySpot(clientId string) error {
	if s.computerService.Count() > s.repository.CountOpenedSessions() {
		return model.FreeComputerAvailableError
	}

	return s.clientService.SetWaiting(clientId)
}

func (s *SessionService) DeleteClientWithCloseSession(id string, endTime time.Timestamp) error {
	_ = s.CloseSession(id, endTime)

	err := s.clientService.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SessionService) TryToCreateFromWaitList(startTime time.Timestamp) (entity.Session, error) {
	waiting, err := s.clientService.GetFirstWaiting()
	if err != nil {
		return entity.Session{}, err
	}

	var availableComputer entity.Computer

	allComputers := s.computerService.List()
	for _, curComputer := range allComputers {
		busy, _ := s.repository.IsSessionOpenForComputer(curComputer.Id)
		if !busy {
			availableComputer, _ = s.computerService.Get(curComputer.Id)
			break
		}
	}

	newSession := entity.Session{
		Computer: availableComputer,
		Client: entity.Client{
			Id:     waiting.Id,
			Status: entity.StatusBusy,
		},
		StartedAt: startTime,
		EndedAt:   time.Timestamp{},
	}

	err = s.Create(newSession)
	if err != nil {
		return entity.Session{}, err
	}

	err = s.clientService.SetBusy(waiting.Id)
	if err != nil {
		return entity.Session{}, err
	}
	return newSession, nil
}

func (s *SessionService) CloseSession(id string, endTime time.Timestamp) error {

	err := s.repository.CloseSession(id, endTime)
	if err != nil {
		return err
	}

	return nil
}

func (s *SessionService) CloseAllSessions(endTime time.Timestamp) ([]string, error) {
	return s.repository.CloseAllSessions(endTime)
}

func (s *SessionService) CalculateProfitForComputer(id int) (int, error) {
	curComputer, err := s.computerService.Get(id)
	if err != nil {
		return -1, err
	}

	sessions := s.repository.GetAllClosedSessionsForComputer(id)

	var profit int

	for _, curSession := range sessions {
		profit += curComputer.PricePerHour * curSession.GetLengthInHoursRounded()
	}

	return profit, nil
}

func (s *SessionService) CalculateTimeOfUsageForComputer(id int) (time.Timestamp, error) {
	exists := s.computerService.IsComputerExists(id)
	if !exists {
		return time.Timestamp{}, model.ComputerDoesNotExistsError
	}
	sessions := s.repository.GetAllClosedSessionsForComputer(id)

	var timeSum int

	for _, curSession := range sessions {
		timeSum += curSession.GetLengthInMinutes()
	}

	timeStamp, err := time.NewTimestamp(timeSum)
	if err != nil {
		return time.Timestamp{}, err
	}

	return timeStamp, nil
}
