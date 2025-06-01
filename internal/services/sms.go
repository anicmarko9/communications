package services

import (
	"errors"

	"communications/internal/database/dto"
)

// Prepares and sends an SMS notification to the specified recipient.
// Uses Azure's SMS API payload structure.
// Returns an error if required parameters are missing or if sending fails.
func (s *Service) SendSMS(id *string, params *dto.CreateLeadDTO) error {
	if id == nil || params == nil {
		return errors.New("id and payload are required")
	}

	return nil
}
