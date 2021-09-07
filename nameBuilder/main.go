package nameBuilder

import "templater/nameBuilder/strcase"

type NameFormats struct {
	CamelCase  *string
	PascalCase *string
	SnakeCase  *string
}

func GetNames(name *string) *NameFormats {
	camelName := strcase.ToLowerCamel(*name)
	pascalName := strcase.ToCamel(*name)
	snakeName := strcase.ToSnake(*name)
	names := NameFormats{
		CamelCase:  &camelName,
		PascalCase: &pascalName,
		SnakeCase:  &snakeName,
	}
	return &names
}
