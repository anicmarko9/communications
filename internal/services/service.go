package services

import (
	"communications/internal/config"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Provides access to the database pool and configuration for business logic operations.
type Service struct {
	Pool *pgxpool.Pool
	Cfg  *config.Config
}

// Creates a new Service instance with the provided database pool and config.
func NewService(db *pgxpool.Pool, cfg *config.Config) *Service {
	return &Service{Pool: db, Cfg: cfg}
}

// Pings the database to verify connectivity.
func (s *Service) CheckHealth() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.Pool.Ping(ctx)
}
