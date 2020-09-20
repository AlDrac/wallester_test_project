package controllers

import (
	"net/http"

	"github.com/AlDrac/wallister_test_project/app/api/repositories"
	"github.com/darahayes/go-boom"
	"github.com/liamylian/jsontime"
	"github.com/sirupsen/logrus"
)

type Action func(writer http.ResponseWriter, request *http.Request) error

type Controller struct {
	repository repositories.Repository
	logger     *logrus.Logger
}

type ResStruct struct {
	Error		string		`json:"error"`
	Message		string		`json:"message"`
	StatusCode	int			`json:"statusCode"`
	Result		interface{}	`json:"result"`
}

var json = jsontime.ConfigWithCustomTimeFormat

func InitialiseController(repository repositories.Repository, logger *logrus.Logger) Controller {
	return Controller{
		repository: repository,
		logger:     logger,
	}
}

func (c *Controller) responseJson(writer http.ResponseWriter, body interface{}, code int) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	if err := json.NewEncoder(writer).Encode(body); err != nil {
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
			} else if err == repositories.ErrRecordExist {
				boom.NotAcceptable(writer, err)
			} else {
				boom.Internal(writer, err)
			}
		}
	})
}
