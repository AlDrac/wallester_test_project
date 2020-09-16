package routers

import (
	"database/sql"
	"net/http"

	"github.com/AlDrac/wallister_test_project/app/api/controllers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Router struct {
	*mux.Router
	database *sql.DB
}

func InitialiseRouter(database *sql.DB) *Router {
	return &Router{
		mux.NewRouter(),
		database,
	}
}

func (router *Router) GetRouterHandlers() {
	customerController := controllers.Customer
	router.HandleFunc(
		"/customers}",
		customerController.Handler(customerController.Index),
	).Methods(http.MethodGet)
	router.HandleFunc(
		"/customer/{id:[0-9+]}",
		customerController.Handler(customerController.GetCustomer),
	).Methods(http.MethodGet)
}
