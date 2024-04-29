package utils

import (
	"sort"
)

func BinarySearch(successValue int, options []int) bool {
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

type Item struct {
	Key   string
	Value int
}

func SortByHighestValue(m map[string]int) []Item {
	var items []Item

	for i, v := range m {
		items = append(items, Item{Key: i, Value: v})
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].Value > items[j].Value
	})

	return items
}
