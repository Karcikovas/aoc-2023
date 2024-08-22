package day12

import (
	"aoc2023/internal/utils"
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

func isValid(springsString string, arrangement []int) bool {
	parts := strings.Split(springsString, ".")

	var groups []int
	for _, part := range parts {
		if part != "" {
			groups = append(groups, len(part))
		}
	}

	return utils.ArraysEqual(groups, arrangement)
}

func travers(springsString string, arrangement []int) {
	nextUnknown := utils.IndexOf(springsString, "?")

	if nextUnknown == -1 {
		if isValid(springsString, arrangement) {
			totalSum += 1
		}
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
	numbers := strings.Split(arrangements, ",")
	var values []int

	for _, nmb := range numbers {
		values = append(values, utils.StringToInt(nmb))
	}

	return values
}
