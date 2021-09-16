package helpers

import (
	"log"
	"os"
)

func WriteFile(path string, data *string) {
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
