package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	// input := readInput("input-test.txt")
	// printMat(input)
	// fmt.Println("----------")

	input := readInput("input.txt")

	indices := make(map[int][]int, 0)

	c := 0
	for i, line := range input {
		for j, char := range line {
			if char == "#" {
				indices[c+1] = []int{i, j}
				c++
			}
		}
	}

	spacesX, spacesY := expand(input)

	fmt.Println(spacesX, spacesY)

	rate := 1000000
	distances := make([]int, 0)

	for i := 1; i < len(indices)+1; i++ {
		p1 := indices[i]
		// fmt.Println(i, p1)
		for j := i; j < len(indices)+1; j++ {
			p2 := indices[j]
			distance := int(math.Abs(float64(p2[1]-p1[1])) + math.Abs(float64(p2[0]-p1[0])))

			extraX := countBetween(spacesX, p1[0], p2[0])
			extraY := countBetween(spacesY, p1[1], p2[1])

			distance = distance + extraX*(rate-1)
			distance = distance + extraY*(rate-1)

			// fmt.Println(i, j, distance)
			distances = append(distances, distance)
		}
	}

	// fmt.Println(distances)

	fmt.Println("----------")
	fmt.Println(sum(distances))
}

func countBetween(arr []int, p1 int, p2 int) int {
	low := 0
	high := 0
	if p1 < p2 {
		low = p1
		high = p2
	} else {
		low = p2
		high = p1
	}

	c := 0
	for _, v := range arr {
		if v > low && v < high {
			c++
		}
	}

	return c
}

func expand(in [][]string) ([]int, []int) {
	// Expand rows
	targetX := make([]int, 0)

	for i := 0; i < len(in); i++ {
		line := in[i]
		test := countArr(line, ".")
		if test == len(line) {
			targetX = append(targetX, i)
		}
	}

	// Expand cols

	targetY := make([]int, 0)
	for i := 0; i < len(in[0]); i++ {
		include := true
		for j := 0; j < len(in); j++ {
			if in[j][i] != "." {
				include = false
				break
			}
		}
		if include {
			targetY = append(targetY, i)
		}
	}

	return targetX, targetY
}

func sum(arr []int) int {
	s := 0
	for _, v := range arr {
		s = s + v
	}
	return s
}

func insert(arr []string, index int, val string, repeat int) []string {
	if len(arr) == index {
		return append(arr, val)
	}
	arr = append(arr[:index+repeat], arr[index:]...)
	for r := 0; r < repeat; r++ {
		arr[index+r] = val
	}
	return arr
}

func countArr(in []string, char string) int {
	count := 0
	for _, val := range in {
		if val == char {
			count++
		}
	}
	return count
}

func printMat(in [][]string) {
	for i := 0; i < len(in); i++ {
		fmt.Println(in[i])
	}
}

func readInput(fname string) [][]string {
	input := make([][]string, 0)

	file, err := os.Open(fname)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		points := strings.Split(line, "")
		input = append(input, points)
	}
	return input
}
