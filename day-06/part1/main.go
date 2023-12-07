package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"regexp"
)

func main() {
	input := utils.ReadFile("input.txt")
	times, records := parseInput(input)

	var values []int

	for i := 0; i < len(times); i++ {
		time := times[i]
		record := records[i]

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
	}

	fmt.Println(utils.Multiply(values))
}

func parseInput(input string) ([]int, []int) {
	re := regexp.MustCompile(`Time:\s+([\d ]+)`)
	match := re.FindStringSubmatch(input)

	re2 := regexp.MustCompile(`Distance:\s+([\d ]+)`)
	match2 := re2.FindStringSubmatch(input)

	times := utils.Split(match[1], ' ')
	distances := utils.Split(match2[1], ' ')
	return utils.ToIntSlice(times), utils.ToIntSlice(distances)
}
