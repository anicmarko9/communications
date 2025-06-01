package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"communications/internal/database/dto"
	"communications/internal/services"
	"communications/internal/utils"
)

// Handles GET requests to check the application's health status.
// Pings the database and returns a success or error response based on the connectivity.
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

// Handles POST requests for incoming leads.
// Validates the request body, and triggers Email and SMS notifications to the client.
// Returns appropriate success or error responses based on the outcome.
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

	emailError := service.SendEmail(&id, &body)
	smsError := service.SendSMS(&id, &body)

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

// Helper to create a new Service instance with the current DB pool and config.
// Used to provide services with access to environment variables and the database.
func (h *Handler) newService() *services.Service {
	return services.NewService(h.Pool, h.Cfg)
}
