package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type game struct {
	id      int
	winning []int
	played  []int
}

func main() {
	lines := utils.ReadLines("input.txt")

	var games []game
	winningCount := map[int]int{}
	cardCount := map[int]int{}

	for _, line := range lines {
		game := parseGame(line)
		games = append(games, game)
		winningCount[game.id] = game.score(game)

		// Count the original card
		cardCount[game.id] = 1
	}

	// Calculate the copies and add it to cardCount
	for _, game := range games {
		countToAdd := winningCount[game.id]

		for j := game.id + 1; j <= game.id+countToAdd; j++ {
			// add the number of times we have copies of this card
			cardCount[j] += 1 * cardCount[game.id]
		}
	}

	fmt.Println("----------")
	count := 0
	for _, i := range cardCount {
		count += i
	}
	fmt.Println(count)
}

func (g game) score(game game) int {
	score := 0
	for _, num := range game.played {
		if slices.Contains(game.winning, num) {
			score++
		}
	}
	return score
}

func parseGame(input string) game {
	re := regexp.MustCompile(`^Card\s+(?P<id>\d+): (?P<win>.+) \| (?P<played>.+)$`)
	match := re.FindStringSubmatch(input)

	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	id, _ := strconv.Atoi(result["id"])
	win := strings.Split(result["win"], " ")
	played := strings.Split(result["played"], " ")
	return game{id: id, winning: toIntSlice(win), played: toIntSlice(played)}
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
