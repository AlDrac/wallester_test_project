package repositories

import (
	"net/url"

	"github.com/AlDrac/wallister_test_project/app/api/models"
)

type CustomerRepository interface {
	Create(customer *models.Customer) error
	GetById(int) (*models.Customer, error)
	Get(values url.Values) ([]models.Customer, error)
}
