package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	content, error := os.ReadFile("./assets/day1input.txt")
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	if error != nil {
		log.Fatal(error)
	}

	words := strings.Split(string(content), "\n")

	for i := 0; i < len(words); i++ {
		var str string = words[i]
		values := re.FindAllString(str, -1)

		if len(values) == 1 {
			nmb1 := values[0]
			fmt.Println(nmb1 + nmb1)
		}
	}

}
