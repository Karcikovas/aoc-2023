package day2

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Cubes struct {
	red   int
	green int
	blue  int
}

func CalculatePart1() {
	file, err := os.Open("./assets/day2input.txt")
	totalGameNumber := 0
	totalGameNumberSum := 0

	cubes := Cubes{
		red:   12,
		green: 13,
		blue:  14,
	}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splits := strings.Split(scanner.Text(), ";")
		possibleSplits := 0
		cubePattern := regexp.MustCompile(` (\d+) (\w+)`)

		for _, split := range splits {
			matches := cubePattern.FindAllStringSubmatch(split, -1)
			items := make(map[string]int)

			for _, match := range matches {
				count, _ := strconv.Atoi(match[1])
				color := match[2]

				items[color] = count
			}

			if items["red"] <= cubes.red && items["blue"] <= cubes.blue && items["green"] <= cubes.green {
				possibleSplits++
			}
		}

		if possibleSplits == (len(splits)) {
			totalGameNumberSum += totalGameNumber + 1
		}

		totalGameNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Game Number:", totalGameNumber)
	log.Println("Tota games numbers sum:", totalGameNumberSum)
}
