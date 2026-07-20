package service

import (
	"context"
)

type HealthService interface {
	Ready(ctx context.Context) error
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type healthService struct {
	db Pinger
}

func NewHealthService(db Pinger) *healthService {
	return &healthService{
		db: db,
	}
}

func (s *healthService) Ready(ctx context.Context) error {
	return s.db.Ping(ctx)
}
