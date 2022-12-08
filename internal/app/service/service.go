package service

import (
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int64 {
	data := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)

	duration := time.Until(data)

	return int64(duration.Hours() / 24)
}
