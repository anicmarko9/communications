package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"communications/internal/database/dto"
	"communications/internal/services"
	"communications/internal/utils"
)

func (h *Handler) HealthHandler(c *gin.Context) {
	service := h.newService()

	if err := service.CheckHealth(); err != nil {
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

func (h *Handler) LeadHandler(c *gin.Context) {
	service := h.newService()

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

	emailError := service.SendEmail()
	smsError := service.SendSMS()

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

func (h *Handler) newService() *services.Service {
	return services.NewService(h.Pool, h.Cfg)
}
