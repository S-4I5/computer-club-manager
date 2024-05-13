package config

import (
	"bufio"
	"computer-club-manager/internal/lib/time"
	"errors"
	"fmt"
)

var IncorrectDatesError = errors.New("startTime cannot be later then endTime")

type Config struct {
	NumberOfComputers int
	PricePerHour      int
	WorkdayStart      time.Timestamp
	WorkdayEnd        time.Timestamp
}

func ReadFrom(reader *bufio.Reader) (Config, error) {
	var numberOfComputers, pricePerHour int
	var workdayStartString, workdayEndString string

	_, err := fmt.Fscan(
		reader,
		&numberOfComputers,
		&workdayStartString,
		&workdayEndString,
		&pricePerHour,
	)
	if err != nil {
		return Config{}, err
	}

	workdayStart, err := time.NewTimestampFromString(workdayStartString)
	if err != nil {
		return Config{}, err
	}

	workdayEnd, err := time.NewTimestampFromString(workdayEndString)
	if err != nil {
		return Config{}, err
	}

	if workdayStart.Value > workdayEnd.Value {
		return Config{}, IncorrectDatesError
	}

	return Config{
		NumberOfComputers: numberOfComputers,
		PricePerHour:      pricePerHour,
		WorkdayStart:      workdayStart,
		WorkdayEnd:        workdayEnd,
	}, nil
}
