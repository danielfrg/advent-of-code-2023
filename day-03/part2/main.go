package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	utils "advent-of-code-2023/utils"
)

func main() {
	lines := utils.ReadLines("input.txt")

	numbersPerLine := make([][]string, 0)
	numbersIndicesPerLine := make([][][]int, 0)
	asteriscIndicesPerLine := make([][]int, 0)

	for _, line := range lines {
		reNumbers := regexp.MustCompile("[0-9]+")
		numbers := reNumbers.FindAllString(line, -1)
		indices := reNumbers.FindAllStringIndex(line, -1)

		reSymbols := regexp.MustCompile("[*]")
		asteriscSymbols := reSymbols.FindAllStringIndex(line, -1)

		cleanAsteriscSymbols := make([]int, 0)

		for i := 0; i < len(asteriscSymbols); i++ {
			cleanAsteriscSymbols = append(cleanAsteriscSymbols, asteriscSymbols[i][0])
		}

		numbersPerLine = append(numbersPerLine, numbers)
		numbersIndicesPerLine = append(numbersIndicesPerLine, indices)
		asteriscIndicesPerLine = append(asteriscIndicesPerLine, cleanAsteriscSymbols)

	}
	// fmt.Println(numbersPerLine)
	// fmt.Println(asteriscIndicesPerLine)

	// Iterate the asteric indices and see if they are next to numbers

	gearRatios := make([]int, 0)

	for i := 0; i < len(asteriscIndicesPerLine); i++ {
		asterisctIndicesOnThisLine := asteriscIndicesPerLine[i]

		for g := 0; g < len(asterisctIndicesOnThisLine); g++ {
			asteriscIndex := asterisctIndicesOnThisLine[g]

			numbersAdj := make([]int, 0)

			// Compare with the surounding lines
			comp := []int{-1, 0, 1}
			for j := 0; j < len(comp); j++ {
				compareIdx := i + comp[j]
				if compareIdx < 0 || compareIdx >= len(asteriscIndicesPerLine) {
					continue
				}

				partNumberLine := numbersPerLine[compareIdx]
				partNumberIdxLine := numbersIndicesPerLine[compareIdx]
				for k := 0; k < len(partNumberIdxLine); k++ {
					partNumberIndeces := partNumberIdxLine[k]
					if isNextToNumber(asteriscIndex, partNumberIndeces) {
						num, err := strconv.Atoi(partNumberLine[k])
						if err != nil {
							log.Panic("Error parsing number: ", partNumberLine[k])
							os.Exit(1)
						}
						numbersAdj = append(numbersAdj, num)
					}
				}
			}

			// fmt.Println(numbersAdj)
			if len(numbersAdj) == 2 {
				gearRatios = append(gearRatios, numbersAdj[0]*numbersAdj[1])
			}
		}
	}

	fmt.Println("----------")
	fmt.Println(utils.Sum(gearRatios))
}

func isNextToNumber(asteriscIndex int, partNumberIndices []int) bool {
	low, high := partNumberIndices[0], partNumberIndices[1]
	if asteriscIndex >= low && asteriscIndex <= high {
		return true
	}
	if asteriscIndex == low-1 {
		return true
	}
	if asteriscIndex == high {
		return true
	}
	return false
}
