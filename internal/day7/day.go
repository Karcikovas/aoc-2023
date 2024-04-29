package day7

import (
	"aoc2023/internal/utils"
	"log"
	"reflect"
	"strings"
)

const (
	Ace   = 14
	King  = 13
	Queen = 12
	Jack  = 11
	T     = 10
)

func CalculatePart1() int {
	hands := initialValues()
	handsWithStrength := getHandsWithStrength(hands)

	log.Println(handsWithStrength)
	return 0
}

type HandWithStrength struct {
	Cards    string
	Bid      int
	Strength int
}

func getHandsWithStrength(hands []Hand) []HandWithStrength {
	var handsWithStrength []HandWithStrength
	for _, hand := range hands {
		counts := make(map[string]int)

		for _, card := range hand.Cards {
			c := string(card)

			if utils.NumToBool(counts[c]) {
				counts[c]++
			} else {
				counts[c] = 1
			}
		}

		utils.SortByHighestValue(counts)

		keys := reflect.ValueOf(counts).MapKeys()
		firstValue := counts[keys[0].String()]
		secondValue := counts[keys[1].String()]

		hand := HandWithStrength{
			Cards:    hand.Cards,
			Bid:      hand.Bid,
			Strength: 2*firstValue + secondValue,
		}

		handsWithStrength = append(handsWithStrength, hand)
	}

	return handsWithStrength
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
