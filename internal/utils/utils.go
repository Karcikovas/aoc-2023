package utils

import (
	"log"
	"os"
	"regexp"
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

func GetIntArrayFromString(stringToConvert string) []int {
	var number []int

	re := regexp.MustCompile(`(\d+)`)
	values := re.FindAllString(stringToConvert, -1)

	for _, value := range values {
		intVal, _ := StringToInt(value)

		number = append(number, intVal)
	}

	return number
}
