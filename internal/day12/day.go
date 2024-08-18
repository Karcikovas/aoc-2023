package day12

import (
	"aoc2023/internal/utils"
	"log"
	"math"
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
	spaceGrid := initialValues()
	stars := createStartMap(spaceGrid)

	log.Println(stars)
	sum := 0
	for i := 0; i < len(stars); i++ {
		for j := 0; j < len(stars); j++ {
			if i == j {
				continue
			}
			sum += manhattanDistance(stars[i], stars[j])
		}
	}

	return sum / 2
}

func manhattanDistance(a, b Star) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func createStartMap(spaceGrid []string) []Star {
	var stars []Star
	for col := range spaceGrid {
		for row := 0; row < len(spaceGrid[col]); row++ {
			if spaceGrid[col][row] == Galaxies {
				star := Star{
					x: row,
					y: col,
				}

				stars = append(stars, star)
			}
		}
	}

	return stars
}

func initialValues() []string {
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

	return spaceGrid
}

func isColumnEmpty(rows []string, col int) bool {
	for _, row := range rows {
		if col >= len(row) || row[col] != EmptySpace {
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
