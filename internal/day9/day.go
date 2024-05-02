package day9

import (
	"aoc2023/internal/utils"
	"strings"
)

type NodesMap = map[string]Node
type NetworkInstructions = []string

type Node struct {
	Key   string
	Left  string
	Right string
}

var stepIndex int = 0

func CalculatePart1() int {
	network, networkMap := initialValues()
	startKey := "AAA"
	destinationKey := "ZZZ"

	getStepsCount(startKey, network, networkMap, destinationKey)

	return stepIndex
}

func getStepsCount(key string, network NetworkInstructions, networkMap NodesMap, destination string) {
	if destination != key {
		if stepIndex >= len(network) {
			network = append(network, network...)
		}

		direction := string(network[stepIndex])
		stepIndex += 1

		if direction == "L" {
			getStepsCount(networkMap[key].Left, network, networkMap, destination)
		} else {
			getStepsCount(networkMap[key].Right, network, networkMap, destination)
		}
	}
}

func initialValues() (NetworkInstructions, NodesMap) {
	var network NetworkInstructions
	var networkMap = make(NodesMap)

	content := utils.ReadInput("./assets/day8input.txt")
	rows := strings.Split(string(content), "\n")

	for i, row := range rows {
		if i == 0 {
			for _, char := range row {
				network = append(network, string(char))
			}
		}

		if len(row) > 0 && i != 0 {
			slices := strings.Split(row, " = ")
			key := slices[0]
			stringValues := slices[1]
			stringValues = strings.TrimLeft(stringValues, "(")
			stringValues = strings.TrimRight(stringValues, ")")
			values := strings.Split(stringValues, ", ")

			node := Node{
				Key:   key,
				Left:  values[0],
				Right: values[1],
			}
			networkMap[key] = node
		}
	}

	return network, networkMap
}
