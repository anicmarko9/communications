package services

import (
	"context"
	"time"

	"communications/internal/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Health(db *pgxpool.Pool) utils.APIResponse[utils.DefaultResponse] {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	status := utils.StatusSuccess
	message := "Server is up and running."

	if err := db.Ping(ctx); err != nil {
		status = utils.StatusError
		message = "Database connection failed."
	}

	return utils.APIResponse[utils.DefaultResponse]{
		Meta: utils.Meta{
			Status:    status,
			Message:   message,
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
		Data: utils.DefaultResponse{},
	}
}

func Throttler() utils.APIResponse[utils.DefaultResponse] {
	return utils.APIResponse[utils.DefaultResponse]{
		Meta: utils.Meta{
			Status:    utils.StatusError,
			Message:   "Too many requests. Please try again later.",
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
		Data: utils.DefaultResponse{},
	}
}
