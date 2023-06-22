package model

import "time"

type Contacts struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
