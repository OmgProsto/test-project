package service

import (
	"encoding/json"
	"log"
	"os"
	"testproject/internal/cmd/entity"
)

type EntityModel struct {
}

func (e EntityModel) JsonToEntity(jsonString string) entity.Urls {
	var urls entity.Urls

	err := json.Unmarshal([]byte(jsonString), &urls)

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	return urls
}
