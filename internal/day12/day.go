package day12

import (
	"aoc2023/internal/utils"
	"log"
	"strings"
)

type Star struct {
	x int
	y int
}

func CalculatePart1() int {
	initialValues()

	return 0
}

func initialValues() {

	content := utils.ReadInput("./assets/day12input.txt")
	rows := strings.Split(content, "\n")

	for _, row := range rows {
		log.Println(row)
	}
}
