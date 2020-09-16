package repositories

type Repository interface {
	Customer() CustomerRepository
}
