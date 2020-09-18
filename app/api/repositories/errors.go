package repositories

import "errors"

var (
	ErrRecordNotFound = errors.New("Record not found")
	ErrRecordExist    = errors.New("Record exist")
)
