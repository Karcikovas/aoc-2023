package day5

import (
	"aoc2023/internal/utils"
	"regexp"
	"strings"
)

func CalculatePart1() int {
	var location []int
	almanacMap, seeds := InitialValues("./assets/day5input.txt")

	for _, seed := range seeds {
		location = append(location, GetLocationNumber(almanacMap, seed))
	}

	return utils.GetLowestNumberInCollection(location)
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

func InitialValues(filePath string) (*AlmanacMap, []int) {
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

	return almanacMap, seeds
}
func GetMapByTitle(almanac *AlmanacMap, title string) *DestinationToSourceMap {
	for i := range almanac.Maps {
		if almanac.Maps[i].Title == title {
			return &almanac.Maps[i]
		}
	}
	return nil
}
func GetLocationNumber(almanacMap *AlmanacMap, seedNumber int) int {
	var destinationToSourceNumber = seedNumber
	var maps []string

	for _, m := range almanacMap.Maps {
		maps = append(maps, m.Title)
	}

	for _, v := range maps {
		destinationToSourceMap := GetMapByTitle(almanacMap, v)

		for _, i := range destinationToSourceMap.Items {
			sum := i.Source + i.Length - 1

			if i.Source < seedNumber && sum > seedNumber {
				destinationSum := i.Destination + i.Length - 1
				diff := sum - seedNumber

				destinationToSourceNumber = destinationSum - diff

				break
			}
		}

	}

	return destinationToSourceNumber
}
