package config

import "path/filepath"

var fabricsDir = "./templatesFabrics/"

func GetTemplatePath(fabric string) *string {
	path := filepath.Join(fabric, "templates", "template.tmpl")
	fullPath := filepath.Join(fabricsDir, path)
	return &fullPath
}
