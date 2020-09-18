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
	if err := customer.BeforeCreate(); err != nil {
		return err
	}

	return r.repository.db.QueryRow(
		"INSERT INTO customers c ("+
			"c.first_name, c.last_name, c.birth_date, c.gender, c.email, c.encrypted_password, c.address) "+
			"VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id;",
		customer.FirstName,
		customer.LastName,
		customer.BirthDate,
		customer.Gender,
		customer.Email,
		customer.EncryptedPassword,
		customer.Address,
	).Scan(&customer.ID)
}

func (r *CustomerRepository) Edit(customer *models.Customer) error {
	if err := customer.BeforeCreate(); err != nil {
		return err
	}

	_, err := r.repository.db.Exec(
		"UPDATE customers c SET "+
			"c.first_name = $2, c.last_name = $3, c.birth_date = $4, c.gender = $5, c.email = $6, c.address = $7 "+
			"WHERE c.id = $1;",
		customer.ID,
		customer.FirstName,
		customer.LastName,
		customer.BirthDate,
		customer.Gender,
		customer.Email,
		customer.Address,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerRepository) Delete(id *repositories.RequestId) error {
	_, err := r.repository.db.Exec(
		"UPDATE customers c SET "+
			"c.active = false "+
			"WHERE c.id = $1;",
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerRepository) Get(req *repositories.RequestSearch) ([]models.Customer, error) {
	rows, err := r.repository.db.Query(
		"SELECT c.id, c.first_name, c.last_name, c.birth_date, c.gender, c.email, c.encrypted_password, c.address, c.active, c.registration_date " +
			"FROM customers c " +
			"WHERE c.active = true " +
			"AND lower(c.first_name) SIMILAR TO '%" + req.FirstName + "%' " +
			"AND lower(c.last_name) SIMILAR TO '%" + req.LastName + "%'",
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

func (r *CustomerRepository) GetById(req *repositories.RequestId) (*models.Customer, error) {
	c := &models.Customer{}
	if err := r.repository.db.QueryRow(
		"SELECT c.id, c.first_name, c.last_name, c.birth_date, c.gender, c.email, c.encrypted_password, c.address, c.active, c.registration_date "+
			"FROM customers c "+
			"WHERE c.active = true "+
			"AND c.id = $1",
		req.Id,
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
