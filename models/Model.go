package models

import (
	"strings"
	"templater/nameBuilder"
	"time"
)

type Model struct {
	Name       *string
	Names      *nameBuilder.NameFormats
	Properties Properties
}

func (i *Model) SetNames() {
	i.Names = nameBuilder.GetNames(i.Name)
}

const (
	ModelPlaceholderPascalSingular = "%Model"
	ModelPlaceholderCamelSingular  = "%model"
	ModelPlaceholderSnakeSingular  = "%_model"
)

const (
	PlaceholderCurrentTimestamp = "%timestamp"
)

func (i *Model) SetNameToPlaceholder(name *string) {
	*name = strings.Replace(*name, ModelPlaceholderPascalSingular, *i.Names.PascalSingular, -1)
	*name = strings.Replace(*name, ModelPlaceholderCamelSingular, *i.Names.CamelSingular, -1)
	*name = strings.Replace(*name, ModelPlaceholderSnakeSingular, *i.Names.SnakeSingular, -1)
	*name = strings.Replace(*name, PlaceholderCurrentTimestamp, time.Now().Format("20060102150405"), -1)
}
