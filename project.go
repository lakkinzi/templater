package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type Project struct {
	Name       *string `json:"name"`
	ClientPath *string `json:"clientPath"`
	ServerPath *string `json:"serverPath"`
}

type Projects []*Project

func CreateProject(name *string) *Project {
	projects := Projects{}
	err := json.Unmarshal(readJson(), &projects)
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

func readJson() []byte {
	data, err := os.ReadFile("./.projects.json")
	if err != nil {
		log.Fatal(err)
	}
	return data
}
