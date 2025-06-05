package models

import "time"

// Represents a contact or inquiry submitted by a user through the website.
// Each lead is associated with a client and contains information from the user's submission (e.g., Contact Us form).
// Used to track and retrieve all leads for a specific client.
type Lead struct {
	ID       int       `json:"id"`               // Unique identifier for the lead.
	Datetime time.Time `json:"datetime"`         // Timestamp when the lead was created.
	Name     *string   `json:"name,omitempty"`   // Name entered by the user (optional).
	Email    *string   `json:"email,omitempty"`  // Email entered by the user (optional).
	Phone    *string   `json:"phone,omitempty"`  // Phone entered by the user (optional).
	ClientID string    `json:"client_id"`        // Associated client ID.
	Client   *Client   `json:"client,omitempty"` // Optional client details.
}
