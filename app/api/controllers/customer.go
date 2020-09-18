package controllers

import (
	"net/http"
	"strconv"

	"github.com/AlDrac/wallister_test_project/app/api/models"
	"github.com/AlDrac/wallister_test_project/app/api/repositories"
	"github.com/gorilla/mux"
)

type CustomerController struct {
	Controller
}

var Customer CustomerController

func (c CustomerController) GetCustomers(writer http.ResponseWriter, request *http.Request) error {
	req := &repositories.RequestSearch{}
	if err := json.NewDecoder(request.Body).Decode(req); err != nil {
		return err
	}

	customers, err := c.repository.Customer().Get(req)
	if err != nil {
		return err
	}

	if err = c.responseJson(writer, customers, http.StatusOK); err != nil {
		return err
	}

	return nil
}

func (c CustomerController) GetCustomer(writer http.ResponseWriter, request *http.Request) error {
	req := &repositories.RequestId{}
	req.Id, _ = strconv.Atoi(mux.Vars(request)["id"])

	customer, err := c.repository.Customer().GetById(req)
	if err != nil {
		return err
	}

	if err = c.responseJson(writer, customer, http.StatusOK); err != nil {
		return err
	}

	return nil
}

func (c CustomerController) Create(writer http.ResponseWriter, request *http.Request) error {
	req := &models.Customer{}
	if err := json.NewDecoder(request.Body).Decode(req); err != nil {
		return err
	}

	if err := c.repository.Customer().Create(req); err != nil {
		return err
	}

	if err := c.responseJson(writer, "", http.StatusOK); err != nil {
		return err
	}

	return nil
}

func (c CustomerController) Edit(writer http.ResponseWriter, request *http.Request) error {
	req := &models.Customer{}
	req.ID, _ = strconv.Atoi(mux.Vars(request)["id"])
	if err := json.NewDecoder(request.Body).Decode(req); err != nil {
		return err
	}

	if err := c.repository.Customer().Edit(req); err != nil {
		return err
	}

	if err := c.responseJson(writer, "", http.StatusOK); err != nil {
		return err
	}

	return nil
}

func (c CustomerController) Delete(writer http.ResponseWriter, request *http.Request) error {
	req := &repositories.RequestId{}
	req.Id, _ = strconv.Atoi(mux.Vars(request)["id"])

	if err := c.repository.Customer().Delete(req); err != nil {
		return err
	}

	if err := c.responseJson(writer, "", http.StatusOK); err != nil {
		return err
	}

	return nil
}
