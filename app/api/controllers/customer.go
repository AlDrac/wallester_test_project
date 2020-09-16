package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerController struct {
	Controller
}

var Customer CustomerController

func (controller CustomerController) Index(writer http.ResponseWriter, request *http.Request) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte("\"Name\":\"Alex\", Hobbies\":[\"snowboarding\",\"programming\"]}"))

	return nil
}

func (controller CustomerController) GetCustomer(writer http.ResponseWriter, request *http.Request) error {
	id := mux.Vars(request)["id"]
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte("{\"Id\": " + id + ", \"Name\":\"Alex\", Hobbies\":[\"snowboarding\",\"programming\"]}"))

	return nil
}
