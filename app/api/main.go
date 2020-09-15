package main

import (
	"log"
	"net/http"

	"github.com/AlDrac/wallister_test_project/app/api/configs"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	config *configs.Config
	logger *logrus.Logger
	router *mux.Router
}

func main() {
	config := configs.New()
	apiServer := &ApiServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}

	err := apiServer.start()
	if err != nil {
		log.Fatal(err)
	}
}

func (apiServer *ApiServer) start() error {
	err := apiServer.configureLogger()
	if err != nil {
		log.Fatal(err)
	}

	apiServer.configureRouter()

	apiServer.logger.Info("Starting api server")
	return http.ListenAndServe(":"+apiServer.config.Http.Port, apiServer.router)
}

func (apiServer *ApiServer) configureLogger() error {
	level, err := logrus.ParseLevel(apiServer.config.Logger.Level)
	if err != nil {
		log.Fatal(err)
	}

	apiServer.logger.SetLevel(level)
	return nil
}

func (apiServer *ApiServer) configureRouter() {

}
