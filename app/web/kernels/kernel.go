package kernels

import (
	"github.com/AlDrac/wallister_test_project/app/web/configs"
	"github.com/AlDrac/wallister_test_project/app/web/routers"
	serviceApi "github.com/AlDrac/wallister_test_project/app/web/services/api"
	serviceTemplate "github.com/AlDrac/wallister_test_project/app/web/services/template"
	"html/template"
	"log"
	"time"
)

type kernel struct {
	server *server
	port   string
}

func Initialise(config *configs.Config) *kernel {
	serviceTemplate.SetTemplateConfig(config.Template.Layout, config.Template.Include)
	serviceTemplate.SetTemplateFunction(&serviceTemplate.TemplateFuncMap{
		FM: template.FuncMap{
			"now": time.Now,
		},
	})

	if err := serviceTemplate.LoadTemplates(); err != nil {
		log.Fatal(err)
	}

	routers.SetStore(config.Session.Key)
	serviceApi.InitializeServiceApi(config.Http.ApiUrl)

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
