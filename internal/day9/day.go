package day9

import (
	"aoc2023/internal/utils"
	"strings"
)

type HistoryValuesMap = map[int]HistoryValues
type HistoryValues = []int

func CalculatePart1() int {
	var totalSum int
	historyValuesMap, totalValuesCount := initialValues()
	predictions := getPredictions(historyValuesMap, totalValuesCount)

	for _, v := range predictions {

		totalSum = totalSum + v
	}

	return totalSum
}

var historyMap = make(HistoryValuesMap)

func getPredictions(m HistoryValuesMap, count int) []int {
	var predictions []int

	for i := 0; i < count; i++ {
		var history []int
		history = m[i]

		historyMap[0] = history
		prediction(history, false, 0)

		count := 0
		for _ = range historyMap {
			count++
		}

		for c := count; c > 0; c-- {
			idx := c - 1
			collection := historyMap[idx]
			l := len(collection)
			lastValue := collection[l-1]

			if c == count {
				historyMap[idx] = append(historyMap[idx], 0)
			} else {
				collectionBefore := historyMap[c]
				valueBefore := collectionBefore[len(collectionBefore)-1]

				historyMap[idx] = append(historyMap[idx], lastValue+valueBefore)
			}
		}

		lCol := historyMap[0]

		predictions = append(predictions, lCol[len(lCol)-1])
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
