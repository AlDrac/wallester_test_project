package kernels

import (
	"github.com/AlDrac/wallister_test_project/app/web/configs"
	"github.com/AlDrac/wallister_test_project/app/web/services"
	"log"
)

type kernel struct {
	server *server
	port   string
}

func Initialise(config *configs.Config) *kernel {
	template_service.SetTemplateConfig(config.Template.Layout, config.Template.Include)
	if err := template_service.LoadTemplates(); err != nil {
		log.Fatal(err)
	}

	return &kernel{
		server: initialiseServer(),
		port:   config.Http.Port,
	}
}

func (kernel *kernel) Run() {
	if err := kernel.server.StartServer(kernel.port); err != nil {
		log.Fatal(err)
	}
}
