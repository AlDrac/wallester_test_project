package repositories

import (
	"github.com/AlDrac/wallister_test_project/app/api/models"
	"net/url"
)

type CustomerRepository interface {
	Create(customer *models.Customer) error
	Edit(customer *models.Customer) error
	Delete(int) error
	Get(values url.Values) ([]models.Customer, error)
	GetById(int) (*models.Customer, error)
}
