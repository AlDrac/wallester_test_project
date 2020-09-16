package repositories

import "github.com/AlDrac/wallister_test_project/app/api/models"

type CustomerRepository interface {
	Create(customer *models.Customer) error
	GetById(int) (*models.Customer, error)
}
