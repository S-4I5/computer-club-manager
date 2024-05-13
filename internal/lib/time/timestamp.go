package time

import (
	"fmt"
	"strconv"
	"strings"
)

type Timestamp struct {
	Value int
}

func (t Timestamp) GetHours() int {
	return t.Value / 60
}

func (t Timestamp) GetHoursRounded() int {
	if t.GetMinutes() > 0 {
		return t.GetHours() + 1
	}

	return t.GetHours()
}

func (t Timestamp) GetMinutes() int {
	return t.Value % 60
}

func (t Timestamp) ToString() string {

	hours := fmt.Sprint(t.GetHours())
	minutes := fmt.Sprint(t.GetMinutes())

	if t.GetHours() < 10 {
		hours = "0" + hours
	}

	if t.GetMinutes() < 10 {
		minutes = "0" + minutes
	}

	return fmt.Sprint(hours, ":", minutes)
}

func NewTimestamp(time int) (Timestamp, error) {
	return Timestamp{Value: time}, nil
}

func NewTimestampFromString(time string) (Timestamp, error) {
	splitted := strings.Split(time, ":")

	hours, err := strconv.Atoi(splitted[0])
	if err != nil {
		return Timestamp{}, err
	}

	minutes, err := strconv.Atoi(splitted[1])
	if err != nil {
		return Timestamp{}, err
	}

	return NewTimestamp(hours*60 + minutes)
}
