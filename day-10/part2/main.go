package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"math"
	"strings"
)

func main() {
	test := 0

	test = solve("input-test-part2.txt", "F")
	utils.Assert(test, 4)

	test = solve("input-test-part2-large.txt", "F")
	utils.Assert(test, 8)

	test = solve("input.txt", "F")
	fmt.Println(test)
}

func solve(fname string, initialShape string) int {
	input, starting := readInput(fname)
	input[starting[0]][starting[1]] = initialShape

	loop := [][]int{starting}
	pos := starting
	dir := ""
	c := 0
	for {
		// fmt.Println(input[pos[0]][pos[1]], pos, dir)
		pos, dir = iter(input, pos, dir)

		c++
		if pos[0] == starting[0] && pos[1] == starting[1] {
			break
		}

		loop = append(loop, pos)
	}

	// fmt.Println(loop)

	// Shoelace formula
	area := interiorArea(loop)

	// pick's theorem - find the number of points in a shape given its area
	numInteriorPoints := area - 0.5*float64(len(loop)) + 1
	fmt.Println(numInteriorPoints)
	return int(numInteriorPoints)
}

func interiorArea(points [][]int) float64 {
	// Shoelace formula
	// https://en.wikipedia.org/wiki/Shoelace_formula

	sum := 0
	p0 := points[len(points)-1]
	for _, p1 := range points {
		// sum += p0.y*p1.x - p0.x*p1.y
		sum += p0[1]*p1[0] - p0[0]*p1[1]
		p0 = p1
		// fmt.Println("_", p0[1]*p1[0]-p0[0]*p1[1])
	}
	return math.Abs(float64(sum)) / 2

}

func iter(input [][]string, pos []int, dir string) ([]int, string) {
	X := pos[0]
	Y := pos[1]
	letter := input[X][Y]

	// fmt.Println(letter, dir)

	switch letter {
	case "|":
		if dir == "D" {
			return []int{X + 1, Y}, "D"
		}

		if dir == "U" {
			return []int{X - 1, Y}, "U"
		}

	case "-":
		if dir == "R" {
			return []int{X, Y + 1}, "R"
		}

		if dir == "L" {
			return []int{X, Y - 1}, "L"
		}
	case "L":
		if dir == "D" {
			return []int{X, Y + 1}, "R"
		}

		if dir == "L" {
			return []int{X - 1, Y}, "U"
		}

	case "J":
		if dir == "D" {
			return []int{X, Y - 1}, "L"
		}

		if dir == "R" {
			return []int{X - 1, Y}, "U"
		}

	case "7":
		if dir == "R" {
			return []int{X + 1, Y}, "D"
		}

		if dir == "U" {
			return []int{X, Y - 1}, "L"
		}

	case "F":
		if dir == "" || dir == "U" {
			return []int{X, Y + 1}, "R"
		}

		if dir == "L" {
			return []int{X + 1, Y}, "D"
		}
	}

	return nil, ""
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
