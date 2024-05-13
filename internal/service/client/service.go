package client

import "computer-club-manager/internal/model/entity"

type Service interface {
	Create(entity.Client) error
	Delete(string) error
	IsClientExists(string) bool
	SetWaiting(string) error
	SetBusy(string) error
	GetFirstWaiting() (entity.Client, error)
}
