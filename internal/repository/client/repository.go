package client

import "computer-club-manager/internal/model/entity"

type Repository interface {
	Create(client entity.Client) error
	Get(clientName string) (entity.Client, error)
	Delete(clientName string) error
	CountWhereStatusIsWaiting() int
	GetFirstWaiting() (entity.Client, error)
	UpdateStatus(clientName string, status entity.ClientStatus) error
}
