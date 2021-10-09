package models

import "templater/nameBuilder"

type Model struct {
	Name       *string
	Names      *nameBuilder.NameFormats
	Properties Properties
}

func (i *Model) SetNames() {
	i.Names = nameBuilder.GetNames(i.Name)
}
