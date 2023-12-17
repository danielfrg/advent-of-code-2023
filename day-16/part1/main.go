package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// input := readInput("input-test.txt")
	// counts := solve(input)
	// utils.Assert(calc(counts), 46)

	input := readInput("input.txt")
	// printLines(input)
	counts := solve(input)
	// printLinesInt(counts[:5])
	fmt.Println(calc(counts))
}

func calc(counts [][]int) int {
	X := len(counts)
	Y := len(counts[0])

	c := 0

	for _, row := range counts {
		for _, v := range row {
			if v == 0 {
				c++
			}
		}
	}

	return X*Y - c
}

func solve(input [][]string) [][]int {
	X := len(input)
	Y := len(input[0])

	counts := [][]int{}
	pairs := map[string]bool{}

	for i := 0; i < X; i++ {
		a := make([]int, Y)
		counts = append(counts, a)
	}
	// printLines(input)
	// printLinesInt(counts)

	pos := [2]int{0, 0}
	initial := Move{}

	symbol := input[pos[0]][pos[1]]
	if symbol == "." || symbol == "-" {
		initial = Move{
			x:   0,
			y:   0,
			dir: "R",
		}
	} else if symbol == "\\" || symbol == "|" {
		initial = Move{
			x:   0,
			y:   0,
			dir: "D",
		}
	}
	moves := []Move{initial}

	for i := 0; i < len(moves); i++ {
		move := moves[i]

		fmt.Printf("Move #%v %v\n", i, move)
		if move.x < 0 || move.y < 0 || move.x > X || move.y > Y {
			continue
		}

		counts[move.y][move.x]++

		if move.dir == "R" {
			if move.x+1 < X {
				symbol := input[move.y][move.x+1]
				_, ok := pairs[fmt.Sprintf("%v-%v-R", move.x+1, move.y)]
				if ok {
					continue
				}
				pairs[fmt.Sprintf("%v-%v-R", move.x+1, move.y)] = true

				if symbol == "." || symbol == "-" {
					// Pass trough
					moves = append(moves, Move{x: move.x + 1, y: move.y, dir: "R"})
				}
				if (symbol) == "|" {
					// Split
					moves = append(moves, Move{x: move.x + 1, y: move.y, dir: "U"})
					moves = append(moves, Move{x: move.x + 1, y: move.y, dir: "D"})
				}
				if (symbol) == "\\" {
					moves = append(moves, Move{x: move.x + 1, y: move.y, dir: "D"})
				}
				if (symbol) == "/" {
					moves = append(moves, Move{x: move.x + 1, y: move.y, dir: "U"})
				}
			}
		}
		if move.dir == "L" {
			if move.x-1 >= 0 {
				symbol := input[move.y][move.x-1]
				_, ok := pairs[fmt.Sprintf("%v-%v-L", move.x-1, move.y)]
				if ok {
					continue
				}
				pairs[fmt.Sprintf("%v-%v-L", move.x-1, move.y)] = true

				if symbol == "." || symbol == "-" {
					// Pass trough
					moves = append(moves, Move{x: move.x - 1, y: move.y, dir: "L"})
				}
				if (symbol) == "|" {
					// Split
					moves = append(moves, Move{x: move.x - 1, y: move.y, dir: "U"})
					moves = append(moves, Move{x: move.x - 1, y: move.y, dir: "D"})
				}
				if (symbol) == "\\" {
					moves = append(moves, Move{x: move.x - 1, y: move.y, dir: "U"})
				}
				if (symbol) == "/" {
					moves = append(moves, Move{x: move.x - 1, y: move.y, dir: "D"})
				}
			}
		}
		if move.dir == "D" {
			if move.y+1 < Y {
				symbol := input[move.y+1][move.x]
				_, ok := pairs[fmt.Sprintf("%v-%v-D", move.x, move.y+1)]
				if ok {
					continue
				}
				pairs[fmt.Sprintf("%v-%v-D", move.x, move.y+1)] = true

				if symbol == "." || symbol == "|" {
					// Pass trough
					moves = append(moves, Move{x: move.x, y: move.y + 1, dir: "D"})
				}
				if (symbol) == "-" {
					// Split
					moves = append(moves, Move{x: move.x, y: move.y + 1, dir: "L"})
					moves = append(moves, Move{x: move.x, y: move.y + 1, dir: "R"})
				}
				if (symbol) == "\\" {
					moves = append(moves, Move{x: move.x, y: move.y + 1, dir: "R"})
				}
				if (symbol) == "/" {
					moves = append(moves, Move{x: move.x, y: move.y + 1, dir: "L"})
				}
			}
		}
		if move.dir == "U" {
			if move.y-1 >= 0 {
				symbol := input[move.y-1][move.x]
				_, ok := pairs[fmt.Sprintf("%v-%v-U", move.x, move.y-1)]
				if ok {
					continue
				}
				pairs[fmt.Sprintf("%v-%v-U", move.x, move.y-1)] = true

				if symbol == "." || symbol == "|" {
					// Pass trough
					moves = append(moves, Move{x: move.x, y: move.y - 1, dir: "U"})
				}
				if (symbol) == "-" {
					// Split
					moves = append(moves, Move{x: move.x, y: move.y - 1, dir: "L"})
					moves = append(moves, Move{x: move.x, y: move.y - 1, dir: "R"})
				}
				if (symbol) == "\\" {
					moves = append(moves, Move{x: move.x, y: move.y - 1, dir: "L"})
				}
				if (symbol) == "/" {
					moves = append(moves, Move{x: move.x, y: move.y - 1, dir: "R"})
				}
			}
		}
		// fmt.Println(moves)
		// printLinesInt(counts[:8])
		// if i == 10 {
		// 	break
		// }
		// fmt.Println("---")
		// break
	}

	return counts
}

type Move struct {
	x   int
	y   int
	dir string
}

func readInput(fname string) [][]string {
	file, _ := os.Open(fname)

	scanner := bufio.NewScanner(file)

	mat := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")
		mat = append(mat, parts)
	}

	return mat
}

func printLines(in [][]string) {
	for _, line := range in {
		fmt.Println(line)
	}
}

func printLinesInt(in [][]int) {
	for _, line := range in {
		fmt.Println(line)
	}
}
