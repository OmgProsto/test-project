package service

import (
	"fmt"
	"strings"
)

type PrintResultChecksService struct {
}

func (p PrintResultChecksService) Print(url string, ok bool, invalidChecks []string) {
	status := "ok"
	invalidChecksString := ""

	if !ok {
		status = "fail"
		invalidChecksString = "(" + strings.Join(invalidChecks, ",") + ")"
	}

	fmt.Println(url, status, invalidChecksString)
}
