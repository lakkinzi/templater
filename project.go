package main

import (
	"encoding/json"
	"errors"
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

type Project struct {
	Name   *string `json:"name"`
	Server *Server
}

type Server struct {
	Path          *string  `json:"path"`
	Models        *string  `json:"models"`
	Api           *string  `json:"api"`
	Routing       *string  `json:"routing"`
	FormatCommand *Command `json:"formatCommand"`
}

type Command struct {
	Command *string
	Params  []string
}

type Projects []*Project

func CreateProject(name *string) *Project {
	projects := Projects{}
	err := json.Unmarshal(helpers.ReadJson(), &projects)
	if err != nil {
		log.Fatal(err)
	}
	for _, project := range projects {
		if *project.Name == *name {
			return project
		}
	}
	log.Fatal(errors.New("project was not found in configuration"))
	return nil
}

func (p *Project) CreateModel(name *nameBuilder.NameFormats) {
	res := model.Fabric(name)
	path := fmt.Sprintf("%s/%s/%s.go", *p.Server.Path, *p.Server.Models, *name.PascalCase)
	helpers.WriteFile(path, res)
}

// CreateApi func
func (p *Project) CreateApi(names *nameBuilder.NameFormats) {
	p.CreateModel(names)
	apiPath := fmt.Sprintf("%s/%s/%s", *p.Server.Path, *p.Server.Api, *names.CamelCase)
	err := os.Mkdir(apiPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	helpers.WriteFile(fmt.Sprintf("%s/init.go", apiPath), initFabric.Fabric(names))
	helpers.WriteFile(fmt.Sprintf("%s/handler.go", apiPath), handler.Fabric(names))
	helpers.WriteFile(fmt.Sprintf("%s/service.go", apiPath), service.Fabric(names))
	helpers.WriteFile(fmt.Sprintf("%s/repository.go", apiPath), repository.Fabric(names))

	err = os.Mkdir(fmt.Sprintf("%s/%s/%s", *p.Server.Path, *p.Server.Routing, *names.CamelCase), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	helpers.WriteFile(fmt.Sprintf("%s/%s/%s/init.go", *p.Server.Path, *p.Server.Routing, *names.CamelCase), routing.Fabric(names))
}

func (p *Project) CreateService(names *nameBuilder.NameFormats) {
	p.CreateModel(names)
	err := os.Mkdir(fmt.Sprintf("handlers/%s", *names.CamelCase), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	helpers.WriteFile(fmt.Sprintf("%s/%s/init.go", *p.Server.Path, *names.CamelCase), initMany.Fabric(names))
	helpers.WriteFile(fmt.Sprintf("%s/%s/handler.go", *p.Server.Path, *names.CamelCase), serviceMany.Fabric(names))
	helpers.WriteFile(fmt.Sprintf("%s/%s/service.go", *p.Server.Path, *names.CamelCase), repositoryMany.Fabric(names))
}

func (p *Project) Format() {
	_, err := exec.Command(*p.Server.FormatCommand.Command, p.Server.FormatCommand.Params...).Output()
	if err != nil {
		log.Fatal(err)
	}
}
