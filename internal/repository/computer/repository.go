package computer

import "computer-club-manager/internal/model/entity"

type Repository interface {
	Create(entity.Computer) error
	Get(int) (entity.Computer, error)
	List() []entity.Computer
	Count() int
}
