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
	var galaxies int
	q := utils.NewQueue()
	content := utils.ReadInput("./assets/day11input.txt")
	rows := strings.Split(content, "\n")

	for _, row := range rows {
		galaxiesCount := utils.CheckSubstrings(row, string(Galaxies))

		if galaxiesCount > 0 {
			q.Push(row)

			galaxies = galaxies + galaxiesCount

			continue
		}

		q.Push(row)
		q.Push(row)
	}

	for q.Len() != 0 {
		n := q.Pop()

		log.Println(n)
	}

	pairAmount := utils.Combinations(galaxies, 2)

	log.Println(pairAmount, galaxies)
}
