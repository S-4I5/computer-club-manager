package impl

import (
	"computer-club-manager/internal/model"
	"computer-club-manager/internal/model/entity"
)

type ComputerRepository struct {
	computers []int
}

func NewComputerRepository(numberOfComputers int) *ComputerRepository {
	return &ComputerRepository{
		computers: make([]int, numberOfComputers),
	}
}

func (r *ComputerRepository) Count() int {
	return len(r.computers)
}

func (r *ComputerRepository) Create(computer entity.Computer) error {
	if computer.Id < 0 || computer.Id >= len(r.computers) {
		return model.IncorrectComputerNameError
	}

	if r.computers[computer.Id] != 0 {
		return model.ComputerAlreadyExistsError
	}

	r.computers[computer.Id] = computer.PricePerHour

	return nil
}

func (r *ComputerRepository) Get(computerId int) (entity.Computer, error) {
	if computerId < 0 || computerId >= len(r.computers) || r.computers[computerId] == 0 {
		return entity.Computer{}, model.ComputerDoesNotExistsError
	}

	return entity.Computer{
		Id:           computerId,
		PricePerHour: r.computers[computerId],
	}, nil
}

func (r *ComputerRepository) List() []entity.Computer {
	var computers []entity.Computer

	for id, price := range r.computers {
		computers = append(computers, entity.Computer{
			Id:           id,
			PricePerHour: price,
		})
	}

	return computers
}
