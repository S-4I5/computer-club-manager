package impl

import (
	"computer-club-manager/internal/model/entity"
	"computer-club-manager/internal/repository/computer"
)

type ComputerService struct {
	repository computer.Repository
}

func NewComputerService(repository computer.Repository) *ComputerService {
	return &ComputerService{repository: repository}
}

func (s *ComputerService) Create(computer entity.Computer) error {
	return s.repository.Create(computer)
}

func (s *ComputerService) IsComputerExists(id int) bool {
	_, err := s.repository.Get(id)
	return err == nil
}

func (s *ComputerService) List() []entity.Computer {
	return s.repository.List()
}

func (s *ComputerService) Count() int {
	return s.repository.Count()
}

func (s *ComputerService) Get(id int) (entity.Computer, error) {
	return s.repository.Get(id)
}
