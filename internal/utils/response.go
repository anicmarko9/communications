package utils

import "github.com/gin-gonic/gin"

type Status string

const (
	StatusSuccess Status = "success"
	StatusError   Status = "error"
)

type DefaultResponse map[string]string

type Meta struct {
	Status    Status `json:"status"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type APIResponse[T any] struct {
	Meta Meta `json:"meta"`
	Data T    `json:"data"`
}

func Reject(c *gin.Context, code int, message string) {
	c.JSON(code, APIResponse[DefaultResponse]{
		Meta: Meta{
			Status:    StatusError,
			Message:   message,
			Timestamp: GetCurrentTimestamp(),
		},
		Data: DefaultResponse{},
	})
}
