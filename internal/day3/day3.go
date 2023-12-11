package day3

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func CalculatePart1() {
	totalSum := 0
	content, err := os.ReadFile("./../../assets/day3input.txt")
	re := regexp.MustCompile(`\d+`)

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	for i := 0; i < len(lines); i++ {
		var prevLine string
		var nextLine string

		currentLine := lines[i]
		numbers := re.FindAllString(currentLine, -1)

		if i > 0 {
			prevLine = lines[i-1]
		}

		if i+1 < len(lines) {
			nextLine = lines[i+1]
		}

		for _, number := range numbers {
			var startFrom int
			var endFrom int

			intVar, _ := stringToInt(number)
			startIndex := strings.Index(currentLine, number)
			endIndex := startIndex + len(number) - 1

			if strings.Index(currentLine, number) > 0 {
				startFrom = startIndex - 1
			} else {
				startFrom = startIndex
			}

			if endIndex+1 < len(currentLine) {
				endFrom = endIndex + 1
			} else {
				endFrom = endIndex
			}

			symbolSF := currentLine[startFrom]
			symbolEF := currentLine[endFrom]

			if (startFrom > 0 && symbolSF != '.') || (endFrom < len(currentLine) && symbolEF != '.') {
				totalSum += intVar
			} else {
				for b := startFrom; b <= endFrom; b++ {
					if len(prevLine) > 0 {

						symbol := string(prevLine[b])

						if symbol != "." {

							totalSum += intVar
							break
						}
					}
					if len(nextLine) > 0 {
						symbol := string(nextLine[b])

						if symbol != "." {

							totalSum += intVar
							break
						}
					}
				}
			}
		}
	}

	log.Println("Total sum:", totalSum)
}

func stringToInt(value string) (int, error) {
	intVar, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}

	return intVar, err
}
