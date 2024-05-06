package day9

import (
	"aoc2023/internal/utils"
	"log"
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

func getPredictions(m HistoryValuesMap, count int) []int {
	var predictions []int

	for i := 0; i < count; i++ {
		var sequence [][]int

		sequence = append(sequence, m[i])

		allZeroes := false

		for !allZeroes {
			seq := sequence[len(sequence)-1]
			newSeq := generateSequence(seq)

			sequence = append(sequence, newSeq)

			if allZero(newSeq) {
				allZeroes = true
			}
		}

		count := 0
		for _ = range sequence {
			count++
		}

		sum := 0
		prev := 0
		for c := count - 2; c >= 0; c-- {
			num := prev + sequence[c][len(sequence[c])-2]
			prev = num
			sum += num
		}

		log.Println(sequence)
		predictions = append(predictions, sum)
	}

	return predictions
}

func generateSequence(numbers []int) []int {
	var seq []int
	for i := 1; i < len(numbers); i++ {
		seq = append(seq, numbers[i]-numbers[i-1])
	}

	return seq
}

func allZero(s []int) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
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
