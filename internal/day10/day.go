package day10

import (
	"aoc2023/internal/utils"
	"fmt"
	"log"
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

func getPipeLoopItems(pipes [][]Node, start Node) {
	var stack utils.Container[Node] = &utils.Stack[Node]{}

	stack.Push(start)
	stack.Len()

	for stack.Len() != 0 {
		n := stack.Pop()
		log.Println(n)
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
