package routers

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/AlDrac/wallister_test_project/app/api/controllers"
	"github.com/AlDrac/wallister_test_project/app/api/repositories/postgres"
	"github.com/gorilla/mux"
)

const version = "/v1"

type Router struct {
	*mux.Router
	controllers.Controller
}

func InitialiseRouter(db *sql.DB, logger *logrus.Logger) *Router {
	return &Router{
		mux.NewRouter(),
		controllers.InitialiseController(
			postgres.InitialiseRepository(db),
			logger,
		),
	}
}

func (router *Router) GetRouterHandlers() {
	customerController := controllers.Customer
	customerController.Controller = router.Controller

	router.HandleFunc(
		version+"/customers",
		customerController.Handler(customerController.GetCustomers),
	).Methods(http.MethodGet)
	router.HandleFunc(
		version+"/customer/{id:[0-9]+}",
		customerController.Handler(customerController.GetCustomer),
	).Methods(http.MethodGet)
	router.HandleFunc(
		version+"/customer/create",
		customerController.Handler(customerController.Create),
	).Methods(http.MethodPost)
	router.HandleFunc(
		version+"/customer/edit/{id:[0-9]+}",
		customerController.Handler(customerController.Edit),
	).Methods(http.MethodPut)
	router.HandleFunc(
		version+"/customer/delete/{id:[0-9]+}",
		customerController.Handler(customerController.Delete),
	).Methods(http.MethodDelete)

	notFoundController := controllers.NotFound
	notFoundController.Controller = router.Controller
	router.NotFoundHandler = notFoundController.Handler(notFoundController.Index)
}
