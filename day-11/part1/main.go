package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	// input := readInput("input-test.txt")
	// printMat(input)
	// fmt.Println("----------")

	input := readInput("input.txt")

	input = expand(input, 2)
	// printMat(input)

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

	// fmt.Println(indices)

	distances := make([]int, 0)

	for i := 1; i < len(indices)+1; i++ {
		p1 := indices[i]
		// fmt.Println(i, p1)
		for j := i; j < len(indices)+1; j++ {
			p2 := indices[j]
			distance := math.Abs(float64(p2[1]-p1[1])) + math.Abs(float64(p2[0]-p1[0]))
			// fmt.Println(i, j, distance)
			distances = append(distances, int(distance))
		}
	}

	// fmt.Println(distances)

	fmt.Println("----------")
	fmt.Println(sum(distances))
}

func sum(arr []int) int {
	s := 0
	for _, v := range arr {
		s = s + v
	}
	return s
}

func expand(in [][]string, rate int) [][]string {
	// Expand rows
	ret := make([][]string, 0)

	for i := 0; i < len(in); i++ {
		line := in[i]
		test := countArr(line, ".")
		if test == len(line) {
			for e := 0; e < rate; e++ {
				ret = append(ret, line)
			}
		} else {
			ret = append(ret, line)
		}
	}

	// Expand cols

	target := make([]int, 0)
	for i := 0; i < len(in[0]); i++ {
		include := true
		for j := 0; j < len(in); j++ {
			if in[j][i] != "." {
				include = false
				break
			}
		}
		if include {
			target = append(target, i)
		}
	}

	// fmt.Println(target)

	for i := 0; i < len(ret); i++ {
		line := make([]string, 0)
		for j := 0; j < len(ret[i]); j++ {
			val := ret[i][j]
			if slices.Contains(target, j) {
				for r := 0; r < rate; r++ {
					line = append(line, ".")
				}
			} else {
				line = append(line, val)
			}
		}
		ret[i] = line
	}

	return ret
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
