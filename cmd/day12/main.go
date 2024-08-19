package main

import (
	"aoc2023/internal/day12"
	"log"
)

func main() {
	totalSumOfArrangements := day12.CalculatePart1()

	log.Print("Day 12 result Part1 sum of steps between galaxies: ", totalSumOfArrangements)
}
