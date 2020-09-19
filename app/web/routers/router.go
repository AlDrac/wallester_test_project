package routers

import (
	"flag"
	template_service "github.com/AlDrac/wallister_test_project/app/web/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

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

	router.HandleFunc("/", indexHandler).Methods(http.MethodGet)
	router.HandleFunc("/customers", customersHandler).Methods(http.MethodGet)

	router.HandleFunc("/customer/create", customerEditHandler).Methods(http.MethodGet)
	router.HandleFunc("/customer/create", customerPostCreateHandler).Methods(http.MethodPost)

	router.HandleFunc("/customer/{id:[0-9]+}", customerViewHandler).Methods(http.MethodGet)
	router.HandleFunc("/customer/edit/{id:[0-9]+}", customerEditHandler).Methods(http.MethodGet)
	router.HandleFunc("/customer/edit/{id:[0-9]+}", customerPostEditHandler).Methods(http.MethodPost)
	router.HandleFunc("/customer/delete/{id:[0-9]+}", customerViewHandler).Methods(http.MethodGet)

	router.NotFoundHandler = http.HandlerFunc(pageNotFoundHandler)
}

func pageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	err := template_service.RenderTemplate(w, "404.tmpl", TemplateData{
		Page: "404",
	})
	if err != nil {
		log.Println(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := template_service.RenderTemplate(w, "index.tmpl", TemplateData{
		Page: "home",
	})
	if err != nil {
		log.Println(err)
	}
}

func customersHandler(w http.ResponseWriter, r *http.Request) {

	err := template_service.RenderTemplate(w, "customers.tmpl", TemplateData{
		Page: "customers",
	})
	if err != nil {
		log.Println(err)
	}
}

func customerCreateHandler(w http.ResponseWriter, r *http.Request) {

}

func customerPostCreateHandler(w http.ResponseWriter, r *http.Request) {
	customersHandler(w, r)
}

func customerViewHandler(w http.ResponseWriter, r *http.Request) {

}

func customerEditHandler(w http.ResponseWriter, r *http.Request) {

}

func customerPostEditHandler(w http.ResponseWriter, r *http.Request) {

	customerViewHandler(w, r)
}

func customerDeleteHandler(w http.ResponseWriter, r *http.Request) {

	customerViewHandler(w, r)
}
