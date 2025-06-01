package models

import "time"

// Represents a business or user who receives leads from the website.
// This entity is manually added to the database and is used to associate incoming leads with the correct recipient.
// When a new lead is generated, the app checks for the corresponding client and notifies them (e.g., via email).
type Client struct {
	ID        string     `json:"id"`                   // Unique identifier for the client.
	Name      string     `json:"name"`                 // Name of the client.
	Email     string     `json:"email"`                // Contact email where this app sends lead notifications.
	Phone     string     `json:"phone"`                // Contact phone number for receiving lead notifications via SMS.
	Website   *string    `json:"website,omitempty"`    // Optional website URL.
	Verified  bool       `json:"verified"`             // Indicates if the client is verified.
	CreatedAt time.Time  `json:"created_at"`           // Timestamp when the client was created.
	UpdatedAt time.Time  `json:"updated_at"`           // Timestamp when the client was last updated.
	DeletedAt *time.Time `json:"deleted_at,omitempty"` // Timestamp when the client was deleted (if applicable).
}
