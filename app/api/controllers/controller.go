package controllers

import (
	"github.com/sirupsen/logrus"
	"io"
	"net/http"

	"github.com/AlDrac/wallister_test_project/app/api/repositories"
	"github.com/darahayes/go-boom"
)

type Action func(writer http.ResponseWriter, request *http.Request) error

type Controller struct {
	repository repositories.Repository
	logger     *logrus.Logger
}

func InitialiseController(repository repositories.Repository, logger *logrus.Logger) Controller {
	return Controller{
		repository: repository,
		logger:     logger,
	}
}

func (c *Controller) responseJson(writer http.ResponseWriter, body string) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err := io.WriteString(writer, body)
	if err != nil {
		return err
	}

	return nil
}

func (c *Controller) Handler(action Action) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if err := action(writer, request); err != nil {
			c.logger.Error(err.Error())
			if err == repositories.ErrRecordNotFound {
				boom.NotFound(writer, err)
			} else {
				boom.Internal(writer, err)
			}
		}
	})
}
