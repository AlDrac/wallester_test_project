package kernels

import (
	"database/sql"
	"net/http"

	"github.com/AlDrac/wallister_test_project/app/api/routers"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type server struct {
	logger *logrus.Logger
	router *routers.Router
}

func initialiseServer(logger *logrus.Logger, database *sql.DB) *server {
	server := &server{
		logger,
		routers.InitialiseRouter(),
	}

	return server
}

func (server *server) StartServer(port string) error {
	server.logger.Info("The server is running on port " + port)
	return http.ListenAndServe(":"+port, server)
}

func (server *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.router.ServeHTTP(writer, request)
}
