package dto

type CreateLeadDTO struct {
	Name    string  `json:"name" binding:"required,min=2,max=31"`
	Phone   string  `json:"phone" binding:"required,min=10,max=15"`
	Email   string  `json:"email" binding:"required,min=5,max=255,email"`
	Message *string `json:"message" binding:"omitempty,min=2,max=255"`
}
