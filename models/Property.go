package models

import "templater/nameBuilder"

type Property struct {
	Name      *string
	Names     *nameBuilder.NameFormats
	ValueType ValueType `json:"valueType"`
}

type Properties []*Property

func (p *Property) SetValueType(valueType string) {
	switch valueType {
	case "String":
		p.ValueType = String
	case "Number":
		p.ValueType = Number
	case "Date":
		p.ValueType = Date
	}
}

func (p *Property) SetNames() {
	p.Names = nameBuilder.GetNames(p.Name)
}
