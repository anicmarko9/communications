package handlers

import (
	"errors"
	"net/http"
	"sync"

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
		utils.Reject(c, http.StatusServiceUnavailable, "Database connection failed.")
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

	id, err := h.validateID(c)
	if err != nil {
		return
	}

	body, err := h.validateBody(c)
	if err != nil {
		return
	}

	client, err := h.findClientByID(c, id)
	if err != nil {
		return
	}

	emailError, smsError := h.sendNotifications(service, client, &body)
	if emailError != nil && smsError != nil {
		utils.Reject(c, http.StatusInternalServerError, "Failed to send Email and SMS.")
		return
	}

	h.Pool.Exec(
		c.Request.Context(),
		`insert into "leads" ("name", "email", "phone", "client_id") values ($1, $2, $3, $4)`,
		body.Name,
		body.Email,
		body.Phone,
		id,
	)

	c.JSON(http.StatusOK, utils.APIResponse[utils.DefaultResponse]{
		Meta: utils.Meta{
			Status:    utils.StatusSuccess,
			Message:   "Email and SMS has been successfully sent to one of the clients.",
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

// Ensures the id param is a valid UUID to prevent invalid DB queries.
// Stops execution early if the input doesn't meet expected format, improving safety.
func (h *Handler) validateID(c *gin.Context) (string, error) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		utils.Reject(c, http.StatusBadRequest, "id must be a UUID")
		return "", err
	}

	return id, nil
}

// Binds and validates incoming lead data to enforce input integrity.
// Prevents invalid or incomplete data from reaching the notification or DB layers.
func (h *Handler) validateBody(c *gin.Context) (dto.CreateLeadDTO, error) {
	var body dto.CreateLeadDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.Reject(c, http.StatusBadRequest, err.Error())
		return body, err
	}

	if isValid := utils.ValidatePhoneNumber(body.Phone); !isValid {
		utils.Reject(c, http.StatusBadRequest, "phone number must start with + and contain at least 10 digits")
		return body, errors.New("invalid phone")
	}

	return body, nil
}

// Finds a client by ID in the database and retrieves their email and phone number.
// Validates client existence and soft-delete status before proceeding with lead logic.
func (h *Handler) findClientByID(c *gin.Context, id string) (*struct{ Email, Phone string }, error) {
	var client struct {
		Email string
		Phone string
	}

	row := h.Pool.QueryRow(
		c.Request.Context(),
		`select "email", "phone" from "clients" where "id" = $1 and "deleted_at" is null`,
		id,
	)

	if err := row.Scan(&client.Email, &client.Phone); err != nil {
		utils.Reject(c, http.StatusNotFound, "Client not found.")
		return nil, err
	}

	return &client, nil
}

// Sends email and SMS concurrently to reduce total response time and improve user experience.
// Captures and returns both errors to allow the caller to handle partial or complete notification failures.
func (h *Handler) sendNotifications(service *services.Service, client *struct{ Email, Phone string }, body *dto.CreateLeadDTO) (error, error) {
	var emailError, smsError error
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		emailError = service.SendEmail(&client.Email, body)
	}()

	go func() {
		defer wg.Done()
		smsError = service.SendSMS(&client.Phone, body)
	}()

	wg.Wait()

	return emailError, smsError
}
