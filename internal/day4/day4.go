package day4

import (
	"aoc2023/internal/utils"
	"strconv"
	"strings"
)

func CalculatePart1() int {
	var totalCardsValue int
	var cards = getCardsValues("./assets/day4input.txt")

	for _, card := range cards {
		var score int
		cardNumbers := strings.Split(card, "|")
		winningNumbers := utils.GetIntArrayFromString(cardNumbers[0])
		numbersToCheck := utils.GetIntArrayFromString(cardNumbers[1])

		for _, winningNumber := range winningNumbers {
			isSuccessful := utils.BinarySearch(winningNumber, numbersToCheck)

			if isSuccessful && score == 0 {
				score = 1
			} else if isSuccessful && score > 0 {
				score *= 2
			}
		}

		totalCardsValue += score
	}

	return totalCardsValue
}

func getCardsValues(readFile string) [193]string {
	content := utils.ReadInput(readFile)
	rows := strings.Split(string(content), "\n")

	var cards [193]string

	for y, line := range rows {
		cardNumber := strconv.Itoa(y + 1)
		cardName := cardNumber + ":"
		numbers := strings.Split(line, cardName)

		cards[y] = numbers[1]
	}

	return cards
}
