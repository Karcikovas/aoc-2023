package day7

import (
	"aoc2023/internal/utils"
	"log"
	"strings"
)

func CalculatePart1() int {
	input := initialValues()

	log.Println(input)

	return 0
}

type Hand struct {
	Cards string
	Bid   int
}

func initialValues() []Hand {
	var hands []Hand

	content := utils.ReadInput("./assets/day7input.txt")
	rows := strings.Split(string(content), "\n")

	for _, row := range rows {
		slice := strings.Split(row, " ")
		cards := slice[0]
		bid := slice[1]

		if len(slice) >= 2 {
			hand := Hand{
				Cards: cards,
				Bid:   utils.StringToInt(bid),
			}

			hands = append(hands, hand)
		}
	}

	return hands
}
