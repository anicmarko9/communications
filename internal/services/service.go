package services

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CheckHealth(db *pgxpool.Pool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return db.Ping(ctx)
}
