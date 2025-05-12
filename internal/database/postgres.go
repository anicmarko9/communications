package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"communications/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(cfg *config.Config) *pgxpool.Pool {

	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseName,
		cfg.DatabaseSSL,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)

	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
	}
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping the database: %v", err)
	}

	return pool
}
