package impl

import (
	"computer-club-manager/internal/model"
	"computer-club-manager/internal/model/entity"
	"computer-club-manager/internal/repository/client"
)

var _ client.Repository = (*ClientRepository)(nil)

type ClientRepository struct {
	waitingClients []string
	clients        map[string]entity.ClientStatus
}

func NewClientRepository() *ClientRepository {
	return &ClientRepository{
		waitingClients: []string{},
		clients:        make(map[string]entity.ClientStatus),
	}
}

func (r *ClientRepository) Create(client entity.Client) error {
	_, isPresent := r.clients[client.Id]
	if isPresent {
		return model.ClientAlreadyExistsError
	}

	r.clients[client.Id] = client.Status

	return nil
}

func (r *ClientRepository) Get(clientName string) (entity.Client, error) {
	status, isPresent := r.clients[clientName]
	if !isPresent {
		return entity.Client{}, model.ClientDoesNotExistsError
	}

	return entity.Client{Id: clientName, Status: status}, nil
}

func (r *ClientRepository) Delete(clientName string) error {
	curStatus, isPresent := r.clients[clientName]
	if !isPresent {
		return model.ClientDoesNotExistsError
	}

	if curStatus == entity.StatusWaiting {
		for i, id := range r.waitingClients {
			if id == clientName {
				r.waitingClients = append(r.waitingClients[:i], r.waitingClients[i+1:]...)
				break
			}
		}
	}

	delete(r.clients, clientName)

	return nil
}

func (r *ClientRepository) CountWhereStatusIsWaiting() int {
	return len(r.waitingClients)
}

func (r *ClientRepository) GetFirstWaiting() (entity.Client, error) {
	if len(r.waitingClients) == 0 {
		return entity.Client{}, model.NoOneWaitingError
	}

	return entity.Client{Id: r.waitingClients[0], Status: entity.StatusWaiting}, nil
}

func (r *ClientRepository) UpdateStatus(clientName string, status entity.ClientStatus) error {
	curStatus, isPresent := r.clients[clientName]
	if !isPresent {
		return model.ClientDoesNotExistsError
	}

	if status == entity.StatusWaiting && curStatus != entity.StatusWaiting {
		r.waitingClients = append(r.waitingClients, clientName)
	} else if curStatus == entity.StatusWaiting {
		r.waitingClients = r.waitingClients[1:]
	}

	r.clients[clientName] = status

	return nil
}
