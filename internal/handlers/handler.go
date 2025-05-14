package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"communications/internal/database/dto"
	"communications/internal/services"
	"communications/internal/utils"
)

func (dbHandler *DatabaseHandler) HealthHandler(c *gin.Context) {
	if err := services.CheckHealth(dbHandler.Pool); err != nil {
		utils.Reject(c, http.StatusInternalServerError, "Database connection failed.")
		return
	}

	c.JSON(http.StatusOK, utils.APIResponse[utils.DefaultResponse]{
		Meta: utils.Meta{
			Status:    utils.StatusSuccess,
			Message:   "Database is up and running.",
			Timestamp: utils.GetCurrentTimestamp(),
		},
		Data: utils.DefaultResponse{},
	})
}

func (dbHandler *DatabaseHandler) LeadHandler(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		utils.Reject(c, http.StatusBadRequest, "id must be a UUID")
		return
	}

	var body dto.CreateLeadDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.Reject(c, http.StatusBadRequest, err.Error())
		return
	}

	if isValid := utils.ValidatePhoneNumber(body.Phone); !isValid {
		utils.Reject(c, http.StatusBadRequest, "phone number must start with + and contain at least 10 digits")
		return
	}

	emailError := services.SendEmail(dbHandler.Pool, dbHandler.Cfg)
	smsError := services.SendSMS(dbHandler.Pool, dbHandler.Cfg)

	if emailError != nil && smsError != nil {
		utils.Reject(c, http.StatusInternalServerError, "Failed to send Email and SMS.")
		return
	}

	c.JSON(http.StatusOK, utils.APIResponse[utils.DefaultResponse]{
		Meta: utils.Meta{
			Status:    utils.StatusSuccess,
			Message:   "Email and SMS have been successfully sent to one of the clients.",
			Timestamp: utils.GetCurrentTimestamp(),
		},
		Data: utils.DefaultResponse{},
	})
}
