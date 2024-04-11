package day3

import (
	"aoc2023/internal/utils"
	"regexp"
	"strconv"
	"strings"
)

func CalculatePart1() int {
	var totalSum int
	content := utils.ReadInput("./assets/day3input.txt")
	rows := strings.Split(string(content), "\n")

	var grid [140][]string

	for y, line := range rows {
		characters := strings.Split(line, "")
		grid[y] = characters
	}

	for y, row := range grid {
		number := ""

		for x, character := range row {
			isNumber, _ := regexp.MatchString(`[0-9]`, character)

			if isNumber {
				number += grid[y][x]

			} else if number != "" || x == len(row) {

				totalSum += validateNumber(number, x, y, grid)
				number = ""
			} else {
				number = ""
			}
		}

		totalSum += validateNumber(number, len(row), y, grid)
	}

	return totalSum
}

func validateNumber(number string, x int, y int, grid [140][]string) int {

	end := x - 1
	start := x - len(number)
	middle := len(number)/2 + start

	coordinates := [][]int{
		{y, start - 1},
		{y, end + 1},

		{y + 1, start - 1},
		{y + 1, start},
		{y + 1, middle},
		{y + 1, end},
		{y + 1, end + 1},

		{y - 1, start - 1},
		{y - 1, start},
		{y - 1, middle},
		{y - 1, end},
		{y - 1, end + 1},
	}

	for _, coordinate := range coordinates {
		if coordinate[0] >= 0 && coordinate[0] < len(grid) && coordinate[1] >= 0 && coordinate[1] < len(grid[0]) {
			isNumber, _ := regexp.MatchString(`[0-9]`, grid[coordinate[0]][coordinate[1]])

			if grid[coordinate[0]][coordinate[1]] != "." && !isNumber {
				convertedNumber, _ := strconv.Atoi(number)

				return convertedNumber
			}
		}
	}

	return 0
}
