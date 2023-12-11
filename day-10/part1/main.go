package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strings"
)

func main() {
	// input, starting := readInput("input-test.txt")
	// initialShape := "F" // Saw manually

	// input, starting := readInput("input-test-2.txt")
	// initialShape := "F" // Saw manually

	input, starting := readInput("input.txt")
	initialShape := "L" // Saw manually

	// fmt.Println(input)
	input[starting[0]][starting[1]] = initialShape

	matrix := make([][]int, len(input))
	for i := range matrix {
		matrix[i] = make([]int, len(input[0]))

		for j := 0; j < len(input[0]); j++ {
			matrix[i][j] = 0
		}
	}

	matrix[starting[0]][starting[1]] = -1

	iters := make([][]int, 0)
	iters = append(iters, starting)

	c := 0
	for {
		nIters := make([][]int, 0)
		fmt.Println(c, iters)
		for _, pos := range iters {
			n := iter(input, &matrix, pos)
			nIters = append(nIters, n...)
		}

		if len(nIters) == 0 {
			break
		}
		iters = nIters
		c++
		// if c == 4 {
		// 	fmt.Println(iters)
		// 	break
		// }
	}

	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}

	fmt.Println("----------")
	fmt.Println(c)
}

func iter(input [][]string, matrix *[][]int, pos []int) [][]int {
	X := pos[0]
	Y := pos[1]
	xLim := len(input[0])
	yLim := len(input)
	shape := input[X][Y]
	// fmt.Println(shape)

	value := (*matrix)[X][Y]
	if value == -1 {
		value = 0
	}
	next := make([][]int, 0)

	switch shape {
	case "|":
		// North
		if X-1 > 0 {
			if (*matrix)[X-1][Y] == 0 {
				(*matrix)[X-1][Y] = value + 1
				next = append(next, []int{X - 1, Y})
			}
		}
		// South
		if X+1 < yLim {
			if (*matrix)[X+1][Y] == 0 {
				(*matrix)[X+1][Y] = value + 1
				next = append(next, []int{X + 1, Y})
			}
		}
	case "-":
		// West
		if Y-1 >= 0 {
			if (*matrix)[X][Y-1] == 0 {
				(*matrix)[X][Y-1] = value + 1
				next = append(next, []int{X, Y - 1})
			}
		}
		// East
		if Y+1 < xLim {
			if (*matrix)[X][Y+1] == 0 {
				(*matrix)[X][Y+1] = value + 1
				next = append(next, []int{X, Y + 1})
			}
		}
	case "L":
		// North
		if X-1 > 0 {
			if (*matrix)[X-1][Y] == 0 {
				(*matrix)[X-1][Y] = value + 1
				next = append(next, []int{X - 1, Y})
			}
		}
		// East
		if Y+1 < xLim {
			if (*matrix)[X][Y+1] == 0 {
				(*matrix)[X][Y+1] = value + 1
				next = append(next, []int{X, Y + 1})
			}
		}
	case "J":
		// North
		if X-1 >= 0 {
			if (*matrix)[X-1][Y] == 0 {
				(*matrix)[X-1][Y] = value + 1
				next = append(next, []int{X - 1, Y})
			}
		}
		// West
		if Y-1 >= 0 {
			if (*matrix)[X][Y-1] == 0 {
				(*matrix)[X][Y-1] = value + 1
				next = append(next, []int{X, Y - 1})
			}
		}
	case "7":
		// South
		if X+1 < yLim {
			if (*matrix)[X+1][Y] == 0 {
				(*matrix)[X+1][Y] = value + 1
				next = append(next, []int{X + 1, Y})
			}
		}
		// West
		if Y-1 >= 0 {
			if (*matrix)[X][Y-1] == 0 {
				(*matrix)[X][Y-1] = value + 1
				next = append(next, []int{X, Y - 1})
			}
		}
	case "F":
		// South
		if X+1 < yLim {
			if (*matrix)[X+1][Y] == 0 {
				(*matrix)[X+1][Y] = value + 1
				next = append(next, []int{X + 1, Y})
			}
		}
		// East
		if Y+1 < xLim {
			if (*matrix)[X][Y+1] == 0 {
				(*matrix)[X][Y+1] = value + 1
				next = append(next, []int{X, Y + 1})
			}
		}
	}

	return next

}

func readInput(fname string) ([][]string, []int) {
	lines := utils.ReadLines(fname)

	rows := make([][]string, 0)
	starting := make([]int, 2)

	for i, line := range lines {
		cols := strings.Split(line, "")
		rows = append(rows, cols)

		idx := strings.Index(line, "S")
		if idx >= 0 {
			starting[0] = i
			starting[1] = idx
		}
	}

	return rows, starting
}
