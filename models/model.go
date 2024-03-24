package models

import "time"

// Employees Model
type Employees struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone   string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Responsoe Model
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}