package models

import (
	"log"
	"os/exec"
)

type Module struct {
	Name          *string  `json:"name"`
	Path          *string  `json:"path"`
	FormatCommand *Command `json:"formatCommand"`
	Works         Works    `json:"works"`
	ActualWorks   []string
}

type Modules []*Module

func (modules Modules) Build(builder *Builder) {
	for i := range modules {
		modules[i].build(builder)
	}
}

func (s *Module) build(builder *Builder) {
	s.Works.Build(builder, s.Path)
	s.Format()
}

func (s *Module) setWorks(builder *Builder) {
	s.Format()
}

//func (s *Module) createModel(builder *Builder) {
//	path := fmt.Sprintf("%s/%s/%s.go", *s.Path, *s.Models, *builder.Data.Model.Names.Pascal)
//	builder.Build(config.GetTemplatePath("model"), path)
//}
//
//// CreateApi func
//func (s *Server) createApi(builder *Builder) {
//	apiPath := fmt.Sprintf("%s/%s/%s", *s.Path, *s.Api, *builder.Data.Model.Names.Camel)
//	err := os.Mkdir(apiPath, os.ModePerm)
//	if err != nil {
//		log.Fatal(err)
//	}
//	builder.Build(config.GetTemplatePath("init"), fmt.Sprintf("%s/init.go", apiPath))
//	builder.Build(config.GetTemplatePath("handler"), fmt.Sprintf("%s/handler.go", apiPath))
//	builder.Build(config.GetTemplatePath("service"), fmt.Sprintf("%s/service.go", apiPath))
//	builder.Build(config.GetTemplatePath("repository"), fmt.Sprintf("%s/repository.go", apiPath))
//}
//
//func (s *Server) createRouting(builder *Builder) {
//	routingPath := fmt.Sprintf("%s/%s/%s", *s.Path, *s.Routing, *builder.Data.Model.Names.Camel)
//	err := os.Mkdir(routingPath, os.ModePerm)
//	if err != nil {
//		log.Fatal(err)
//	}
//	builder.Build(config.GetTemplatePath("routing"), fmt.Sprintf("%s/init.go", routingPath))
//}
//
//func (s *Server) createManyService(builder *Builder) {
//	apiPath := fmt.Sprintf("%s/%s/%s", *s.Path, *s.Api, *builder.Data.Model.Names.Camel)
//	err := os.Mkdir(apiPath, os.ModePerm)
//	if err != nil {
//		log.Fatal(err)
//	}
//	builder.Build(config.GetTemplatePath("initMany"), fmt.Sprintf("%s/init.go", apiPath))
//	builder.Build(config.GetTemplatePath("serviceMany"), fmt.Sprintf("%s/handler.go", apiPath))
//	builder.Build(config.GetTemplatePath("repositoryMany"), fmt.Sprintf("%s/service.go", apiPath))
//}

func (s *Module) Format() {
	if s.FormatCommand == nil {
		return
	}
	_, err := exec.Command(*s.FormatCommand.Command, s.FormatCommand.Params...).Output()
	if err != nil {
		log.Fatal(err)
	}
}
