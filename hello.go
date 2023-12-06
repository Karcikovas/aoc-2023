package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var totalSum int = 0
	content, error := os.ReadFile("./assets/day1input.txt")
	re := regexp.MustCompile(`(\d+)(,\d{0,2})?(\.\d{2})?`)

	if error != nil {
		log.Fatal(error)
	}

	words := strings.Split(string(content), "\n")

	for i := 0; i < len(words); i++ {
		str := words[i]
		values := re.FindAllString(str, -1)
		nmb1 := values[0]

		if len(values) == 1 {
			intVar, err := strconv.Atoi(nmb1 + nmb1)
			if err != nil {
				log.Fatal(err)
			}
			totalSum += intVar
		} else {
			nmb2 := values[len(values)-1]
			kebas := nmb1 + nmb2

			firstNumber := kebas[0:1]
			lastNumber := kebas[len(kebas)-1:]
			totalNum := firstNumber + lastNumber

			intVar, err := strconv.Atoi(totalNum)
			if err != nil {
				log.Fatal(err)
			}

			totalSum += intVar
		}
	}

	fmt.Println("Total sum of coordinates is:", totalSum)
}
