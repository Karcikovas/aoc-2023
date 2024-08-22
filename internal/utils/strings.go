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

func IndexOf(str string, value string) int {

	for index, char := range str {
		if string(char) == value {
			return index
		}
	}

	return -1
}
