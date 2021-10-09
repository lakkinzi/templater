package cli

import (
	"github.com/AlecAivazis/survey/v2"
	"log"
	"templater/models"
)

func GetOptions() *models.Project {
	project := models.CreateProject(writeProjectName())
	writeModel(project)
	for i := range project.Modules {
		writeWorks(project.Modules[i])
	}
	for {
		if !addOnceMorePropertyToModel() {
			break
		}
		prop := models.Property{}
		writePropertyName(&prop)
		writePropertyValueType(&prop)
		project.Builder.Data.Model.Properties = append(project.Builder.Data.Model.Properties, &prop)
	}
	return project
}

func addOnceMorePropertyToModel() bool {
	add := false
	prompt := &survey.Confirm{
		Message: "Add property?",
	}
	err := survey.AskOne(prompt, &add)
	if err != nil {
		log.Fatal(err)
	}
	return add
}
