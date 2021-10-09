package models

import (
	"fmt"
	"path/filepath"
	"strings"
	"templater/config"
)

type Work struct {
	Name      string `json:"name"`
	Extension string `json:"extension"`
	Path      string `json:"path"`
	Case      string `json:"case"`
	Template  string `json:"template"`
	Doing     bool
	FileName  string `json:"fileName"`
}

type Works []*Work

func (works Works) GetNames() []string {
	names := make([]string, 0)
	for _, work := range works {
		names = append(names, work.Name)
	}
	return names
}

func (works Works) SetDoing(actualWorks []string) {
	for i := range works {
		for _, actualWork := range actualWorks {
			if actualWork == works[i].Name {
				works[i].Doing = true
				break
			}
		}
	}
}

func (work *Work) build(builder *Builder, projectPath *string) {
	work.setFileName(builder)
	work.Path = strings.Replace(work.Path, "%model", *builder.Data.Model.Names.CamelSingular, -1)
	path := filepath.Join(*projectPath, work.Path)

	builder.Build(config.GetTemplatePath(work.Name), path, work.FileName)
}

func (work *Work) setFileName(builder *Builder) {
	if work.FileName != "%model" {
		work.FileName = fmt.Sprintf("%s.%s", work.Name, work.Extension)
	} else {
		fileName := strings.Replace(work.FileName, "%model", *builder.Data.Model.Names.PascalSingular, -1)
		work.FileName = fmt.Sprintf("%s.%s", fileName, work.Extension)
	}

}

func (works Works) Build(builder *Builder, projectPath *string) {
	for _, work := range works {
		if work.Doing {
			work.build(builder, projectPath)
		}
	}
}
