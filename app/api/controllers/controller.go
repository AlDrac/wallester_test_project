package controllers

import (
	"net/http"

	"github.com/AlDrac/wallister_test_project/app/api/repositories"
	"github.com/darahayes/go-boom"
)

type Action func(writer http.ResponseWriter, request *http.Request) error

type Controller struct {
	repository repositories.Repository
}

func InitialiseController(repository repositories.Repository) Controller {
	return Controller{
		repository: repository,
	}
}

func (c *Controller) Handler(action Action) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if err := action(writer, request); err != nil {
			boom.Internal(writer, err)
		}
	})
}
