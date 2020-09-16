package routers

import (
	"database/sql"
	"net/http"

	"github.com/AlDrac/wallister_test_project/app/api/controllers"
	"github.com/AlDrac/wallister_test_project/app/api/repositories/postgres"
	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
	controllers.Controller
}

func InitialiseRouter(db *sql.DB) *Router {
	return &Router{
		mux.NewRouter(),
		controllers.InitialiseController(
			postgres.InitialiseRepository(db),
		),
	}
}

func (router *Router) GetRouterHandlers() {
	customerController := controllers.Customer
	customerController.Controller = router.Controller

	router.HandleFunc(
		"/customers",
		customerController.Handler(customerController.Index),
	).Methods(http.MethodGet)
	router.HandleFunc(
		"/customer/{id:[0-9+]}",
		customerController.Handler(customerController.GetCustomer),
	).Methods(http.MethodGet)
}
