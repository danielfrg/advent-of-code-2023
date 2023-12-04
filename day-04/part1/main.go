package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadLines("input.txt")

	var scores []int
	for _, line := range lines {
		winning, played := parseGame(line)

		score := 0
		for _, num := range played {
			if slices.Contains(winning, num) {
				if score == 0 {
					score = 1
				} else {
					score = score * 2
				}
			}
		}

		scores = append(scores, score)
		fmt.Println(winning, played, score)
	}

	fmt.Println("----------")
	fmt.Println(utils.Sum(scores))
}

func parseGame(input string) ([]int, []int) {
	re := regexp.MustCompile(`^Card\s+(?P<idx>\d+): (?P<win>.+) \| (?P<played>.+)$`)
	match := re.FindStringSubmatch(input)

	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	// fmt.Println(result)
	win := strings.Split(result["win"], " ")
	played := strings.Split(result["played"], " ")
	return toIntSlice(win), toIntSlice(played)
}

func toIntSlice(input []string) []int {
	var s []int
	for _, number := range input {
		if number != "" {
			n, _ := strconv.Atoi(number)
			s = append(s, n)
		}
	}
	return s
}
