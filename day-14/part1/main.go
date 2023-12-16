package main

import (
	"advent-of-code-2023/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readInput("input-test.txt")
	// printLines(input)

	tiltNorth(&input)
	utils.Assert(calc(&input), 136)

	input = readInput("input.txt")

	tiltNorth(&input)
	fmt.Println(calc(&input))
}

func calc(input *[][]string) int {
	val := 0

	for i := len(*input) - 1; i >= 0; i-- {
		val = val + (len(*input)-i)*count(&((*input)[i]), "O")
	}

	return val
}

func count(input *[]string, el string) int {
	c := 0
	for i := 0; i < len(*input); i++ {
		if (*input)[i] == el {
			c++
		}
	}
	return c
}

func tiltNorth(input *[][]string) {
	N := len(*input)
	for i := 1; i < N; i++ {
		for s := 0; s < i; s++ {
			moveUp(input, i-s)
		}
	}
}

func moveUp(input *[][]string, i int) {
	M := len((*input)[0])

	for c := 0; c < M; c++ {
		prevEl := (*input)[i-1][c]
		currentEl := (*input)[i][c]

		if currentEl == "O" {
			if prevEl == "." {
				(*input)[i-1][c] = "O"
				(*input)[i][c] = "."
			}
		}
	}
}

func readInput(fname string) [][]string {
	file, _ := os.Open(fname)
	scanner := bufio.NewScanner(file)

	mat := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		els := strings.Split(line, "")
		mat = append(mat, els)
	}

	return mat
}

func printLines(in [][]string) {
	for _, row := range in {
		fmt.Println(row)
	}
}
