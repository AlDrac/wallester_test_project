package kernels

import (
	"github.com/AlDrac/wallister_test_project/app/web/routers"
	"net/http"
)

type server struct {
	router *routers.Router
}

func initialiseServer() *server {
	server := &server{
		routers.InitialiseRouter(),
	}

	server.router.GetRouterHandlers()

	return server
}

func (server *server) StartServer(port string) error {
	return http.ListenAndServe(":"+port, server)
}

func (server *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.router.ServeHTTP(writer, request)
}
