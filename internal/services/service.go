package services

import (
	"github.com/gin-gonic/gin"
)

func Health() map[string]any {
	return gin.H{"status": "ok", "message": "success"}
}

func Throttler() map[string]any {
	return gin.H{"status": "error", "message": "rate limit exceeded"}
}
