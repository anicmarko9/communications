package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"communications/internal/services"
	"communications/internal/utils"
)

func (dbHandler *DatabaseHandler) HealthHandler(c *gin.Context) {
	response := services.Health(dbHandler.Pool)
	code := http.StatusOK

	if response.Meta.Status == utils.StatusError {
		code = http.StatusInternalServerError
	}

	c.JSON(code, response)
}
