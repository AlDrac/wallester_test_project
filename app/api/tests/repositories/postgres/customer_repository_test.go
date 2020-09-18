package postgres_test

import (
	"github.com/AlDrac/wallister_test_project/app/api/models"
	repositories2 "github.com/AlDrac/wallister_test_project/app/api/repositories"
	"github.com/AlDrac/wallister_test_project/app/api/repositories/postgres"
	"github.com/AlDrac/wallister_test_project/app/api/tests/repositories"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"

	"github.com/AlDrac/wallister_test_project/app/api/configs"
)

var url string

func TestMain(m *testing.M) {
	config := configs.New(true)
	url = config.Database.UrlTest

	os.Exit(m.Run())
}

func getCustomer(t *testing.T) *models.Customer {
	t.Helper()

	return &models.Customer{
		FirstName: "Test",
		LastName:  "Tester",
		BirthDate: time.Date(2000, time.November, 10, 0, 0, 0, 0, time.UTC),
		Gender:    "male",
		Email:     "test.tester100@test.com",
		Password:  "qwerty",
		Address:   "Test Test 50-4 12345",
	}
}

func TestCustomerRepository_Create(t *testing.T) {
	db, teardown := repositories.InitialiseTestDatabase(t, url)
	defer teardown("customers")

	cR := postgres.InitialiseRepository(db).Customer()
	c := getCustomer(t)

	assert.NoError(t, cR.Create(c))
	assert.NotNil(t, c.ID)
}

func TestCustomerRepository_Edit(t *testing.T) {
	db, teardown := repositories.InitialiseTestDatabase(t, url)
	defer teardown("customers")

	cR := postgres.InitialiseRepository(db).Customer()
	c := getCustomer(t)
	_ = cR.Create(c)

	c.FirstName = "TesterTester"
	assert.NoError(t, cR.Edit(c))
}

func TestCustomerRepository_Delete(t *testing.T) {
	db, teardown := repositories.InitialiseTestDatabase(t, url)
	defer teardown("customers")

	cR := postgres.InitialiseRepository(db).Customer()
	c := getCustomer(t)
	_ = cR.Create(c)

	req := &repositories2.RequestId{
		Id: c.ID,
	}
	assert.NoError(t, cR.Delete(req))
}

func TestCustomerRepository_Get(t *testing.T) {
	db, teardown := repositories.InitialiseTestDatabase(t, url)
	defer teardown("customers")

	cR := postgres.InitialiseRepository(db).Customer()
	c := getCustomer(t)
	_ = cR.Create(c)

	req := &repositories2.RequestSearch{
		FirstName: c.FirstName,
		LastName:  c.LastName,
	}

	customers, err := cR.Get(req)
	assert.NoError(t, err)
	assert.NotNil(t, customers)
}

func TestCustomerRepository_GetById(t *testing.T) {
	db, teardown := repositories.InitialiseTestDatabase(t, url)
	defer teardown("customers")

	cR := postgres.InitialiseRepository(db).Customer()
	c := getCustomer(t)
	_ = cR.Create(c)

	req := &repositories2.RequestId{
		Id: c.ID,
	}

	customer, err := cR.GetById(req)
	assert.NoError(t, err)
	assert.NotNil(t, customer)
}

func TestCustomerRepository_InvalidGetById(t *testing.T) {
	db, teardown := repositories.InitialiseTestDatabase(t, url)
	defer teardown("customers")

	cR := postgres.InitialiseRepository(db).Customer()
	c := getCustomer(t)
	_ = cR.Create(c)

	req := &repositories2.RequestId{
		Id: 99999,
	}

	customer, err := cR.GetById(req)
	assert.EqualError(t, err, repositories2.ErrRecordNotFound.Error())
	assert.Nil(t, customer)
}

func TestCustomerRepository_GetByEmail(t *testing.T) {
	db, teardown := repositories.InitialiseTestDatabase(t, url)
	defer teardown("customers")

	cR := postgres.InitialiseRepository(db).Customer()
	c := getCustomer(t)
	_ = cR.Create(c)

	customer, err := cR.GetByEmail(c.Email)
	assert.NoError(t, err)
	assert.NotNil(t, customer)
}
