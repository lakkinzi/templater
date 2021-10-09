package models

import (
	"encoding/json"
	"errors"
	"log"
	"templater/helpers"
)

type Project struct {
	Name    *string `json:"name"`
	Mode    Mode
	Builder *Builder
	Modules Modules `json:"modules"`
}

type Projects []*Project

func CreateProject(name *string) *Project {
	projects := Projects{}
	err := json.Unmarshal(helpers.ReadJson(), &projects)
	if err != nil {
		log.Fatal(err)
	}
	project := Project{}
	for _, proj := range projects {
		if *proj.Name == *name {
			project = *proj
			return &project
		}
	}
	log.Fatal(errors.New("project was not found in configuration"))
	return nil
}

func (p *Project) DoWorks() {
	p.Modules.Build(p.Builder)
}

func (p *Project) migrate() {

}
