package day4

import (
	"aoc2023/internal/utils"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func CalculatePart1() int {
	var totalCardsValue int
	var cards = getCardsValues("./../../assets/day4input.txt")

	for _, card := range cards {
		var score int
		cardNumbers := strings.Split(card, "|")
		winningNumbers := getIntArrayFromString(cardNumbers[0])
		numbersToCheck := getIntArrayFromString(cardNumbers[1])

		for _, winningNumber := range winningNumbers {
			isSuccessful := binarySearch(winningNumber, numbersToCheck)

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

		log.Print(numbers[1])

		cards[y] = numbers[1]
	}

	return cards
}

func getIntArrayFromString(stringToConvert string) []int {
	var number []int

	re := regexp.MustCompile(`(\d+)`)
	values := re.FindAllString(stringToConvert, -1)

	for _, value := range values {
		intVal, _ := utils.StringToInt(value)

		number = append(number, intVal)
	}

	return number
}

func binarySearch(successValue int, options []int) bool {
	low := 0
	high := len(options) - 1

	for low <= high {
		middleValueIndex := low + (high-low)/2
		middleValue := options[middleValueIndex]

		if successValue > middleValue {
			low = middleValueIndex + 1
		} else if successValue < middleValue {
			high = middleValueIndex - 1
		} else {
			return true
		}
	}

	if low == len(options) || low != successValue {
		return false
	}

	return true
}
