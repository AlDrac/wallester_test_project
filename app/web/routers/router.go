package routers

import (
	"flag"
	"fmt"
	template_service "github.com/AlDrac/wallister_test_project/app/web/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type ActionHandler func(writer http.ResponseWriter, request *http.Request) error

type Router struct {
	*mux.Router
}

type Messages struct {
	Errors  []string
	Success string
}

type TemplateData struct {
	Page string
	Data map[string]interface{}
}

var (
	dir      string
	messages *Messages
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

	router.HandleFunc("/customer/create", middlewareHandler(customerEditHandler)).Methods(http.MethodGet)
	router.HandleFunc("/customer/create", middlewareHandler(customerPostCreateHandler)).Methods(http.MethodPost)

	router.HandleFunc("/customer/{id:[0-9]+}", middlewareHandler(customerViewHandler)).Methods(http.MethodGet)
	router.HandleFunc("/customer/edit/{id:[0-9]+}", middlewareHandler(customerEditHandler)).Methods(http.MethodGet)
	router.HandleFunc("/customer/edit/{id:[0-9]+}", middlewareHandler(customerPostEditHandler)).Methods(http.MethodPost)
	router.HandleFunc("/customer/delete/{id:[0-9]+}", middlewareHandler(customerViewHandler)).Methods(http.MethodGet)

	router.NotFoundHandler = middlewareHandler(pageNotFoundHandler)
}

func middlewareHandler(action ActionHandler) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if err := action(writer, request); err != nil {
			fmt.Println(err)
			pageInternalServerErrorHandler(writer, request)
		}
	})
}

func pageInternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	err := template_service.RenderTemplate(w, "500.tmpl", TemplateData{
		Page: "500",
	})
	if err != nil {
		log.Println(err)
	}
}

func pageNotFoundHandler(w http.ResponseWriter, r *http.Request) error {
	err := template_service.RenderTemplate(w, "404.tmpl", TemplateData{
		Page: "404",
	})
	if err != nil {
		return err
	}

	return nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) error {
	err := template_service.RenderTemplate(w, "index.tmpl", TemplateData{
		Page: "home",
	})
	if err != nil {
		return err
	}

	return nil
}

func customersHandler(w http.ResponseWriter, r *http.Request) error {

	err := template_service.RenderTemplate(w, "customers.tmpl", TemplateData{
		Page: "customers",
	})
	if err != nil {
		return err
	}

	return nil
}

func customerCreateHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func customerPostCreateHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func customerViewHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func customerEditHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func customerPostEditHandler(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func customerDeleteHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}
