package init

import (
	"templater/config"
	"templater/nameBuilder"
	"templater/templatesFabrics"
)

type Data struct {
	Model   string
	Package string
}

func Fabric(name *nameBuilder.NameFormats) *string {
	data := templatesFabrics.CreateData(name)
	return templatesFabrics.ParseTemplate(config.GetTemplatePath("init"), data)
}
