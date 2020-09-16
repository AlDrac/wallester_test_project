package postgres

import (
	"database/sql"
	"github.com/AlDrac/wallister_test_project/app/api/repositories"
	_ "github.com/lib/pq"
)

type Repository struct {
	db                 *sql.DB
	customerRepository *CustomerRepository
}

func InitialiseRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Customer() repositories.CustomerRepository {
	if r.customerRepository != nil {
		return r.customerRepository
	}

	r.customerRepository = &CustomerRepository{
		repository: r,
	}

	return r.customerRepository
}
