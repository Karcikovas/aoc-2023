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
	x     int
	y     int
	value string
}

func CalculatePart1() int {
	nodes, start := initialValues()

	getPipeLoopItems(nodes, start)

	return 0
}

func isParent(currX int, currY int, parentX int, parentY int) bool {
	return currX == parentX && currY == parentY
}

func isPossibleNeighbor(pipes [][]Node, x int, y int, fromX int, fromY int, combinator []string) bool {

	if x >= 0 && y >= 0 {
		log.Println(y, x, pipes[y])
		targetNode := pipes[y][x]

		return slices.Contains(combinator, targetNode.value) && !isParent(x, y, fromX, fromY)
	} else {
		return false
	}
}

func getPipeLoopItems(pipes [][]Node, start Node) {
	var stack utils.Container[Node] = &utils.Stack[Node]{}
	stack.Push(start)

	for stack.Len() != 0 {
		n := stack.Pop()

		if isPossibleNeighbor(pipes, n.x, n.y-1, start.x, start.y, UpCombinations) {
			log.Println("UpCombination")
		}

		if isPossibleNeighbor(pipes, n.x, n.y+1, start.x, start.y, DownCombinations) {
			log.Println("DownCombination")
		}

		if isPossibleNeighbor(pipes, n.x-1, n.y, start.x, start.y, LeftCombinations) {
			log.Println("LeftCombination")
		}

		if isPossibleNeighbor(pipes, n.x+1, n.y, start.x, start.y, RightCombinations) {
			log.Println("RightCombination")
		}
	}
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
