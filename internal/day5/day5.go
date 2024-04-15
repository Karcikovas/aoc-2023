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

type Item struct {
	Destination int
	Source      int
	Length      int
}

type DestinationToSourceMap struct {
	Title string
	Items []Item
}

type AlmanacMap struct {
	Maps []DestinationToSourceMap
}

func (almanacMap *AlmanacMap) AddMap(item DestinationToSourceMap) []DestinationToSourceMap {
	almanacMap.Maps = append(almanacMap.Maps, item)
	return almanacMap.Maps
}

func initialValues(filePath string) {
	var seeds []int
	var activeMapName string

	almanacMap := &AlmanacMap{}

	content := utils.ReadInput(filePath)

	rows := strings.Split(string(content), "\n")

	for _, row := range rows {
		fmtMatchSeeds := "^seeds:( \\d+)+$"
		fmtMatchMap := "\\bmap\\b"
		fmtMatchEmptyLine := "^\\s*$"
		seedsPattern := regexp.MustCompile(fmtMatchSeeds)
		mapsPattern := regexp.MustCompile(fmtMatchMap)
		emptyLinePattern := regexp.MustCompile(fmtMatchEmptyLine)
		isSeeds := bool(seedsPattern.MatchString(row))
		isNewMap := bool(mapsPattern.MatchString(row))
		isEmptyLine := bool(emptyLinePattern.MatchString(row))

		if isSeeds {
			values := strings.Split(row, ":")
			seedsNumbers := utils.GetIntArrayFromString(values[1])

			seeds = seedsNumbers

			continue
		}

		if isNewMap {
			m := DestinationToSourceMap{Title: row, Items: make([]Item, 0)}
			almanacMap.AddMap(m)
			activeMapName = row
		}

		if !isEmptyLine && !isNewMap {
			values := strings.Split(row, " ")

			i := Item{
				Destination: utils.StringToInt(values[0]),
				Source:      utils.StringToInt(values[1]),
				Length:      utils.StringToInt(values[2]),
			}

			m := GetMapByTitle(almanacMap, activeMapName)
			m.Items = append(m.Items, i)
		}
	}
	log.Print(GetMapByTitle(almanacMap, "light-to-temperature map:"), seeds)
}
func GetMapByTitle(almanac *AlmanacMap, title string) *DestinationToSourceMap {
	for i := range almanac.Maps {
		if almanac.Maps[i].Title == title {
			return &almanac.Maps[i]
		}
	}
	return nil
}
