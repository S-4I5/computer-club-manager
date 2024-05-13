package model

import (
	"errors"
)

var ClientDoesNotExistsError = errors.New("ClientUnknown")
var ComputerDoesNotExistsError = errors.New("computer does not exists")
var SessionDoesNotExistsError = errors.New("session does not exists")

var ClientAlreadyExistsError = errors.New("YouShallNotPass")
var ComputerAlreadyExistsError = errors.New("computer already exists")

var TooEarlyError = errors.New("NotOpenYet")

var NoOneWaitingError = errors.New("no one waiting")

var IncorrectComputerNameError = errors.New("incorrect computer name")

var FreeComputerAvailableError = errors.New("ICanWaitNoLonger")
var ComputerAlreadyOccupiedError = errors.New("PlaceIsBusy")

var QueueIsTooLongError = errors.New("queue is too long")
