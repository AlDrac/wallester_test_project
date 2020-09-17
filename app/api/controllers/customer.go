package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CustomerController struct {
	Controller
}

var Customer CustomerController

func (c CustomerController) Index(writer http.ResponseWriter, request *http.Request) error {
	customers, err := c.repository.Customer().Get()
	if err != nil {
		return err
	}

	body, err := json.Marshal(customers)
	if err != nil {
		return err
	}

	if err = c.responseJson(writer, string(body)); err != nil {
		return err
	}

	return nil
}

func (c CustomerController) GetCustomer(writer http.ResponseWriter, request *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		return err
	}

	customer, err := c.repository.Customer().GetById(id)
	if err != nil {
		return err
	}

	body, err := json.Marshal(customer)
	if err != nil {
		return err
	}

	if err = c.responseJson(writer, string(body)); err != nil {
		return err
	}

	return nil
}
