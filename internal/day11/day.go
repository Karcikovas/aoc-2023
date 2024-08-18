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

type Star struct {
	x int
	y int
}

func CalculatePart1() int {
	spaceGrid, galaxies := initialValues()
	stars := createStartMap(spaceGrid)

	log.Println(galaxies)

	log.Println(stars)
	return 0
}

func createStartMap(spaceGrid []string) []Star {
	var stars []Star
	for x := range spaceGrid {
		for y := 0; y < len(spaceGrid[x]); y++ {
			if spaceGrid[x][y] == Galaxies {
				star := Star{
					x: x,
					y: y,
				}

				stars = append(stars, star)
			}
		}
	}

	return stars
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

		if galaxiesCount == 0 {
			spaceGrid = append(spaceGrid, row)
		}
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
