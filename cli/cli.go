package cli

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"log"
	"templater/config"
	"templater/models"
	"templater/nameBuilder"
)

func GetOptions() *models.Options {
	options := models.Options{}
	err := survey.AskOne(getProjectName(), &options.ProjectName)
	if err != nil {
		log.Fatal(err)
	}
	mode := ""
	err = survey.AskOne(getMode(), &mode)
	if err != nil {
		log.Fatal(err)
	}
	options.Mode = models.WriteMode(mode)
	if options.Mode == models.Builder {
		name := ""
		err = survey.AskOne(getName(), &name)
		options.Name = nameBuilder.GetNames(&name)

		var operations []string
		err = survey.AskOne(getServerBuildOperations(), &operations)
		fmt.Println(operations)
		options.Operations.ServerOperations.ServerBuildOperations = models.StringsToServerBuildOperations(operations)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &options
}

func getName() *survey.Input {
	return &survey.Input{Message: "Choose entiry name:"}
}

func getProjectName() *survey.Select {
	return &survey.Select{Message: "Choose Project:", Options: config.GetProjectsNames()}
}

func getServerBuildOperations() *survey.MultiSelect {
	return &survey.MultiSelect{Message: "Choose Server operations:", Options: models.GetServerBuildOperations()}
}

func getMode() *survey.Select {
	return &survey.Select{Message: "Choose mode:", Options: models.GetModels()}
}
