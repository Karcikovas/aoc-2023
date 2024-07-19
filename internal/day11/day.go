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
	var galaxies int = 0
	q := utils.NewQueue()
	content := utils.ReadInput("./assets/day11input.txt")
	rows := strings.Split(string(content), "\n")

	for _, row := range rows {
		galaxiesCount := checkSubstrings(row, string(Galaxies))

		if galaxiesCount > 0 {

			q.Push(row)

			galaxies = galaxies + 1

			log.Println(galaxiesCount)

			continue
		}

		q.Push(row)
		q.Push(row)
	}

	for q.Len() != 0 {
		n := q.Pop()

		log.Println(n)
	}

	pairAmount := utils.Combinations(9, 2)

	log.Println(pairAmount, galaxies)
}

func checkSubstrings(str string, value string) int {
	matches := 0

	for _, char := range str {
		if strings.Contains(string(char), value) {
			matches += 1
		}
	}

	return matches
}
