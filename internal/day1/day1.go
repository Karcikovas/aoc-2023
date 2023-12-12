package day1

import (
	"aoc2023/internal/utils"
	"regexp"
	"strings"
)

func CalculatePart1() int {
	var totalSum int = 0

	content := utils.ReadInput("./../../assets/day1input.txt")
	re := regexp.MustCompile(`(\d)|,(\d{0,2})?(\.\d{2})?`)

	words := strings.Split(string(content), "\n")

	for i := 0; i < len(words); i++ {
		str := words[i]
		values := re.FindAllString(str, -1)
		nmb1 := values[0]

		if len(values) == 1 {
			intVar, _ := utils.StringToInt(nmb1 + nmb1)
			totalSum += intVar
		} else {
			nmb2 := values[len(values)-1]
			kebas := nmb1 + nmb2

			firstNumber := kebas[0:1]
			lastNumber := kebas[len(kebas)-1:]
			totalNum := firstNumber + lastNumber

			intVar, _ := utils.StringToInt(totalNum)
			totalSum += intVar
		}
	}

	return totalSum
}
