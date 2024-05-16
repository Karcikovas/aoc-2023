package day10

import (
	"aoc2023/internal/utils"
	"fmt"
	"slices"
	"strings"
)

const (
	Start          = "S"
	VerticalPipe   = "|"
	HorizontalPipe = "-"
	NEBend         = "L"
	SEBend         = "F"
	NWBend         = "J"
	SWBend         = "7"
)

var UpCombinations = []string{Start, SWBend, SEBend, VerticalPipe}
var DownCombinations = []string{Start, NEBend, NWBend, VerticalPipe}

var LeftCombinations = []string{Start, HorizontalPipe, SEBend, NEBend}
var RightCombinations = []string{Start, HorizontalPipe, SWBend, NWBend}

type Node struct {
	name      string
	x         int
	y         int
	value     string
	neighbors []Node
	visited   bool
}

func CalculatePart1() int {
	nodes, start := initialValues()

	getPipeLoopItems(nodes, start)

	return 0
}
func getOrCreateNode(pipes [][]Node, x int, y int, parentNode *Node) Node {
	node := pipes[y][x]

	count := 0
	for _, neighbor := range node.neighbors {
		if neighbor.name == parentNode.name {
			count++
		}
	}

	if count == 0 {
		node.neighbors = append(node.neighbors, *parentNode)
	}

	return node
}

func createNeighborNode(pipes [][]Node, x int, y int, parentNode *Node) (*Node, error) {
	node := getOrCreateNode(pipes, x, y, parentNode)
	var isPrevNeighbour bool

	for _, neighbor := range parentNode.neighbors {
		if neighbor.name == node.name {
			isPrevNeighbour = true
		}
	}

	if isPrevNeighbour {
		return nil, fmt.Errorf("node with same name as parent node already exists")
	}

	return &node, nil
}

func isParent(currX int, currY int, parentX int, parentY int) bool {
	return currX == parentX && currY == parentY
}

func isPossibleNeighbor(pipes [][]Node, x int, y int, fromX int, fromY int, combinator []string) bool {
	if y < 0 || y >= len(pipes) {
		return false
	}

	if x < 0 || x >= len(pipes[y]) {
		return false
	}

	targetNode := pipes[y][x]

	return slices.Contains(combinator, targetNode.value) && !isParent(x, y, fromX, fromY)
}

func getPipeLoopItems(pipes [][]Node, start Node) {
	var stack utils.Container[Node] = &utils.Stack[Node]{}
	stack.Push(start)

	for stack.Len() != 0 {
		n := stack.Pop()

		if isPossibleNeighbor(pipes, n.x, n.y-1, start.x, start.y, UpCombinations) {
			neighbor, err := createNeighborNode(pipes, n.x, n.y-1, &n)

			if err == nil {
				stack.Push(*neighbor)

				continue
			}
		}

		if isPossibleNeighbor(pipes, n.x, n.y+1, start.x, start.y, DownCombinations) {
			neighbor, err := createNeighborNode(pipes, n.x, n.y+1, &n)

			if err == nil {
				stack.Push(*neighbor)

				continue
			}
		}

		if isPossibleNeighbor(pipes, n.x-1, n.y, start.x, start.y, LeftCombinations) {
			neighbor, err := createNeighborNode(pipes, n.x-1, n.y, &n)

			if err == nil {
				stack.Push(*neighbor)

				continue
			}
		}

		if isPossibleNeighbor(pipes, n.x+1, n.y, start.x, start.y, RightCombinations) {
			neighbor, err := createNeighborNode(pipes, n.x+1, n.y, &n)

			if err == nil {
				stack.Push(*neighbor)

				continue
			}
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
