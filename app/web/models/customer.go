package models

import "time"

type Customer struct {
	ID               int       `json:"id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	BirthDate        string    `json:"birth_date" time_format:"sql_date"`
	Gender           string    `json:"gender"`
	Email            string    `json:"email"`
	Password         string    `json:"password,omitempty"`
	Address          string    `json:"address"`
	Active           bool      `json:"active"`
	RegistrationDate time.Time `json:"registration_date" time_format:"sql_datetime"`
}
