package kernels

import (
	"database/sql"
	"github.com/AlDrac/wallister_test_project/app/api/configs"
	"github.com/sirupsen/logrus"
	"log"
)

type kernel struct {
	server *server
	port   string
}

func Initialise(config *configs.Config) (*kernel, error) {
	logger, err := initialiseLogger(config.Logger.Level)
	if err != nil {
		return nil, err
	}

	database, err := initialiseDatabase(config.Database.Url)
	if err != nil {
		return nil, err
	}
	defer closeDatabase(database, &err)

	return &kernel{
		server: initialiseServer(logger, database),
		port:   config.Http.Port,
	}, nil
}

func (kernel *kernel) Run() {
	if err := kernel.server.StartServer(kernel.port); err != nil {
		log.Fatal(err)
	}
}

func initialiseLogger(level string) (*logrus.Logger, error) {
	logger := logrus.New()
	loggerLevel, err := logrus.ParseLevel(level)
	if err != nil {
		log.Fatal(err)
	}

	logger.SetLevel(loggerLevel)
	return logger, nil
}

func initialiseDatabase(url string) (*sql.DB, error) {
	database, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}

func closeDatabase(database *sql.DB, err *error) {
	closeErr := database.Close()
	if *err == nil {
		*err = closeErr
	}
}
