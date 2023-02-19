package print_service

import (
	"fmt"
	"strings"
)

type PrintResultService struct {
}

func (p PrintResultService) Print(url string, ok bool, invalidChecks []string) {
	status := "ok"
	invalidChecksString := ""

	if !ok {
		status = "fail"
		invalidChecksString = "(" + strings.Join(invalidChecks, ",") + ")"
	}

	fmt.Println(url, status, invalidChecksString)
}
