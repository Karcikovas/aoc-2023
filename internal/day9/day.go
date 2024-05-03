package day9

import (
	"aoc2023/internal/utils"
	"log"
	"strings"
)

type HistoryValuesMap = map[int]HistoryValues
type HistoryValues = []int

func CalculatePart1() int {
	historyValuesMap, totalValuesCount := initialValues()
	predictions := getPredictions(historyValuesMap, totalValuesCount)

	log.Println(predictions)

	return 0
}

var historyMap = make(HistoryValuesMap)

func getPredictions(m HistoryValuesMap, count int) []int {
	var predictions []int

	for i := 0; i < count; i++ {
		var history []int
		history = m[i]

		log.Print(history)
		historyMap[0] = history
		prediction(history, false, 0)
		log.Print(historyMap)
	}

	return predictions
}

func prediction(m []int, b bool, i int) {
	var n []int
	var lastIsZero bool

	if !b {
		for i, v := range m {
			if i+1 < len(m) {
				diff := m[i+1] - v
				lastIsZero = diff == 0

				n = append(n, diff)
			} else {
				break
			}
		}

		historyMap[i+1] = n
		prediction(n, lastIsZero, i+1)
	}
}

func initialValues() (HistoryValuesMap, int) {
	var totalValuesCount int
	var historyValuesMap = make(HistoryValuesMap)

	content := utils.ReadInput("./assets/day9input.txt")
	rows := strings.Split(string(content), "\n")

	for i, row := range rows {
		var values []int
		valuesStrings := strings.Split(row, " ")

		for _, value := range valuesStrings {
			values = append(values, utils.StringToInt(value))
		}

		historyValuesMap[i] = values
		totalValuesCount += 1
	}

	return historyValuesMap, totalValuesCount
}
