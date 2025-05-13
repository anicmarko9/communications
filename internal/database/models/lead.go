package models

import "time"

type Lead struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	Datetime  time.Time `json:"datetime"`
	UserName  *string   `json:"user_name,omitempty"`
	UserEmail *string   `json:"user_email,omitempty"`
	UserPhone *string   `json:"user_phone,omitempty"`
	ClientID  string    `json:"client_id"`
	Client    *Client   `json:"client,omitempty"`
}
