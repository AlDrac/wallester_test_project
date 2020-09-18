package repositories

import (
	"github.com/AlDrac/wallister_test_project/app/api/models"
)

type RequestId struct {
	Id int `json:"id"`
}

type RequestSearch struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type CustomerRepository interface {
	Create(*models.Customer) error
	Edit(*models.Customer) error
	Delete(*RequestId) error
	Get(*RequestSearch) ([]models.Customer, error)
	GetById(*RequestId) (*models.Customer, error)
	GetByEmail(string) (*models.Customer, error)
}
