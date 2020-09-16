package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CustomerController struct {
	Controller
}

var Customer CustomerController

func (c CustomerController) Index(writer http.ResponseWriter, request *http.Request) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte("\"Name\":\"Alex\", Hobbies\":[\"snowboarding\",\"programming\"]}"))

	return nil
}

func (c CustomerController) GetCustomer(writer http.ResponseWriter, request *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		return err
	}

	fmt.Println(c.repository.Customer().GetById(id))

	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte("{\"Id\": " + mux.Vars(request)["id"] + ", \"Name\":\"Alex\", Hobbies\":[\"snowboarding\",\"programming\"]}"))

	return nil
}
