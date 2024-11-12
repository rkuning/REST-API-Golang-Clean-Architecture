package util

import (
	"regexp"
	"strings"
)

func ExactError(error error) string {
	regex := regexp.MustCompile(`Error:(.*?)(\n|$)`)
	matches := regex.FindAllStringSubmatch(error.Error(), -1)

	var result []string

	for _, match := range matches {
		if len(match) > 1 {
			result = append(result, strings.Trim(match[1], " "))
		}
	}

	return strings.Join(result, ", ")
}
