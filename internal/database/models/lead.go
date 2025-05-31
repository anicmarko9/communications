package models

import "time"

// Represents a contact or inquiry submitted by a user through the website.
// Each lead is associated with a client and contains information from the user's submission (e.g., Contact Us form).
// The "type" field indicates how the lead should be delivered (Email or SMS, or both).
// Used to track and retrieve all leads for a specific client.
type Lead struct {
	ID        int       `json:"id"`                   // Unique identifier for the lead.
	Type      string    `json:"type"`                 // Delivery type: "email" or "sms".
	Datetime  time.Time `json:"datetime"`             // Timestamp when the lead was created.
	UserName  *string   `json:"user_name,omitempty"`  // Name entered by the user (optional).
	UserEmail *string   `json:"user_email,omitempty"` // Email entered by the user (optional).
	UserPhone *string   `json:"user_phone,omitempty"` // Phone entered by the user (optional).
	ClientID  string    `json:"client_id"`            // Associated client ID.
	Client    *Client   `json:"client,omitempty"`     // Optional client details.
}
