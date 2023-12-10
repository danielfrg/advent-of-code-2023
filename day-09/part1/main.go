package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// lines := utils.ReadLines("input-test.txt")
	lines := utils.ReadLines("input.txt")
	sequences := parseLines(lines)
	// fmt.Println(sequences)

	values := make([]int, len(sequences))
	for i, seq := range sequences {
		value := extrapolate(seq)
		values[i] = value
		// break
	}

	fmt.Println("----------")
	fmt.Println(utils.Sum(values))
}

func extrapolate(seq []int) int {
	steps := make([][]int, 0)
	// fmt.Println(seq)
	steps = append(steps, seq)
	for {
		seq = reduce(seq)
		// fmt.Println(seq)
		steps = append(steps, seq)
		if allZero(seq) {
			break
		}
	}

	// Extrapolate part
	adding := 0
	for i := len(steps) - 1; i >= 0; i-- {
		seq := steps[i]
		adding = seq[len(seq)-1] + adding
		// fmt.Println(seq, adding)
	}

	return adding
}

func allZero(arr []int) bool {
	for _, val := range arr {
		if val != 0 {
			return false
		}
	}

	return true
}

func reduce(seq []int) []int {
	ret := make([]int, len(seq)-1)

	for i := 1; i < len(seq); i++ {
		ret[i-1] = seq[i] - seq[i-1]
	}

	return ret
}

func parseLines(lines []string) [][]int {
	sequences := make([][]int, len(lines))

	for i, line := range lines {
		sequences[i] = parseLine(line)
	}

	return sequences
}

func parseLine(line string) []int {
	parts := strings.Split(line, " ")

	ret := make([]int, len(parts))
	for i := 0; i < len(parts); i++ {
		val, _ := strconv.Atoi(parts[i])
		ret[i] = val
	}

	return ret
}
