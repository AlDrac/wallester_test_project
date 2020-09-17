package postgres

import (
	"database/sql"
	"github.com/AlDrac/wallister_test_project/app/api/models"
	"github.com/AlDrac/wallister_test_project/app/api/repositories"
)

type CustomerRepository struct {
	repository *Repository
}

func (r *CustomerRepository) Create(customer *models.Customer) error {
	return r.repository.db.QueryRow(
		"INSERT INTO customers ("+
			"first_name, last_name, birth_date, gender, email, encrypted_password, address) "+
			"VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		customer.FirstName,
		customer.LastName,
		customer.BirthDate,
		customer.Gender,
		customer.Email,
		customer.EncryptedPassword,
		customer.Address,
	).Scan(&customer.ID)
}

func (r *CustomerRepository) Get() ([]models.Customer, error) {
	rows, err := r.repository.db.Query(
		"SELECT id, first_name, last_name, birth_date, gender, email, encrypted_password, address, active, registration_date " +
			"FROM customers c " +
			"WHERE c.active = true",
	)
	if err != nil {
		return nil, err
	}

	customers := make([]models.Customer, 0)
	for rows.Next() {
		var c models.Customer

		if err := rows.Scan(
			&c.ID,
			&c.FirstName,
			&c.LastName,
			&c.BirthDate,
			&c.Gender,
			&c.Email,
			&c.EncryptedPassword,
			&c.Address,
			&c.Active,
			&c.RegistrationDate,
		); err != nil {
			return nil, err
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (r *CustomerRepository) GetById(id int) (*models.Customer, error) {
	c := &models.Customer{}
	if err := r.repository.db.QueryRow(
		"SELECT id, first_name, last_name, birth_date, gender, email, encrypted_password, address, active, registration_date "+
			"FROM customers c "+
			"WHERE c.active = true "+
			"AND c.id = $1",
		id,
	).Scan(
		&c.ID,
		&c.FirstName,
		&c.LastName,
		&c.BirthDate,
		&c.Gender,
		&c.Email,
		&c.EncryptedPassword,
		&c.Address,
		&c.Active,
		&c.RegistrationDate,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, repositories.ErrRecordNotFound
		}
		return nil, err
	}
	return c, nil
}
