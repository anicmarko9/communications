package utils

import "github.com/gin-gonic/gin"

type Status string

const (
	StatusSuccess Status = "success" // Indicates a successful API response.
	StatusError   Status = "error"   // Indicates an error in the API response.
)

// Generic map for simple {"key": "value"} responses.
type DefaultResponse map[string]string

// Contains metadata for every API response, such as status, message, and timestamp.
type Meta struct {
	Status    Status `json:"status"`    // "success" or "error"
	Message   string `json:"message"`   // Human-readable message for the response.
	Timestamp string `json:"timestamp"` // ISO timestamp of the response.
}

// Wraps the response body and metadata for all API responses.
// Data is the actual payload returned to the frontend.
type APIResponse[T any] struct {
	Meta Meta `json:"meta"` // Metadata about the response.
	Data T    `json:"data"` // Actual response data.
}

// Helper to send a standardized error response and status code.
// Use this to return errors (e.g., 404 for not found, 400 for bad request) in a consistent format.
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
