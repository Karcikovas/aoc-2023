package day6

import (
	"aoc2023/internal/utils"
	"log"
	"strings"
)

func CalculatePart1() int {
	initialValues()
	return 1
}

type Race struct {
	RaceNumber int
	Duration   int
	Record     int
}

type Races struct {
	Races []Race
}

func initialValues() {
	races := &Races{}

	content := utils.ReadInput("./assets/day6input.txt")
	rows := strings.Split(string(content), "\n")

	for _, row := range rows {
		log.Print(row)
	}

	log.Println(content, races)
}
