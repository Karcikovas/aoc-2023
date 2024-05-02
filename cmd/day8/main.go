package main

import (
	"aoc2023/internal/day8"
	"log"
)

func main() {
	totalPointSum := day8.CalculatePart1()

	log.Print("Day 8 total steps required to reach ZZZ: ", totalPointSum)
}
