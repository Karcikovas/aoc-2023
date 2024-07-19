package utils

import "strings"

func CheckSubstrings(str string, value string) int {
	matches := 0

	for _, char := range str {
		if strings.Contains(string(char), value) {
			matches += 1
		}
	}

	return matches
}
