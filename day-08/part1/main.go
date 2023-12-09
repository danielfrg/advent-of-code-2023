package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"os"
	"regexp"
)

func main() {
	lines := utils.ReadLines("input.txt")

	steps, instructions := parseInput(lines)
	// fmt.Println(steps)
	// fmt.Println(instructions)

	count := 0
	location := "AAA"
	for {
		if location == "ZZZ" {
			break
		}
		inst := instructions[location]
		nextStep := steps[count%len(steps)]
		if nextStep == 'L' {
			location = inst.left
		} else if nextStep == 'R' {
			location = inst.right
		} else {
			fmt.Printf("Unkwon step: %v\n", nextStep)
			os.Exit(1)
		}
		// fmt.Println(nextStep, location)
		count++
	}

	fmt.Println(count)
}

type instruction struct {
	start string
	left  string
	right string
}

func parseInput(lines []string) (string, map[string]instruction) {
	steps := lines[0]

	re := regexp.MustCompile(`^([A-Z]+)\s=\s\(([A-Z]+), ([A-Z]+)\)`)

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
