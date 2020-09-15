package models

import "time"

type Customer struct {
	ID               int
	FirstName        string
	LastName         string
	BirthDate        time.Time
	Gender           string
	Email            string
	Address          string
	Active           bool
	RegistrationDate time.Time
}
