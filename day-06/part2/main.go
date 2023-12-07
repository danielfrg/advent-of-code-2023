package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadFile("input.txt")
	time, record := parseInput(input)
	fmt.Println(time, record)

	var values []int
	counter := 0

	for t := 0; t <= time; t++ {
		speed := t
		distance := (time - t) * speed
		// fmt.Println(distance)

		if distance > record {
			counter++
		}
	}

	values = append(values, counter)

	fmt.Println(utils.Multiply(values))
}

func parseInput(input string) (int, int) {
	re := regexp.MustCompile(`Time:\s+([\d ]+)`)
	match := re.FindStringSubmatch(input)

	re2 := regexp.MustCompile(`Distance:\s+([\d ]+)`)
	match2 := re2.FindStringSubmatch(input)

	timeStr := strings.ReplaceAll(match[1], " ", "")
	recordStr := strings.ReplaceAll(match2[1], " ", "")

	time, _ := strconv.Atoi(timeStr)
	record, _ := strconv.Atoi(recordStr)
	return time, record
}
