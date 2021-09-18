package config

import (
	"encoding/json"
	"log"
	"templater/helpers"
)

type NameStruct struct {
	Name string `json:"name"`
}

func GetProjectsNames() []string {
	namesStructs := make([]*NameStruct, 0)
	err := json.Unmarshal(helpers.ReadJson(), &namesStructs)
	if err != nil {
		log.Fatal(err)
	}
	names := make([]string, 0)
	for _, name := range namesStructs {
		names = append(names, name.Name)
	}
	return names
}
