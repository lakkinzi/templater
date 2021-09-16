package config

import (
	"flag"
)

type Flags struct {
	Project *string
	Mode    *string
	Action  *string
	Name    *string
}

func ParseFlags() Flags {
	mode := flag.String("mode", "", "init/create/createSql/run/rollback")
	action := flag.String("action", "", "init/create/createSql/run/rollback")
	name := flag.String("name", "", "/init")
	project := flag.String("project", "", "")
	flag.Parse()
	return Flags{Project: project, Mode: mode, Action: action, Name: name}
}
