package models

import "templater/nameBuilder"

type Options struct {
	ProjectName string
	Mode        Mode
	Operations  Operations
	Name        *nameBuilder.NameFormats
}
