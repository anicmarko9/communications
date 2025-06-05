package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"communications/internal/database/dto"
	"communications/internal/utils"
)

// Prepares and sends an email notification to the specified recipient.
// Uses Azure's email API payload structure.
// Returns an error if the required parameters are missing or if the sending fails.
func (s *Service) SendEmail(to *string, params *dto.CreateLeadDTO) error {
	if to == nil || params == nil {
		return errors.New("id and payload are required")
	}

	endpoint, key, err := parseACS(s.Cfg.AzureURL)
	if err != nil {
		return err
	}

	message := dto.EmailMessage{
		SenderAddress: s.Cfg.EmailFrom,
		Recipients:    dto.EmailRecipients{To: []dto.EmailRecipientAddress{{Address: *to}}},
		Content:       dto.EmailContent{Subject: "New Website Lead", HTML: setHTML(params)},
		ReplyTo:       []dto.EmailRecipientAddress{{Address: params.Email}},
	}

	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}

	azureError := sendAzureEmail(endpoint, key, payload)
	if azureError != nil {
		return azureError
	}

	return nil
}

// Parses Azure Connection String
func parseACS(url string) (endpoint, key string, err error) {
	parts := utils.SplitString(url, ";")

	for _, part := range parts {
		if strings.HasPrefix(part, "endpoint=") {
			endpoint = strings.TrimPrefix(part, "endpoint=")
		}

		if strings.HasPrefix(part, "accesskey=") {
			key = strings.TrimPrefix(part, "accesskey=")
		}
	}

	if endpoint == "" || key == "" {
		return "", "", errors.New("invalid Azure Connection String")
	}

	return endpoint, key, nil
}

// Sends an email via Azure Communication Services Email REST API.
func sendAzureEmail(endpoint, key string, payload []byte) error {
	url := fmt.Sprintf("%s/emails:send?api-version=2023-03-31", endpoint)
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
		fmt.Printf("Azure Email Service ~ %s: %s\n", res.Status, resBody)
		return errors.New("failed to send email: " + res.Status)
	}

	return nil
}

// Generates the HTML body for the lead notification email.
// Used to format the email content sent to the client.
func setHTML(params *dto.CreateLeadDTO) string {
	message := ""
	if params != nil && params.Message != nil {
		message = *params.Message
	}

	return `
		<!doctype html>
		<html
			xmlns="http://www.w3.org/1999/xhtml"
			xmlns:v="urn:schemas-microsoft-com:vml"
			xmlns:o="urn:schemas-microsoft-com:office:office"
		>
			<head>
				<title>New Website Lead</title>
				<meta name="robots" content="noindex, nofollow" />
				<meta name="referrer" content="no-referrer" />
				<meta charset="UTF-8" />
				<meta http-equiv="Content-Type" content="text/html charset=UTF-8" />
				<meta name="viewport" content="width=device-width, initial-scale=1" />
				<style type="text/css">
					table {
						border-collapse: separate;
					}
					a,
					a:link,
					a:visited {
						text-decoration: none;
						color: #00788a;
					}
					a:hover {
						text-decoration: underline;
					}
					h2,
					h2 a,
					h2 a:visited,
					h3,
					h3 a,
					h3 a:visited,
					h4,
					h5,
					h6,
					.t_cht {
						color: #000 !important;
					}
					.ExternalClass p,
					.ExternalClass span,
					.ExternalClass font,
					.ExternalClass td {
						line-height: 100%;
					}
					.ExternalClass {
						width: 100%;
					}
					body {
						font-family: Arial, sans-serif;
						background-color: #f5f5f5;
						margin: 0;
						padding: 0;
					}
					.container {
						max-width: 600px;
						margin: 20px auto;
						background-color: #fff;
						border-radius: 8px;
						padding: 20px;
						box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
					}
					h1 {
						text-align: center;
						color: #333;
					}
					.info-container {
						margin-top: 20px;
					}
					.info-container p {
						margin: 5px 0;
						color: #555;
					}
					.info-container p strong {
						color: #333;
					}
					.footer {
						margin-top: 20px;
						text-align: center;
						color: #999;
						font-size: 14px;
					}
				</style>
			</head>

			<body style="background: #ffffff; font-family: 'Circular', sans-serif; margin: 0 auto; padding: 0">
				<div class="container">
					<h1>New Website Lead</h1>
					<div class="info-container">
						<p><strong>Name:</strong> ` + params.Name + `</p>
						<p><strong>Email:</strong> ` + params.Email + `</p>
						<p><strong>Phone:</strong> ` + params.Phone + `</p>
						<p><strong>Message:</strong> ` + message + `</p>
					</div>
					<p class="footer">This email was sent via the Contact Us form.</p>
				</div>
			</body>
		</html>
	`
}
