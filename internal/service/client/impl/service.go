package impl

import (
	"computer-club-manager/internal/model"
	"computer-club-manager/internal/model/entity"
	"computer-club-manager/internal/repository/client"
	client2 "computer-club-manager/internal/service/client"
	"computer-club-manager/internal/service/computer"
)

var _ client2.Service = (*ClientService)(nil)

type ClientService struct {
	repository      client.Repository
	computerService computer.Service
}

func NewClientService(
	repository client.Repository,
	computerService computer.Service,
) *ClientService {
	return &ClientService{
		repository:      repository,
		computerService: computerService,
	}
}

func (s *ClientService) Create(client entity.Client) error {
	return s.repository.Create(client)
}

func (s *ClientService) IsClientExists(clientId string) bool {
	_, err := s.repository.Get(clientId)
	return err == nil
}

func (s *ClientService) Delete(clientId string) error {
	return s.repository.Delete(clientId)
}

func (s *ClientService) SetWaiting(clientId string) error {
	if s.repository.CountWhereStatusIsWaiting()+1 > s.computerService.Count() {
		_ = s.Delete(clientId)
		return model.QueueIsTooLongError
	}

	return s.repository.UpdateStatus(clientId, entity.StatusWaiting)
}

func (s *ClientService) GetFirstWaiting() (entity.Client, error) {
	return s.repository.GetFirstWaiting()
}

func (s *ClientService) SetBusy(clientId string) error {
	return s.repository.UpdateStatus(clientId, entity.StatusBusy)
}
