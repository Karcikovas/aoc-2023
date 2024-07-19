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
	var grid []string

	content := utils.ReadInput("./assets/day11input.txt")
	rows := strings.Split(content, "\n")

	for _, row := range rows {
		galaxiesCount := utils.CheckSubstrings(row, string(Galaxies))
		grid = append(grid, row)

		if galaxiesCount > 0 {

			galaxies = galaxies + galaxiesCount

			continue
		} else {
			grid = append(grid, row)
		}

	}

	for row := 0; row < len(grid); row++ {
		l := len(grid[row])

		for col := 0; col < l; col++ {
			log.Println(grid[row], l)
			log.Println(string(grid[row][col]))
		}

		//for col := 0; col < len(grid[row]); col++ {
		//	fmt.Print(string(grid[row][col]), "--------")
		//}
	}

	pairAmount := utils.Combinations(galaxies, 2)

	log.Println(pairAmount, galaxies)
}
