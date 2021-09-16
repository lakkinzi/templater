package templatesFabrics

import (
	"log"
	"strings"
	"templater/nameBuilder"
	"text/template"
)

type Data struct {
	Model   *string
	Package *string
	Snake   *string
}

func CreateData(name *nameBuilder.NameFormats) *Data {
	data := Data{
		Model:   name.PascalCase,
		Package: name.CamelCase,
		Snake:   name.SnakeCase,
	}
	return &data
}

func ParseTemplate(templatePath *string, data *Data) *string {
	var buf strings.Builder
	tmpl, err := template.ParseFiles(*templatePath)
	if err != nil {
		log.Fatal(err)
	}
	_ = tmpl.Execute(&buf, *data)
	strTmpl := buf.String()
	return &strTmpl
}
