package kernels

import (
	"database/sql"
	"net/http"

	"github.com/AlDrac/wallister_test_project/app/api/routers"
	"github.com/sirupsen/logrus"
)

type server struct {
	db     *sql.DB
	logger *logrus.Logger
	router *routers.Router
}

func initialiseServer(logger *logrus.Logger, db *sql.DB) *server {
	server := &server{
		db,
		logger,
		routers.InitialiseRouter(db),
	}

	server.router.GetRouterHandlers()

	return server
}

func (server *server) StartServer(port string) error {
	defer server.db.Close()
	server.logger.Info("The server is running on port " + port)
	return http.ListenAndServe(":"+port, server)
}

func (server *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.router.ServeHTTP(writer, request)
}
