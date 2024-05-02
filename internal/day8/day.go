package day8

import (
	"aoc2023/internal/utils"
	"log"
	"strings"
)

type NodesMap = map[string]Node
type NetworkInstructions = []string

type Node struct {
	Key   string
	Left  string
	Right string
}

func CalculatePart1() int {
	network, networkMap := initialValues()

	log.Println(network, networkMap)

	return 0
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
			log.Println(key, stringValues)

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
