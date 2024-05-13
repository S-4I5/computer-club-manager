package entity

import "computer-club-manager/internal/lib/time"

type Session struct {
	Computer  Computer
	Client    Client
	StartedAt time.Timestamp
	EndedAt   time.Timestamp
}

func (s Session) GetLengthInHoursRounded() int {
	if s.EndedAt.Value == 0 {
		return -1
	}

	if (s.EndedAt.Value-s.StartedAt.Value)%60 != 0 {
		return (s.EndedAt.Value-s.StartedAt.Value)/60 + 1
	}

	return (s.EndedAt.Value - s.StartedAt.Value) / 60
}

func (s Session) GetLengthInMinutes() int {
	if s.EndedAt.Value == 0 {
		return -1
	}

	return s.EndedAt.Value - s.StartedAt.Value
}

func NewSession(
	computer Computer,
	client Client,
	start time.Timestamp,
	end time.Timestamp,
) Session {
	return Session{
		Computer:  computer,
		Client:    client,
		StartedAt: start,
		EndedAt:   end,
	}
}
