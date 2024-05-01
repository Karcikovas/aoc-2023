package day7

import (
	"aoc2023/internal/utils"
	"sort"
	"strings"
)

const (
	Ace   = 14
	King  = 13
	Queen = 12
	Jack  = 11
	T     = 10
)

type HandWithStrength struct {
	Cards    string
	Bid      int
	Strength int
}

type Hand struct {
	Cards string
	Bid   int
}

func CalculatePart1() int {
	hands := initialValues()
	handsWithStrength := getHandsWithStrength(hands)
	handsOrderedByStrength := sortHands(handsWithStrength)
	totalValue := calculateTotalWinnings(handsOrderedByStrength)

	return totalValue
}

func calculateTotalWinnings(hands []HandWithStrength) int {
	var totalValue int

	for i, hand := range hands {
		totalValue = totalValue + (i+1)*hand.Bid
	}

	return totalValue
}

func sortHands(hands []HandWithStrength) []HandWithStrength {

	sort.Slice(hands, func(i, j int) bool {
		a := hands[i]
		b := hands[j]

		if a.Strength != b.Strength {
			return a.Strength < b.Strength
		}

		for i := 0; i <= 5; i++ {
			if a.Cards[i] != b.Cards[i] {
				//TODO: Need to add card values comparison
				return true
			}
		}

		return false
	})

	return hands
}

func getHandsWithStrength(hands []Hand) []HandWithStrength {
	const maxAmountOfPairs = 2

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

		sortedItems := utils.SortMapByHighestValue(counts)
		firstValue := sortedItems[0].Value
		secondValue := sortedItems[1].Value

		hand := HandWithStrength{
			Cards:    hand.Cards,
			Bid:      hand.Bid,
			Strength: maxAmountOfPairs*firstValue + secondValue,
		}

		handsWithStrength = append(handsWithStrength, hand)
	}

	return handsWithStrength
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
