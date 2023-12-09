package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	lines := utils.ReadLines("input.txt")
	// lines := utils.ReadLines("input-test-part2.txt")

	steps, instructions := parseInput(lines)

	// Find staring points
	searching := make(map[string]string, 0)
	lengths := make(map[string]int, 0)
	for key := range instructions {
		if strings.HasSuffix(key, "A") {
			searching[key] = key
			lengths[key] = 0
			// break
		}
	}

	// fmt.Println(searching)

	// Iterate the searching nodes while navigating

	count := 0

	for len(searching) > 0 {
		movement := steps[count%len(steps)]

		for key, node := range searching {
			inst := instructions[node]
			nextLocation := navigate(inst, movement)
			searching[key] = nextLocation

			if strings.HasSuffix(nextLocation, "Z") {
				lengths[key] = count + 1
				delete(searching, key)
			}
		}

		count++
	}

	fmt.Println(lengths)

	// Fund the LCM of the found lengths
	values := make([]int, 0, len(lengths))

	for _, value := range lengths {
		values = append(values, value)
	}

	fmt.Println(LCM(values[0], values[1], values[2:]...))
}

func navigate(inst instruction, movement byte) string {
	if movement == 'L' {
		return inst.left
	} else if movement == 'R' {
		return inst.right
	} else {
		fmt.Printf("Unknown movement: %v\n", movement)
		os.Exit(1)
	}
	return ""
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

type instruction struct {
	start string
	left  string
	right string
}

func parseInput(lines []string) (string, map[string]instruction) {
	steps := lines[0]

	re := regexp.MustCompile(`^([A-Z0-9]+)\s=\s\(([A-Z0-9]+), ([A-Z0-9]+)\)`)

	instructions := make(map[string]instruction, 0)
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		match := re.FindStringSubmatch(line)

		if len(match) == 4 {
			inst := instruction{
				start: match[1],
				left:  match[2],
				right: match[3],
			}
			instructions[match[1]] = inst
		}
	}
	return steps, instructions
}
