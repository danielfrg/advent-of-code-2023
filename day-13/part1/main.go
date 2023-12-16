package main

import (
	"advent-of-code-2023/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	mazes := readInput("input-test.txt")

	// printLines(mazes[0])
	// printLines(mazes[1])

	// printLines(mazes[0])
	// fmt.Println(getCol(mazes[0], 2))
	utils.Assert(findCols(mazes[0]), 5)

	// printLines(mazes[1])
	// fmt.Println(getRow(mazes[1], 2))
	utils.Assert(findRows(mazes[1]), 4)

	utils.Assert(calc(mazes), 405)

	mazes = readInput("input.txt")
	fmt.Println(calc(mazes))
}

func calc(mazes [][][]string) int {

	val := 0
	for i, maze := range mazes {
		cols := findCols(maze)
		if cols != -1 {
			val = val + cols
			continue
		}

		rows := findRows(maze)
		if rows != -1 {
			val = val + rows*100
			continue
		}

		fmt.Println("Maze #%s didnt find mirror", i)
	}

	return val
}

func findRows(mat [][]string) int {
	for i := 1; i < len(mat); i++ {
		c := 1
		test := true
		for {
			if i+c-1 == len(mat) || i-c == -1 {
				break
			}
			// fmt.Println(i+c-1, getRow(mat, i+c-1))
			// fmt.Println(i-c, getRow(mat, i-c))
			if getRow(mat, i+c-1) != getRow(mat, i-c) {
				// fmt.Println("!!")
				test = false
				break
			}
			c++
		}

		if test {
			return i
		}
		// break
	}

	return -1
}

func getRow(mat [][]string, idx int) string {
	vec := ""

	for i := 0; i < len(mat[0]); i++ {
		// vec = append(vec, mat[i][idx])
		vec = vec + mat[idx][i]
	}

	return vec
}

func findCols(mat [][]string) int {
	for i := 1; i < len(mat[0]); i++ {
		c := 1
		test := true
		for {
			if i+c-1 == len(mat[0]) || i-c == -1 {
				break
			}
			// fmt.Println(i+c-1, getCol(mat, i+c-1))
			// fmt.Println(i-c, getCol(mat, i-c))
			if getCol(mat, i+c-1) != getCol(mat, i-c) {
				// fmt.Println("!!")
				test = false
				break
			}
			c++
		}

		if test {
			return i
		}
		// break
	}

	return -1
}

func getCol(mat [][]string, idx int) string {
	vec := ""

	for i := 0; i < len(mat); i++ {
		// vec = append(vec, mat[i][idx])
		vec = vec + mat[i][idx]
	}

	return vec
}

func printLines(in [][]string) {
	for _, row := range in {
		fmt.Println(row)
	}
}

func readInput(fname string) [][][]string {
	file, _ := os.Open(fname)
	scanner := bufio.NewScanner(file)

	mazes := make([][][]string, 0)
	maze := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			mazes = append(mazes, maze)
			maze = make([][]string, 0)
			continue
		}

		row := strings.Split(line, "")
		maze = append(maze, row)
	}

	mazes = append(mazes, maze)
	return mazes
}
