package computer

import (
	"computer-club-manager/internal/model/entity"
)

type Service interface {
	Create(entity.Computer) error
	List() []entity.Computer
	Get(int) (entity.Computer, error)
	IsComputerExists(int) bool
	Count() int
}
