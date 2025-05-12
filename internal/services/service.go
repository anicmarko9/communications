package services

import (
	"time"

	"communications/internal/utils"
)

func Health() utils.APIResponse[utils.DefaultResponse] {
	return utils.APIResponse[utils.DefaultResponse]{
		Meta: utils.Meta{
			Status:    utils.StatusSuccess,
			Message:   "Server is up and running.",
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
