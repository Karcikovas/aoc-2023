package day12

import (
	"aoc2023/internal/utils"
	"log"
	"strings"
)

type RecordRow struct {
	springsString string
	arrangement   []int
}

var totalSum int

func CalculatePart1() int {
	records := initialValues()

	for _, record := range records {
		travers(record.springsString, record.arrangement)
	}

	return totalSum
}

func travers(springsString string, arrangement []int) {
	nextUnknown := utils.IndexOf(springsString, "?")

	if nextUnknown == -1 {
		log.Println(springsString)

		totalSum += 1
	} else {
		travers(string(springsString[:nextUnknown])+"."+string(springsString[nextUnknown+1:]), arrangement)
		travers(string(springsString[:nextUnknown])+"#"+string(springsString[nextUnknown+1:]), arrangement)
	}
}

func initialValues() []RecordRow {
	var records []RecordRow
	content := utils.ReadInput("./assets/day12input.txt")
	rows := strings.Split(content, "\n")

	for _, row := range rows {
		splitRow := strings.Split(string(row), " ")
		arrangement := getArrangementArray(splitRow[1])

		record := RecordRow{
			springsString: splitRow[0],
			arrangement:   arrangement,
		}

		records = append(records, record)
	}

	return records
}

func getArrangementArray(arrangements string) []int {
	var values []int

	for _, arrangement := range arrangements {
		values = append(values, int(arrangement))
	}

	return values
}
