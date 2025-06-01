package services

import (
	"errors"

	"communications/internal/database/dto"
)

// Prepares and sends an email notification to the specified recipient.
// Uses Azure's email API payload structure.
// Returns an error if the required parameters are missing or if the sending fails.
func (s *Service) SendEmail(to *string, params *dto.CreateLeadDTO) error {
	if to == nil || params == nil {
		return errors.New("id and payload are required")
	}

	senderAddress := s.Cfg.EmailFrom
	recipients := dto.EmailRecipients{To: []dto.EmailRecipientAddress{{Address: *to}}}
	replyTo := []dto.EmailRecipientAddress{{Address: params.Email}}
	content := dto.EmailContent{Subject: "New Website Lead", HTML: setHTML(params)}

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
