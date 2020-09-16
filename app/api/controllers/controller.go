package controllers

import (
	"net/http"
)

type Action func(writer http.ResponseWriter, request *http.Request) error

type Controller struct{}

func (controller *Controller) Handler(action Action) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if err := action(writer, request); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})
}
