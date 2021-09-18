package models

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"templater/helpers"
	"templater/nameBuilder"
	"templater/templatesFabrics/handler"
	initFabric "templater/templatesFabrics/init"
	"templater/templatesFabrics/initMany"
	"templater/templatesFabrics/model"
	"templater/templatesFabrics/repository"
	"templater/templatesFabrics/repositoryMany"
	"templater/templatesFabrics/routing"
	"templater/templatesFabrics/service"
	"templater/templatesFabrics/serviceMany"
)

type Server struct {
	Path             *string  `json:"path"`
	Models           *string  `json:"models"`
	Api              *string  `json:"api"`
	Routing          *string  `json:"routing"`
	FormatCommand    *Command `json:"formatCommand"`
	ServerOperations *ServerOperations
}

func (s *Server) build(name *nameBuilder.NameFormats) {
	for _, operation := range s.ServerOperations.ServerBuildOperations {
		if operation == CreateModel {
			s.createModel(name)
		}
		if operation == CreateApi {
			s.createApi(name)
		}
		if operation == CreateRouting {
			s.createRouting(name)
		}
		if operation == CreateServiceMany {
			s.createManyService(name)
		}
	}
	s.Format()
}

func (s *Server) createModel(name *nameBuilder.NameFormats) {
	res := model.Fabric(name)
	path := fmt.Sprintf("%s/%s/%s.go", *s.Path, *s.Models, *name.PascalCase)
	helpers.WriteFile(path, res)
}

// CreateApi func
func (s *Server) createApi(names *nameBuilder.NameFormats) {
	apiPath := fmt.Sprintf("%s/%s/%s", *s.Path, *s.Api, *names.CamelCase)
	err := os.Mkdir(apiPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	helpers.WriteFile(fmt.Sprintf("%s/init.go", apiPath), initFabric.Fabric(names))
	helpers.WriteFile(fmt.Sprintf("%s/handler.go", apiPath), handler.Fabric(names))
	helpers.WriteFile(fmt.Sprintf("%s/service.go", apiPath), service.Fabric(names))
	helpers.WriteFile(fmt.Sprintf("%s/repository.go", apiPath), repository.Fabric(names))

	err = os.Mkdir(fmt.Sprintf("%s/%s/%s", *s.Path, *s.Routing, *names.CamelCase), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) createRouting(name *nameBuilder.NameFormats) {
	helpers.WriteFile(fmt.Sprintf("%s/%s/%s/init.go", *s.Path, *s.Routing, *name.CamelCase), routing.Fabric(name))
}

func (s *Server) createManyService(names *nameBuilder.NameFormats) {
	err := os.Mkdir(fmt.Sprintf("handlers/%s", *names.CamelCase), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	helpers.WriteFile(fmt.Sprintf("%s/%s/init.go", *s.Path, *names.CamelCase), initMany.Fabric(names))
	helpers.WriteFile(fmt.Sprintf("%s/%s/handler.go", *s.Path, *names.CamelCase), serviceMany.Fabric(names))
	helpers.WriteFile(fmt.Sprintf("%s/%s/service.go", *s.Path, *names.CamelCase), repositoryMany.Fabric(names))
}

func (s *Server) Format() {
	_, err := exec.Command(*s.FormatCommand.Command, s.FormatCommand.Params...).Output()
	if err != nil {
		log.Fatal(err)
	}
}
