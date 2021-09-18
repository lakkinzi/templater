package models

import (
	"encoding/json"
	"errors"
	"log"
	"templater/helpers"
	"templater/nameBuilder"
)

type Project struct {
	Name   *string `json:"name"`
	Mode   Mode
	Server *Server `json:"server"`
}

type Projects []*Project

func CreateProject(opt *Options) *Project {
	projects := Projects{}
	err := json.Unmarshal(helpers.ReadJson(), &projects)
	if err != nil {
		log.Fatal(err)
	}
	project := Project{}
	for _, proj := range projects {
		if *proj.Name == opt.ProjectName {
			project = *proj
			project.Mode = opt.Mode
			project.Server.ServerOperations = &opt.Operations.ServerOperations
			return &project
		}
	}
	log.Fatal(errors.New("project was not found in configuration"))
	return nil
}

func (p *Project) DoWork(name *nameBuilder.NameFormats) {
	if p.Mode == Builder {
		p.Server.build(name)
	}
}

func (p *Project) migrate() {

}
