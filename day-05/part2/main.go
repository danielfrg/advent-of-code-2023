package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"math"
	"regexp"
	"strings"
)

func main() {
	input := utils.ReadFile("input.txt")

	seedSoil := parseMapping(input, "seed", "soil")
	soilFertilizer := parseMapping(input, "soil", "fertilizer")
	fertilizerWater := parseMapping(input, "fertilizer", "water")
	waterLight := parseMapping(input, "water", "light")
	lightTemperature := parseMapping(input, "light", "temperature")
	temperatureHumidity := parseMapping(input, "temperature", "humidity")
	humidityLocation := parseMapping(input, "humidity", "location")

	seeds := parseInput(input)

	lowestLocation := int(math.Inf(1))

	for i := 0; i < len(seeds); i = i + 2 {
		start := seeds[i]
		len := seeds[i+1]

		for seed := start; seed < start+len; seed++ {
			soil := findNext(seedSoil, seed)
			fertilizer := findNext(soilFertilizer, soil)
			water := findNext(fertilizerWater, fertilizer)
			light := findNext(waterLight, water)
			temperature := findNext(lightTemperature, light)
			humidity := findNext(temperatureHumidity, temperature)
			location := findNext(humidityLocation, humidity)

			if location < int(lowestLocation) {
				lowestLocation = location
			}
		}
	}

	fmt.Println("----------")
	fmt.Println(lowestLocation)
}

func findNext(mapping [][]int, target int) int {

	for i := 0; i < len(mapping); i++ {
		currentMap := mapping[i]

		destination, source, length := currentMap[0], currentMap[1], currentMap[2]
		if target >= source && target < source+length {
			return destination + (target - source)
		}
	}

	return target
}

func parseMapping(input string, source string, destination string) [][]int {
	expr := fmt.Sprintf(`%s-to-%s map:\n([\d \n]*)`, source, destination)
	re := regexp.MustCompile(expr)
	match := re.FindStringSubmatch(input)

	mappings := utils.Split(match[1], '\n')

	var ret [][]int

	for i := 0; i < len(mappings); i++ {
		currentMap := utils.Split(mappings[i], ' ')
		ret = append(ret, utils.ToIntSlice(currentMap))
	}

	return ret
}

func parseInput(input string) []int {
	re := regexp.MustCompile(`^seeds: ([\d ]*)`)
	match := re.FindStringSubmatch(input)

	textNums := strings.Split(match[1], " ")
	return utils.ToIntSlice(textNums)
}
