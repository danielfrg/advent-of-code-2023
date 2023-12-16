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
	// utils.Assert(calcMaze(mazes[0]), 5)
	// utils.Assert(iterate(mazes[0]), 300)

	// printLines(mazes[1])
	// utils.Assert(calcMaze(mazes[1]), 400)
	// utils.Assert(iterate(mazes[1]), 100)

	// mazes = readInput("input-test2.txt")
	// fmt.Println(calcMaze(mazes[0]))
	// fmt.Println(iterate(mazes[0]))
	// utils.Assert(calcBase(mazes), 32035)

	mazes = readInput("input.txt")
	utils.Assert(calcBase(mazes), 32035)
	fmt.Println(calcMaze(mazes[0]))
	fmt.Println(iterate(mazes[0]))
	// fmt.Println(calc(mazes))
}

func calc(mazes [][][]string) int {
	val := 0
	for i, maze := range mazes {
		v := iterate(maze)

		if v == -1 {
			fmt.Printf("Maze #%v didnt find mirror\n", i)
			continue
		}

		val = val + v
	}
	return val
}

func calcBase(mazes [][][]string) int {
	val := 0
	for i, maze := range mazes {
		v := calcMaze(maze)

		if v == -1 {
			fmt.Printf("Maze #%v didnt find mirror\n", i)
			continue
		}

		val = val + v
	}

	return val
}

func iterate(maze [][]string) int {
	baseAnswer := calcMaze(maze)

	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[0]); j++ {
			m := flip(maze, i, j)
			// fmt.Println("---", i, j)
			// printLines(m)

			v := calcMaze2(m, baseAnswer)
			// fmt.Println(i, j, v)

			if v != -1 && v != baseAnswer {
				return v
			}
			flip(maze, i, j)
		}
	}

	return -1
}

func calcMaze(maze [][]string) int {
	r := findRows(maze)

	if r != -1 {
		return r * 100
	}

	c := findCols(maze)

	if c != -1 {
		return c
	}

	return -1
}

func calcMaze2(maze [][]string, base int) int {
	r := findRows(maze)

	if r != -1 && r*100 != base {
		return r * 100
	}

	c := findCols(maze)

	if c != -1 && r != base {
		return c
	}

	return -1
}

func flip(mat [][]string, i int, j int) [][]string {
	v := mat[i][j]
	if v == "." {
		mat[i][j] = "#"
	}
	if v == "#" {
		mat[i][j] = "."
	}

	return mat
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
