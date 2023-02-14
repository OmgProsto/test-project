package action

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testproject/internal/layer/cmd/domain/sitechecker/entity"
	persistenceEntity "testproject/internal/layer/cmd/persistence/sitechecker/entity"
	"testproject/internal/layer/cmd/persistence/sitechecker/model"
)

type FileParserAction struct {
	SiteModel model.SiteModel
}

func (f FileParserAction) Parse() []entity.Site {
	return f.SiteModel.ToDomain(f.parseFile())
}

func (f *FileParserAction) parseFile() persistenceEntity.Urls {
	file, err := os.Open("/home/user/GolandProjects/test-project/storage/sites.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)
	dataString := ""

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		dataString += string(data[:n])
	}

	var urls persistenceEntity.Urls

	err = json.Unmarshal([]byte(dataString), &urls)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return urls
}
