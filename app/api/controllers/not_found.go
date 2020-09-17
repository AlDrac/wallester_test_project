package controllers

import (
	"github.com/darahayes/go-boom"
	"net/http"
)

type NotFoundController struct {
	Controller
}

var NotFound NotFoundController

func (c NotFoundController) Index(writer http.ResponseWriter, request *http.Request) error {
	boom.NotFound(writer, "404 Page not found")
	return nil
}
