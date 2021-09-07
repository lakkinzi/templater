package repository

import (
	"templater/config"
	"templater/nameBuilder"
	"templater/templatesFabrics"
)

func Fabric(name *nameBuilder.NameFormats) *string {
	data := templatesFabrics.CreateData(name)
	return templatesFabrics.ParseTemplate(config.GetTemplatePath("repository"), data)
}
