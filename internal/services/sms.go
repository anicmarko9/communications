package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"communications/internal/database/dto"
)

// Prepares and sends an SMS notification to the specified recipient.
// Uses Azure's SMS API payload structure.
// Returns an error if required parameters are missing or if sending fails.
func (s *Service) SendSMS(to *string, params *dto.CreateLeadDTO) error {
	if to == nil || params == nil {
		return errors.New("id and payload are required")
	}

	endpoint, key, err := parseACS(s.Cfg.AzureURL)
	if err != nil {
		return err
	}

	from := s.Cfg.SMSFrom
	message := setSMSContent(params)

	smsBody := dto.SMSMessage{
		From:    from,
		To:      []string{*to},
		Message: message,
	}

	payload, err := json.Marshal(smsBody)
	if err != nil {
		return err
	}

	azureError := sendAzureSMS(endpoint, key, payload)
	if azureError != nil {
		return azureError
	}

	return nil
}

// Sends an SMS via Azure Communication Services SMS REST API.
func sendAzureSMS(endpoint, key string, payload []byte) error {
	url := fmt.Sprintf("%s/sms", endpoint)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", key)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	bodyBytes, _ := io.ReadAll(res.Body)
	resBody := string(bodyBytes)

	if res.StatusCode >= 300 {
		fmt.Printf("Azure SMS Service ~ %s: %s\n", res.Status, resBody)
		return errors.New("failed to send SMS: " + res.Status)
	}

	return nil
}

// Generates the Textual message content for the lead notification.
func setSMSContent(params *dto.CreateLeadDTO) string {
	return "New Website Lead\n\n" +
		"Name: " + params.Name + "\n" +
		"Email: " + params.Email + "\n" +
		"Phone: " + params.Phone + "\n"
}
