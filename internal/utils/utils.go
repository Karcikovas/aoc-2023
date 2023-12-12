package utils

import (
	"log"
	"os"
	"strconv"
)

func ReadInput(file string) string {
	rawFile, err := os.ReadFile(file)
	if err != nil {
		log.Panicln(err)
	}
	return string(rawFile)
}

func StringToInt(value string) (int, error) {
	intVar, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}

	return intVar, err
}
