package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	grid := readInput("input-test.txt")

	printGrid(grid)

	// path := make([][]int, len(grid))
	// printGrid(path)
}

func printGrid(in [][]int) {
	for _, l := range in {
		fmt.Println(l)
	}
}

func readInput(fname string) [][]int {
	file, _ := os.Open(fname)

	scanner := bufio.NewScanner(file)

	grid := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "")
		nums := make([]int, len(parts))
		for i, t := range parts {
			n, _ := strconv.Atoi(t)
			nums[i] = n
		}
		grid = append(grid, nums)
	}

	return grid
}
