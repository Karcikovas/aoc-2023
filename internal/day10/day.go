package day10

import (
	"aoc2023/internal/utils"
	"log"
	"strings"
)

const (
	VerticalPipe   = "|"
	HorizontalPipe = "F"
	NEBend         = "L"
	NWBend         = "J"
	SWBend         = "7"
	SEBend         = "F"
	Start          = "S"
)

type StartCoordinates struct {
	y int
	x int
}

func CalculatePart1() int {
	pipes, start := initialValues()

	log.Println(pipes, start)
	return 0
}

func initialValues() ([][]string, StartCoordinates) {
	var pipes [][]string
	var start StartCoordinates
	content := utils.ReadInput("./assets/day10input.txt")
	rows := strings.Split(string(content), "\n")

	for y, row := range rows {
		startIndex := strings.Index(row, "S")
		items := strings.Split(row, "")

		pipes = append(pipes, items)

		if startIndex >= 0 {
			start = StartCoordinates{
				y: y,
				x: startIndex,
			}
		}
	}

	return pipes, start
}
