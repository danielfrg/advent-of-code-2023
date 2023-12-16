package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// input := readInput("input-test.txt")
	// utils.Assert(result(&input), 64)
	// result := result(&input)
	// fmt.Println(result)

	input := readInput("input.txt")
	result := result(&input)
	fmt.Println(result)
}

func result(input *[][]string) int {
	grids := map[string]int{}

	N := 1000000000
	// N := 100
	jumped := false
	i := 0
	for i < N {
		i++
		cycle(input)
		score := calc(input)
		fmt.Println(i, score)
		// calc(&input)
		h := hash(input)
		if !jumped {
			val, ok := grids[h]
			if ok {
				// cycle_length := i - val - 1  # For example
				cycle_length := i - val
				increase := (N - i) / cycle_length
				fmt.Println("!!", score, val, cycle_length, N-i, increase)
				i = i + increase*cycle_length
				jumped = true
			} else {
				grids[h] = i
			}
		}
	}

	return calc(input)
}

func hash(grid *[][]string) string {
	hash := ""
	for i := 0; i < len(*(grid)); i++ {
		for j := 0; j < len((*(grid))[0]); j++ {
			hash = hash + (*grid)[i][j]
		}
	}
	return hash
}

func cycle(input *[][]string) {
	tiltNorth(input)
	// fmt.Println("---")
	// printLines(*input)

	tiltWest(input)
	// fmt.Println("---")
	// printLines(*input)

	tiltSouth(input)
	// fmt.Println("---")
	// printLines(*input)

	tiltEast(input)
	// fmt.Println("---")
	// printLines(*input)
}

func tiltWest(input *[][]string) {
	M := len((*input)[0])
	for i := 1; i < M; i++ {
		for s := 0; s < i; s++ {
			moveWest(input, i-s)
		}
	}
}

func moveWest(input *[][]string, i int) {
	N := len(*input)

	for c := 0; c < N; c++ {
		prevEl := (*input)[c][i-1]
		currentEl := (*input)[c][i]

		if currentEl == "O" {
			if prevEl == "." {
				(*input)[c][i-1] = "O"
				(*input)[c][i] = "."
			}
		}
	}
}

func tiltSouth(input *[][]string) {
	N := len(*input)
	for i := N - 2; i >= 0; i-- {
		for s := 0; s < (N-1)-i; s++ {
			moveSouth(input, i+s)
		}
	}
}

func moveSouth(input *[][]string, i int) {
	M := len((*input)[0])

	for c := 0; c < M; c++ {
		prevEl := (*input)[i+1][c]
		currentEl := (*input)[i][c]

		if currentEl == "O" {
			if prevEl == "." {
				(*input)[i+1][c] = "O"
				(*input)[i][c] = "."
			}
		}
	}
}

func tiltEast(input *[][]string) {
	M := len((*input)[0])
	for i := M - 2; i >= 0; i-- {
		for s := 0; s < (M-1)-i; s++ {
			moveEast(input, i+s)
		}
	}
}

func moveEast(input *[][]string, i int) {
	N := len(*input)

	for c := 0; c < N; c++ {
		prevEl := (*input)[c][i+1]
		currentEl := (*input)[c][i]

		if currentEl == "O" {
			if prevEl == "." {
				(*input)[c][i+1] = "O"
				(*input)[c][i] = "."
			}
		}
	}
}

func tiltNorth(input *[][]string) {
	N := len(*input)
	for i := 1; i < N; i++ {
		for s := 0; s < i; s++ {
			moveNorth(input, i-s)
		}
	}
}

func moveNorth(input *[][]string, i int) {
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
