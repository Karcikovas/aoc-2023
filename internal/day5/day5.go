package day5

import (
	"aoc2023/internal/utils"
	"log"
	"regexp"
	"strings"
)

func CalculatePart1() int {
	initialValues("./assets/day5input.txt")

	return 0
}

type mapItem struct {
	destination []int `json:"destination"`
	source      []int `json:"source"`
	length      []int `json:"length"`
}

type destinationToSourceMap struct {
	title string    `json:"title"`
	items []mapItem `json:"items"`
}

func initialValues(filePath string) {
	var seeds []int
	var almanacMap []destinationToSourceMap
	//var currentStruct mapItem

	content := utils.ReadInput(filePath)

	rows := strings.Split(string(content), "\n")

	for _, row := range rows {
		fmtMatchSeeds := "^seeds:( \\d+)+$"
		fmtMatchMap := "\\bmap\\b"
		seedsPattern := regexp.MustCompile(fmtMatchSeeds)
		mapsPattern := regexp.MustCompile(fmtMatchMap)
		isSeeds := bool(seedsPattern.MatchString(row))
		isNewMap := bool(mapsPattern.MatchString(row))

		if isSeeds {
			values := strings.Split(row, ":")
			seedsNumbers := utils.GetIntArrayFromString(values[1])

			seeds = seedsNumbers

			continue
		}

		if isNewMap {
			m := destinationToSourceMap{title: row}
			almanacMap = append(almanacMap, m)
		}

		//if isMap {
		//	almanacMap = append(almanacMap, currentStruct)
		//	currentStruct = mapItem{}
		//
		//	continue
		//} else {
		//	values := utils.GetIntArrayFromString(row)
		//
		//	if len(values) > 0 {
		//		currentStruct.destination = append(currentStruct.destination, values[0])
		//		currentStruct.source = append(currentStruct.source, values[1])
		//		currentStruct.length = append(currentStruct.length, values[2])
		//	}
		//}
	}
	log.Print(almanacMap, seeds)
}
