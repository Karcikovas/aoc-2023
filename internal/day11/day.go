package day11

import (
	"aoc2023/internal/utils"
	"log"
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
	spaceGrid, galaxies := initialValues()

	log.Println(spaceGrid, galaxies)
	return 0
}

func initialValues() ([]string, int) {
	var galaxies int
	var spaceGrid []string
	expandedCols := 0

	content := utils.ReadInput("./assets/day11input.txt")
	rows := strings.Split(content, "\n")

	for _, row := range rows {
		galaxiesCount := utils.CheckSubstrings(row, string(Galaxies))
		spaceGrid = append(spaceGrid, row)
		galaxies += galaxiesCount
	}

	for col := 0; col < len(rows[0]); col++ {
		if isColumnEmpty(rows, col) {
			insertEmptySpaceInColumn(&spaceGrid, col+expandedCols)
			expandedCols++
		}
	}

	return spaceGrid, galaxies
}

func isColumnEmpty(rows []string, col int) bool {
	for x := 0; x < len(rows[0]); x++ {
		if rows[col][x] == Galaxies {
			return false
		}
	}
	return true
}

func insertEmptySpaceInColumn(spaceGrid *[]string, col int) {
	for i, row := range *spaceGrid {
		(*spaceGrid)[i] = row[:col] + string(EmptySpace) + row[col:]
	}
}
