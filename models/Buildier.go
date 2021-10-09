package models

import (
	"log"
	"strings"
	"templater/helpers"
	"text/template"
)

type Builder struct {
	Data *Data
}

type Data struct {
	Model *Model
}

func CreateBuilder(model *Model) *Builder {
	builder := Builder{
		Data: createData(model),
	}
	return &builder
}

func (b *Builder) Build(templatePath *string, path string, fileName string) {
	res := b.parseTemplate(templatePath)
	helpers.WriteFile(path, res, &fileName)
}

func createData(model *Model) *Data {
	data := Data{
		Model: model,
	}
	return &data
}

func (b *Builder) parseTemplate(templatePath *string) *string {
	var buf strings.Builder
	tmpl, err := template.ParseFiles(*templatePath)
	if err != nil {
		log.Fatal(err)
	}
	_ = tmpl.Execute(&buf, b.Data)
	strTmpl := buf.String()
	return &strTmpl
}
