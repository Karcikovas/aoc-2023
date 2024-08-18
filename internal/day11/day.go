package day11

import (
	"aoc2023/internal/utils"
	"strings"
)

const (
	EmptySpace = '.'
	Galaxies   = '#'
)

type SpaceNode struct {
	x        int
	y        int
	isGalaxy bool
	value    string
}

func CalculatePart1() int {
	initialValues()
	return 0
}

func initialValues() {
	var galaxies int
	var spaceGrid []string
	expandedCols := 0

	content := utils.ReadInput("./assets/day11input.txt")
	rows := strings.Split(content, "\n")

	for _, row := range rows {
		galaxiesCount := utils.CheckSubstrings(row, string(Galaxies))
		spaceGrid = append(spaceGrid, row)

		if galaxiesCount > 0 {

			galaxies = galaxies + galaxiesCount

			continue
		} else {
			spaceGrid = append(spaceGrid, row)
		}
	}

	for col := 0; col < len(rows[0]); col++ {
		empty := true

		for x := 0; x < len(rows[0]); x++ {
			if rows[col][x] == Galaxies {
				empty = false
				continue
			}
		}

		if empty {
			for rowToUpdate := 0; rowToUpdate < len(spaceGrid); rowToUpdate++ {
				updateStr := spaceGrid[rowToUpdate]

				part1 := updateStr[:col+expandedCols]
				part2 := updateStr[col+expandedCols:]

				spaceGrid[rowToUpdate] = part1 + string(EmptySpace) + part2
			}

			expandedCols++
		}
	}
}
