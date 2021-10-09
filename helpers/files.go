package helpers

import (
	"log"
	"os"
	"path/filepath"
)

func WriteFile(path string, data *string, fileName *string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	path = filepath.Join(path, *fileName)
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.WriteString(*data)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadJson() []byte {
	data, err := os.ReadFile("./.projects.json")
	if err != nil {
		log.Fatal(err)
	}
	return data
}
