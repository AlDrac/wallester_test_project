package models

import (
	"time"

	"github.com/AlDrac/wallister_test_project/app/api/models/validations"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

const (
	male   = "male"
	female = "female"
)

type Customer struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	BirthDate         time.Time `json:"birth_date" time_format:"sql_date"`
	Gender            string    `json:"gender"`
	Email             string    `json:"email"`
	Password          string    `json:"password,omitempty"`
	EncryptedPassword string    `json:"-"`
	Address           string    `json:"address"`
	Active            bool      `json:"active"`
	RegistrationDate  time.Time `json:"registration_date" time_format:"sql_datetime"`
}

func (customer *Customer) Validate() error {
	return validation.ValidateStruct(
		customer,
		validation.Field(&customer.FirstName, validation.Required, validation.Length(1, 100)),
		validation.Field(&customer.LastName, validation.Required, validation.Length(1, 100)),
		validation.Field(&customer.BirthDate, validation.Required, validation.By(validations.AgeRange(18, 60))),
		validation.Field(&customer.Gender, validation.Required, validation.In(female, male)),
		validation.Field(&customer.Email, validation.Required, is.Email),
		validation.Field(&customer.Password, validation.By(validations.RequiredIf(customer.EncryptedPassword == "")), validation.Length(6, 100)),
		validation.Field(&customer.Address, validation.Length(0, 200)),
	)
}

func (customer *Customer) ValidateEdit() error {
	return validation.ValidateStruct(
		customer,
		validation.Field(&customer.FirstName, validation.Required, validation.Length(1, 100)),
		validation.Field(&customer.LastName, validation.Required, validation.Length(1, 100)),
		validation.Field(&customer.BirthDate, validation.Required, validation.By(validations.AgeRange(18, 60))),
		validation.Field(&customer.Gender, validation.Required, validation.In(female, male)),
		validation.Field(&customer.Email, validation.Required, is.Email),
		validation.Field(&customer.Address, validation.Length(0, 200)),
	)
}

func (customer *Customer) BeforeCreate() error {
	if err := customer.Validate(); err != nil {
		return err
	}

	if len(customer.Password) > 0 {
		enc, err := encryptString(customer.Password)
		if err != nil {
			return err
		}

		customer.EncryptedPassword = enc
	}

	return nil
}

func (customer *Customer) BeforeEdit() error {
	if err := customer.ValidateEdit(); err != nil {
		return err
	}

	return nil
}

func encryptString(s string) (string, error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(enc), nil
}
