package entity

import "fmt"

type ClientStatus int

const (
	StatusWaiting ClientStatus = iota
	StatusFree
	StatusBusy
)

type Client struct {
	Id     string
	Status ClientStatus
}

func NewClient(id string) (Client, error) {
	for _, symbol := range id {
		if !isAllowed(symbol) {
			return Client{}, fmt.Errorf("incorrect client name")
		}
	}

	return Client{Id: id, Status: StatusFree}, nil
}

func isAllowed(symbol rune) bool {
	return symbol == '-' || symbol == '_' || (symbol >= '0' && symbol <= '9') || (symbol >= 'a' && symbol <= 'z')
}
