package routers

import "github.com/gorilla/mux"

type Router struct {
	*mux.Router
}

func InitialiseRouter() *Router {
	return &Router{
		mux.NewRouter(),
	}
}
