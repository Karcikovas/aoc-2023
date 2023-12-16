package day5

import (
	"aoc2023/internal/utils"
	"encoding/json"
	"log"
	"regexp"
	"strings"
)

func CalculatePart1() int {
	initialValues("./../../assets/day5input.txt")

	return 0
}

type sourceToDestination struct {
	destination []int `json:"destination"`
	source      []int `json:"source"`
	length      []int `json:"length"`
}

func initialValues(filePath string) {
	var seeds []int
	var almanacMap []sourceToDestination
	var currentStruct sourceToDestination

	content := utils.ReadInput(filePath)

	rows := strings.Split(string(content), "\n")

	for _, row := range rows {
		fmtMatchSeeds := "^seeds:( \\d+)+$"
		fmtMatchMap := "\\bmap\\b"
		seedsPattern := regexp.MustCompile(fmtMatchSeeds)
		mapsPattern := regexp.MustCompile(fmtMatchMap)
		isSeeds := bool(seedsPattern.MatchString(row))
		isMap := bool(mapsPattern.MatchString(row))

		if isSeeds {
			values := strings.Split(row, ":")
			seedsNumbers := utils.GetIntArrayFromString(values[1])

			seeds = seedsNumbers

			continue
		}

		if isMap {
			almanacMap = append(almanacMap, currentStruct)
			currentStruct = sourceToDestination{}

			continue
		} else {
			values := utils.GetIntArrayFromString(row)

			if len(values) > 0 {
				currentStruct.destination = append(currentStruct.destination, values[0])
				currentStruct.source = append(currentStruct.source, values[1])
				currentStruct.length = append(currentStruct.length, values[2])
			}
		}
	}

	log.Print(seeds)
	fooMarshalled, _ := json.Marshal(almanacMap)

	log.Print(string(fooMarshalled))
}
