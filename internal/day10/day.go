package day10

import (
	"aoc2023/internal/utils"
	"fmt"
	"log"
	"slices"
	"strings"
)

const (
	Start          = "S"
	VerticalPipe   = "|"
	HorizontalPipe = "-"
	NEBend         = "L"
	NWBend         = "J"
	SWBend         = "7"
	SEBend         = "F"
)

var UpCombinations = []string{Start, SWBend, SEBend, VerticalPipe}
var DownCombinations = []string{Start, NEBend, NWBend, VerticalPipe}
var LeftCombinations = []string{Start, HorizontalPipe, SEBend, NEBend}
var RightCombinations = []string{Start, HorizontalPipe, SWBend, NWBend}

type Node struct {
	name  string
	y     int
	x     int
	value string
}

func CalculatePart1() int {
	nodes, start := initialValues()

	log.Println(nodes, start)
	//getPipeLoopItems(pipes, start)

	return 0
}

func getPipeLoopItems(pipes [][]string, start Node) {
	var LoopItems []string
	var loopIsCompleted = false
	var location = start
	var prevValue Node

	LoopItems = append(LoopItems, start.value)
	log.Println(slices.Contains(UpCombinations, HorizontalPipe))

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
				location = Node{
					y:     newY,
					x:     location.x,
					value: topValue,
				}

				LoopItems = append(LoopItems, topValue)

				continue
			}

			if topValue == SEBend && topValue != prevValue.value {
				prevValue = location
				location = Node{
					y:     newY,
					x:     location.x - 1,
					value: topValue,
				}

				LoopItems = append(LoopItems, topValue)

				continue
			}

			if topValue == SWBend && topValue != prevValue.value {
				prevValue = location
				location = Node{
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
				location = Node{
					y:     newY,
					x:     location.x,
					value: bottomValue,
				}

				LoopItems = append(LoopItems, bottomValue)
				continue
			}

			if bottomValue == NWBend && bottomValue != prevValue.value {
				prevValue = location

				location = Node{
					y:     newY,
					x:     location.x,
					value: bottomValue,
				}

				LoopItems = append(LoopItems, bottomValue)
				continue
			}

			if bottomValue == NEBend && bottomValue != prevValue.value {
				prevValue = location

				location = Node{
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

				location = Node{
					y:     location.y,
					x:     newX,
					value: leftValue,
				}

				LoopItems = append(LoopItems, leftValue)
				continue
			}

			if leftValue == HorizontalPipe && leftValue != prevValue.value {
				prevValue = location

				location = Node{
					y:     location.y,
					x:     newX,
					value: leftValue,
				}

				LoopItems = append(LoopItems, leftValue)
				continue
			}

			if leftValue == SEBend && leftValue != prevValue.value {
				prevValue = location

				location = Node{
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

				location = Node{
					y:     location.y,
					x:     newX,
					value: rightValue,
				}

				LoopItems = append(LoopItems, rightValue)
				continue
			}

			if rightValue == HorizontalPipe && rightValue != prevValue.value {
				prevValue = location

				location = Node{
					y:     location.y,
					x:     newX,
					value: rightValue,
				}

				LoopItems = append(LoopItems, rightValue)
				continue
			}

			if rightValue == SWBend && rightValue != prevValue.value {
				prevValue = location

				location = Node{
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

func getNodeName(x int, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

func initialValues() ([][]Node, Node) {
	var nodes [][]Node
	var startNode Node
	content := utils.ReadInput("./assets/day10input.txt")
	rows := strings.Split(string(content), "\n")

	for y, row := range rows {
		startIndex := strings.Index(row, "S")

		nodes = append(nodes, make([]Node, 0))

		if startIndex >= 0 {
			startNode = Node{
				y:     y,
				x:     startIndex,
				value: Start,
				name:  getNodeName(startIndex, y),
			}
		} else {
			for x, value := range row {
				node := Node{
					y:     y,
					x:     x,
					value: string(value),
					name:  getNodeName(x, y),
				}
				nodes[y] = append(nodes[y], node)
			}
		}
	}

	return nodes, startNode
}
