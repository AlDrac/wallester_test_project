package routers

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	serviceTemplate "github.com/AlDrac/wallister_test_project/app/web/services/template"
	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

type TemplateData struct {
	Page string
	Data map[string]interface{}
}

var (
	dir string
)

func InitialiseRouter() *Router {
	flag.StringVar(&dir, "dir", "./static", "the directory to serve files from.")
	flag.Parse()

	return &Router{
		mux.NewRouter(),
	}
}

func (router *Router) GetRouterHandlers() {
	router.PathPrefix("/asserts/").Handler(http.StripPrefix("/asserts/", http.FileServer(http.Dir(dir))))

	router.HandleFunc("/", middlewareHandler(indexHandler)).Methods(http.MethodGet)
	router.HandleFunc("/customers", middlewareHandler(customersHandler)).Methods(http.MethodGet)

	router.HandleFunc("/customer/create", middlewareHandler(customerCreateHandler)).Methods(http.MethodGet)
	router.HandleFunc("/customer/create", middlewareHandler(customerPostCreateHandler)).Methods(http.MethodPost)

	router.HandleFunc("/customer/{id:[0-9]+}", middlewareHandler(customerViewHandler)).Methods(http.MethodGet)

	router.HandleFunc("/customer/edit/{id:[0-9]+}", middlewareHandler(customerEditHandler)).Methods(http.MethodGet)
	router.HandleFunc("/customer/edit/{id:[0-9]+}", middlewareHandler(customerPostEditHandler)).Methods(http.MethodPost)

	router.HandleFunc("/customer/delete/{id:[0-9]+}", middlewareHandler(customerDeleteHandler)).Methods(http.MethodPost)

	router.NotFoundHandler = middlewareHandler(pageNotFoundHandler)
}

type ActionHandler func(writer http.ResponseWriter, request *http.Request) error

func middlewareHandler(action ActionHandler) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if err := action(writer, request); err != nil {
			fmt.Println(err)
			pageInternalServerErrorHandler(writer, request)
		}
	})
}

func pageInternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	err := serviceTemplate.RenderTemplate(w, "500.tmpl", TemplateData{
		Page: "500",
	})
	if err != nil {
		log.Println(err)
	}
}

func pageNotFoundHandler(w http.ResponseWriter, r *http.Request) error {
	err := serviceTemplate.RenderTemplate(w, "404.tmpl", TemplateData{
		Page: "404",
	})
	if err != nil {
		return err
	}

	return nil
}
