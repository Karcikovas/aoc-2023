package day11

import (
	"aoc2023/internal/utils"
	"fmt"
	"gonum.org/v1/gonum/stat/combin"
	"log"
	"strings"
)

const (
	EmptySpace = '.'
	Galaxies   = '#'
)

type Space struct {
	name     string
	x        int
	y        int
	isGalaxy bool
	value    string
}

func CalculatePart1() int {
	n := 9
	k := 2
	list := combin.Combinations(n, k)
	fmt.Println(len(list))

	initialValues()
	return 0
}

func initialValues() {
	q := utils.NewQueue()

	//var spaceMap [][]Space
	content := utils.ReadInput("./assets/day11input.txt")
	rows := strings.Split(string(content), "\n")

	for _, row := range rows {
		startIndex := strings.Index(row, "#")

		if startIndex < 0 {
			q.Push(row)
		}
		q.Push(row)
	}

	for q.Len() != 0 {
		n := q.Pop()

		log.Println(n)
	}
}
