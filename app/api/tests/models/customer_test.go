package models_test

import (
	"testing"
	"time"

	"github.com/AlDrac/wallister_test_project/app/api/models"
	"github.com/stretchr/testify/assert"
)

func getCustomer(t *testing.T) *models.Customer {
	t.Helper()

	return &models.Customer{
		FirstName:        "Test",
		LastName:         "Tester",
		BirthDate:        time.Date(2000, time.November, 10, 0, 0, 0, 0, time.UTC),
		Gender:           "male",
		Email:            "test.tester@test.com",
		Password:         "qwerty",
		Address:          "Test Test 50-4 12345",
		Active:           true,
		RegistrationDate: time.Now(),
	}
}

func getCustomerWithoutBirthDate(t *testing.T) *models.Customer {
	t.Helper()

	return &models.Customer{
		FirstName:        "Test",
		LastName:         "Tester",
		Gender:           "male",
		Email:            "test.tester@test.com",
		Password:         "qwerty",
		Address:          "Test Test 50-4 12345",
		Active:           true,
		RegistrationDate: time.Now(),
	}
}

func getCustomerWithoutAddress(t *testing.T) *models.Customer {
	t.Helper()

	return &models.Customer{
		FirstName:        "Test",
		LastName:         "Tester",
		BirthDate:        time.Date(2000, time.November, 10, 0, 0, 0, 0, time.UTC),
		Gender:           "male",
		Email:            "test.tester@test.com",
		Password:         "qwerty",
		Active:           true,
		RegistrationDate: time.Now(),
	}
}

func TestCustomer_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		c       func() *models.Customer
		isValid bool
	}{
		{
			name: "valid customer",
			c: func() *models.Customer {
				return getCustomer(t)
			},
			isValid: true,
		},
		{
			name: "invalid first name required",
			c: func() *models.Customer {
				c := getCustomer(t)
				c.FirstName = ""
				return c
			},
			isValid: false,
		},
		{
			name: "invalid first name max length 100",
			c: func() *models.Customer {
				c := getCustomer(t)
				c.FirstName = "qwertyuiopasdfghjklzxcvb" +
					"qwertyuiopasdfghjklzxcvb" +
					"qwertyuiopasdfghjklzxcvb" +
					"qwertyuiopasdfghjklzxcvb" +
					"qwertyuiopasdfghjklzxcvb"
				return c
			},
			isValid: false,
		},
		{
			name: "invalid last name required",
			c: func() *models.Customer {
				c := getCustomer(t)
				c.LastName = ""
				return c
			},
			isValid: false,
		},
		{
			name: "invalid last name max length 100",
			c: func() *models.Customer {
				c := getCustomer(t)
				c.LastName = "qwertyuiopasdfghjklzxcvb" +
					"qwertyuiopasdfghjklzxcvb" +
					"qwertyuiopasdfghjklzxcvb" +
					"qwertyuiopasdfghjklzxcvb" +
					"qwertyuiopasdfghjklzxcvb"
				return c
			},
			isValid: false,
		},
		{
			name: "invalid birth date required",
			c: func() *models.Customer {
				return getCustomerWithoutBirthDate(t)
			},
			isValid: false,
		},
		{
			name: "invalid birth date < 18",
			c: func() *models.Customer {
				c := getCustomer(t)
				c.BirthDate = time.Date(time.Now().Year(), time.November, 10, 0, 0, 0, 0, time.UTC)
				return c
			},
			isValid: false,
		},
		{
			name: "invalid birth date > 60",
			c: func() *models.Customer {
				c := getCustomer(t)
				c.BirthDate = time.Date(1950, time.November, 10, 0, 0, 0, 0, time.UTC)
				return c
			},
			isValid: false,
		},
		{
			name: "valid customer gender",
			c: func() *models.Customer {
				c := getCustomer(t)
				c.Gender = "female"
				return c
			},
			isValid: true,
		},
		{
			name: "invalid gender required",
			c: func() *models.Customer {
				c := getCustomer(t)
				c.Gender = ""
				return c
			},
			isValid: false,
		},
		{
			name: "invalid gender any",
			c: func() *models.Customer {
				c := getCustomer(t)
				c.Gender = "any"
				return c
			},
			isValid: false,
		},
		{
			name: "invalid email required",
			c: func() *models.Customer {
				c := getCustomer(t)
				c.Email = ""
				return c
			},
			isValid: false,
		},
		{
			name: "invalid email type",
			c: func() *models.Customer {
				c := getCustomer(t)
				c.Email = "test.test@"
				return c
			},
			isValid: false,
		},
		{
			name: "valid address empty",
			c: func() *models.Customer {
				return getCustomerWithoutAddress(t)
			},
			isValid: true,
		},
		{
			name: "invalid address > 200",
			c: func() *models.Customer {
				c := getCustomer(t)
				c.Address = "qwertyuiopasdfghjklzxcvb" +
					"qwertyuiopasdfghjklzxcvb" +
					"qwertyuiopasdfghjklzxcvb" +
					"qwertyuiopasdfghjklzxcvb" +
					"qwertyuiopasdfghjklzxcvb" +
					"qwertyuiopasdfghjklzxcvb" +
					"qwertyuiopasdfghjklzxcvb" +
					"qwertyuiopasdfghjklzxcvb" +
					"qwertyuiopasdfghjklzxcvb"
				return c
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.c().Validate())
			} else {
				assert.Error(t, tc.c().Validate())
			}
		})
	}
}

func TestCustomer_ValidateEdit(t *testing.T) {
	testCases := []struct {
		name    string
		c       func() *models.Customer
		isValid bool
	}{
		{
			name: "valid customer for edit",
			c: func() *models.Customer {
				c := getCustomer(t)
				c.Password = ""
				return c
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.c().ValidateEdit())
			} else {
				assert.Error(t, tc.c().ValidateEdit())
			}
		})
	}
}
