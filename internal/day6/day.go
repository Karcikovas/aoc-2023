package day6

import (
	"aoc2023/internal/utils"
	"regexp"
	"strings"
)

func CalculatePart1() int {
	input := initialValues()
	margin := 1

	for _, race := range input.Races {
		n := calculateNumberOfWinnings(race)

		margin = margin * n
	}

	return margin
}

type Race struct {
	RaceNumber int
	Duration   int
	Record     int
}

type Races struct {
	Races []Race
}

func initialValues() Races {
	var times []string
	var distances []string
	races := Races{}

	content := utils.ReadInput("./assets/day6input.txt")
	rows := strings.Split(string(content), "\n")

	for _, row := range rows {
		sliceIndex := strings.Index(row, ":")
		valuesInString := row[sliceIndex+1:]

		fmtTime := "\\bTime\\b"
		timePattern := regexp.MustCompile(fmtTime)
		isTime := bool(timePattern.MatchString(row))
		re := regexp.MustCompile(`\d+`)

		if isTime {
			times = re.FindAllString(valuesInString, -1)
		} else {
			distances = re.FindAllString(valuesInString, -1)
		}
	}

	for idx, time := range times {
		race := Race{
			RaceNumber: idx + 1,
			Duration:   utils.StringToInt(time),
			Record:     utils.StringToInt(distances[idx]),
		}

		races.Races = append(races.Races, race)
	}

	return races
}

func calculateNumberOfWinnings(race Race) int {
	var numberOfWays int

	for s := 0; s <= race.Duration; s++ {
		timeLeft := race.Duration - s
		distanceCovered := timeLeft * s

		if distanceCovered > race.Record {

			numberOfWays++
		}
	}

	return numberOfWays
}
