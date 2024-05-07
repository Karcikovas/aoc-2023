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

type Coordinates struct {
	y     int
	x     int
	value string
}

func CalculatePart1() int {
	pipes, start := initialValues()
	getPipeLoopItems(pipes, start)

	return 0
}

func getPipeLoopItems(pipes [][]string, start Coordinates) {
	var LoopItems []string
	var loopIsCompleted = false
	var location = start

	for !loopIsCompleted {
		//When whole loop is completed
		if len(LoopItems) != 0 && location.value == Start {
			LoopItems = append(LoopItems, location.value)
			loopIsCompleted = true

			break
		}

		log.Println(location)
		//Top check
		if location.y > 0 {
			topValue := pipes[location.y-1][location.x]

			log.Println(topValue)
		}

		//Bottom check
		if location.y < len(pipes) {
			newY := location.y + 1
			bottomValue := pipes[newY][location.x]

			log.Println(bottomValue)

			if bottomValue == "|" {
				location = Coordinates{
					y:     newY,
					x:     location.x,
					value: bottomValue,
				}

				continue
			}

			if bottomValue == "L" {
				log.Println(bottomValue)
				break
			}
		}

		loopIsCompleted = true
	}

	log.Println(len(pipes), start)

}

func initialValues() ([][]string, Coordinates) {
	var pipes [][]string
	var start Coordinates
	content := utils.ReadInput("./assets/day10input.txt")
	rows := strings.Split(string(content), "\n")

	for y, row := range rows {
		startIndex := strings.Index(row, "S")
		items := strings.Split(row, "")

		pipes = append(pipes, items)

		if startIndex >= 0 {
			start = Coordinates{
				y:     y,
				x:     startIndex,
				value: "S",
			}
		}
	}

	return pipes, start
}
