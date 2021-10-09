package cli

import (
	"github.com/AlecAivazis/survey/v2"
	"log"
	"templater/models"
)

func writePropertyName(prop *models.Property) {
	name := ""
	err := survey.AskOne(getPropertyName(), &name)
	if err != nil {
		log.Fatal(err)
	}
	prop.Name = &name
	prop.SetNames()
}

func writePropertyValueType(prop *models.Property) {
	valueType := ""
	err := survey.AskOne(getPropertyValueType(), &valueType)
	if err != nil {
		log.Fatal(err)
	}
	prop.SetValueType(valueType)
}

func writeProjectName() *string {
	projectName := ""
	err := survey.AskOne(getProjectName(), &projectName)
	if err != nil {
		log.Fatal(err)
	}
	return &projectName
}

func writeMode(options *models.Options) {
	mode := ""
	err := survey.AskOne(getMode(), &mode)
	if err != nil {
		log.Fatal(err)
	}
	options.Mode = models.WriteMode(mode)
}

func writeModel(project *models.Project) {
	model := models.Model{}
	name := ""
	err := survey.AskOne(getName(), &name)
	if err != nil {
		log.Fatal(err)
	}
	model.Name = &name
	model.SetNames()
	project.Builder = models.CreateBuilder(&model)
}

func writeServerBuildOperations(options *models.Options) {
	var operations []string
	err := survey.AskOne(getServerBuildOperations(), &operations)
	options.Operations.ServerOperations.ServerBuildOperations = models.StringsToServerBuildOperations(operations)
	if err != nil {
		log.Fatal(err)
	}
}

func writeWorks(module *models.Module) {
	var works []string
	err := survey.AskOne(getWorks(module.Works), &works)
	if err != nil {
		log.Fatal(err)
	}
	module.Works.SetDoing(works)
}
