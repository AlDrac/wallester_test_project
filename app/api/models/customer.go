package models

import "time"

type Customer struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	BirthDate         time.Time `json:"birth_date"`
	Gender            string    `json:"gender"`
	Email             string    `json:"email"`
	EncryptedPassword string    `json:"encrypted_password"`
	Address           string    `json:"address"`
	Active            bool      `json:"active"`
	RegistrationDate  time.Time `json:"registration_date"`
}
