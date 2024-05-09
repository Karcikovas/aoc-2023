package day10

import (
	"aoc2023/internal/utils"
	"log"
	"strings"
)

const (
	VerticalPipe   = "|"
	HorizontalPipe = "-"
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
	var prevValue Coordinates

	LoopItems = append(LoopItems, start.value)

	for !loopIsCompleted {
		//When whole loop is completed
		if len(LoopItems) != 1 && location.value == Start {
			LoopItems = append(LoopItems, location.value)
			loopIsCompleted = true

			break
		}

		log.Println("New has been started: ", location)
		//Top check
		if location.y-1 > 0 {
			newY := location.y - 1
			topValue := pipes[newY][location.x]

			log.Println(topValue)

			if topValue == VerticalPipe && topValue != prevValue.value {
				prevValue = location
				location = Coordinates{
					y:     newY,
					x:     location.x,
					value: topValue,
				}

				LoopItems = append(LoopItems, topValue)

				continue
			}

			if topValue == SEBend && topValue != prevValue.value {
				prevValue = location
				location = Coordinates{
					y:     newY,
					x:     location.x - 1,
					value: topValue,
				}

				LoopItems = append(LoopItems, topValue)

				continue
			}

			if topValue == SWBend && topValue != prevValue.value {
				prevValue = location
				location = Coordinates{
					y:     newY,
					x:     location.x + 1,
					value: topValue,
				}

				LoopItems = append(LoopItems, topValue)

				continue
			}
		}

		//Bottom check
		if location.y+1 < len(pipes) {
			newY := location.y + 1
			bottomValue := pipes[newY][location.x]

			if bottomValue == VerticalPipe && bottomValue != prevValue.value {
				prevValue = location
				location = Coordinates{
					y:     newY,
					x:     location.x,
					value: bottomValue,
				}

				LoopItems = append(LoopItems, bottomValue)
				continue
			}

			if bottomValue == NWBend && bottomValue != prevValue.value {
				prevValue = location

				location = Coordinates{
					y:     newY,
					x:     location.x,
					value: bottomValue,
				}

				LoopItems = append(LoopItems, bottomValue)
				continue
			}

			if bottomValue == NEBend && bottomValue != prevValue.value {
				prevValue = location

				location = Coordinates{
					y:     newY,
					x:     location.x,
					value: bottomValue,
				}

				LoopItems = append(LoopItems, bottomValue)
				continue
			}
		}

		//Left check
		if location.x-1 >= 0 {
			newX := location.x - 1
			leftValue := pipes[location.y][newX]

			if leftValue == NEBend && leftValue != prevValue.value {
				prevValue = location

				location = Coordinates{
					y:     location.y,
					x:     newX,
					value: leftValue,
				}

				LoopItems = append(LoopItems, leftValue)
				continue
			}

			if leftValue == HorizontalPipe && leftValue != prevValue.value {
				prevValue = location

				location = Coordinates{
					y:     location.y,
					x:     newX,
					value: leftValue,
				}

				LoopItems = append(LoopItems, leftValue)
				continue
			}

			if leftValue == SEBend && leftValue != prevValue.value {
				prevValue = location

				location = Coordinates{
					y:     location.y,
					x:     newX,
					value: leftValue,
				}

				LoopItems = append(LoopItems, leftValue)
				continue
			}
		}

		//Right check
		if location.x+1 < len(pipes[location.y]) {
			newX := location.x + 1
			rightValue := pipes[location.y][newX]

			if rightValue == NWBend && rightValue != prevValue.value {
				prevValue = location

				location = Coordinates{
					y:     location.y,
					x:     newX,
					value: rightValue,
				}

				LoopItems = append(LoopItems, rightValue)
				continue
			}

			if rightValue == HorizontalPipe && rightValue != prevValue.value {
				prevValue = location

				location = Coordinates{
					y:     location.y,
					x:     newX,
					value: rightValue,
				}

				LoopItems = append(LoopItems, rightValue)
				continue
			}

			if rightValue == SWBend && rightValue != prevValue.value {
				prevValue = location

				location = Coordinates{
					y:     location.y,
					x:     newX,
					value: rightValue,
				}

				LoopItems = append(LoopItems, rightValue)
				continue
			}
		}

		loopIsCompleted = true
	}

	log.Println(LoopItems)

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
