package json_service

import (
	"encoding/json"
	"log"
	"os"
	"testproject/pkg/json/json_service/entity"
)

type UrlModel struct {
}

func (e UrlModel) JsonToEntity(jsonString string) entity.Urls {
	var urls entity.Urls

	err := json.Unmarshal([]byte(jsonString), &urls)

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	return urls
}
