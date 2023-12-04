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
	symbolIndicesPerLine := make([][]int, 0)

	for _, line := range lines {
		reNumbers := regexp.MustCompile("[0-9]+")
		numbers := reNumbers.FindAllString(line, -1)
		indices := reNumbers.FindAllStringIndex(line, -1)

		reSymbols := regexp.MustCompile("[^0-9.]")
		indicesSymbols := reSymbols.FindAllStringIndex(line, -1)

		cleanIndicesSymbols := make([]int, 0)

		for i := 0; i < len(indicesSymbols); i++ {
			cleanIndicesSymbols = append(cleanIndicesSymbols, indicesSymbols[i][0])
		}

		// fmt.Println(line, cleanIndicesSymbols)

		numbersPerLine = append(numbersPerLine, numbers)
		numbersIndicesPerLine = append(numbersIndicesPerLine, indices)
		symbolIndicesPerLine = append(symbolIndicesPerLine, cleanIndicesSymbols)
	}

	// Iterate the number indices and search on the symbol indices

	parts := make([]int, 0)
	for i := 0; i < len(numbersPerLine); i++ {
		lineNumbers := numbersPerLine[i]
		lineNumberIndices := numbersIndicesPerLine[i]

		partsAddedOnLine := make([]int, 0)
		for j := 0; j < len(lineNumbers); j++ {
			partNumber := lineNumbers[j]
			partNumberIndeces := lineNumberIndices[j]

			comp := [3]int{-1, 0, 1}
			for k := 0; k < len(comp); k++ {
				compareLineIndex := i + comp[k]
				if compareLineIndex < 0 || compareLineIndex >= len(symbolIndicesPerLine) {
					continue
				}

				symbolIndices := symbolIndicesPerLine[compareLineIndex]
				if isNextToSymbol(partNumberIndeces, symbolIndices) {
					num, err := strconv.Atoi(partNumber)
					if err != nil {
						log.Panic("Could not parse number")
						os.Exit(1)
					}
					partsAddedOnLine = append(partsAddedOnLine, num)
					break
				}
			}
		}

		parts = append(parts, partsAddedOnLine...)
		fmt.Println(i+1, partsAddedOnLine)
	}

	fmt.Println("----------")
	fmt.Println(utils.Sum(parts))
}

func isNextToSymbol(partNumberIndices []int, symbolIndices []int) bool {
	low, high := partNumberIndices[0], partNumberIndices[1]
	for _, symbolIndex := range symbolIndices {
		if symbolIndex >= low && symbolIndex <= high {
			return true
		}
		if symbolIndex == low-1 {
			return true
		}
		if symbolIndex == high {
			return true
		}
	}
	return false
}
