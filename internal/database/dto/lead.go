package dto

// Used to validate and bind incoming lead data from the request body.
// Ensures that required fields are present and conform to basic validation rules before further processing.
type CreateLeadDTO struct {
	Name    string  `json:"name" binding:"required,min=2,max=31"`         // Name of the user submitting the lead.
	Phone   string  `json:"phone" binding:"required,min=10,max=15"`       // User's phone number.
	Email   string  `json:"email" binding:"required,min=5,max=255,email"` // User's email address.
	Message *string `json:"message" binding:"omitempty,min=2,max=255"`    // Optional message from the user.
}

// Represents a single recipient's email address.
// Used as part of the Azure's email API payload.
type EmailRecipientAddress struct {
	Address string `json:"address" binding:"required,email"` // Email address of the recipient.
}

// Used to represent a list of email recipients for sending notifications.
// Designed to satisfy Azure's email API requirements.
type EmailRecipients struct {
	To []EmailRecipientAddress `json:"to" binding:"required"` // List of recipient email addresses.
}

// Contains the subject and HTML body for an email message.
// Used to structure the email payload for Azure's email API.
type EmailContent struct {
	Subject string `json:"subject" binding:"required"` // Subject of the email.
	HTML    string `json:"html" binding:"required"`    // HTML content of the email.
}

// This structure is used to create a valid payload for Azure's email API.
// Designed to satisfy Azure's email API requirements.
type EmailMessage struct {
	SenderAddress string                  `json:"senderAddress"` // Email address of the sender.
	Recipients    EmailRecipients         `json:"recipients"`    // List of email recipients.
	Content       EmailContent            `json:"content"`       // Content of the email.
	ReplyTo       []EmailRecipientAddress `json:"replyTo"`       // List of reply-to email addresses.
}

// This structure is used to create a valid payload for Azure's SMS API.
// Designed to satisfy Azure's SMS API requirements.
type SMSMessage struct {
	From    string   `json:"from" binding:"required"`    // Phone number of the sender.
	To      []string `json:"to" binding:"required"`      // List of SMS recipients.
	Message string   `json:"message" binding:"required"` // Textual message content.
}
