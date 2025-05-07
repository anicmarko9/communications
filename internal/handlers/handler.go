package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"communications/internal/services"
)

func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, services.Health())
}
