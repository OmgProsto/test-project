package json_service

import (
	"io"
	"log"
	"os"
)

type FileParser struct {
}

func (f FileParser) Parse(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)

	dataString := ""
	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}

		dataString += string(data[:n])
	}

	return dataString
}
