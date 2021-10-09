package nameBuilder

import (
	"github.com/gertd/go-pluralize"
	"templater/nameBuilder/strcase"
)

type NameFormats struct {
	CamelSingular  *string
	PascalSingular *string
	SnakeSingular  *string
	CamelPlural    *string
	PascalPlural   *string
	SnakePlural    *string
}

func GetNames(name *string) *NameFormats {
	camel := strcase.ToLowerCamel(*name)
	pascal := strcase.ToCamel(*name)
	snake := strcase.ToSnake(*name)

	pluralizator := pluralize.NewClient()

	camelSingular := pluralizator.Singular(camel)
	pascalSingular := pluralizator.Singular(pascal)
	snakeSingular := pluralizator.Singular(snake)

	camelPlural := pluralizator.Plural(camel)
	pascalPlural := pluralizator.Plural(pascal)
	snakePlural := pluralizator.Plural(snake)

	names := NameFormats{
		CamelSingular:  &camelSingular,
		PascalSingular: &pascalSingular,
		SnakeSingular:  &snakeSingular,
		CamelPlural:    &camelPlural,
		PascalPlural:   &pascalPlural,
		SnakePlural:    &snakePlural,
	}
	return &names
}
