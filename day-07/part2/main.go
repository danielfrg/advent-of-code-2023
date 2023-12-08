package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadLines("input.txt")

	var games []game

	for _, line := range lines {
		game := parseLine(line)
		games = append(games, game)
	}

	// Sort the games by the logic described: handType and then highest card in order
	sort.Slice(games, func(i, j int) bool {
		type1 := games[i].handType()
		type2 := games[j].handType()
		if type1 != type2 {
			return games[i].handType() < games[j].handType()
		}

		for c := 0; c < len(games[i].hand); c++ {
			card1 := games[i].hand[c]
			card2 := games[j].hand[c]

			comp := compareCard(card1, card2)
			if comp != 0 {
				return comp < 0
			}
		}

		return false
	})

	var values []int
	for rank, game := range games {
		value := (rank + 1) * game.bid
		values = append(values, value)
	}

	fmt.Println(utils.Sum(values))
}

type game struct {
	hand []string
	bid  int
}

const (
	high_card  = iota // == 0
	one_pair   = iota // == 1
	two_pair   = iota // == 2
	three_kind = iota // == 3
	full_house = iota // == 4
	four_kind  = iota // == 5
	five_kind  = iota // == 6
)

func (g game) handType() int {
	originalCount := countEls(g.hand)

	jokers, ok := originalCount["J"]
	if !ok {
		jokers = 0
	}

	handWithoutJokers := make([]string, 0)
	for _, c := range g.hand {
		if c != "J" {
			handWithoutJokers = append(handWithoutJokers, c)
		}
	}

	count := countEls(handWithoutJokers)

	// Sort count map keys
	keys := make([]string, 0, len(count))
	for key := range count {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return count[keys[i]] > count[keys[j]]
	})

	// Give ranking of the hand

	if jokers == 5 {
		return five_kind
	}

	if count[keys[0]]+jokers == 5 {
		return five_kind
	}

	if count[keys[0]]+jokers == 4 {
		return four_kind
	}

	if count[keys[0]]+jokers == 3 && count[keys[1]] == 2 {
		return full_house
	}

	if count[keys[0]]+jokers == 3 && count[keys[1]] == 1 {
		return three_kind
	}

	if count[keys[0]]+jokers == 2 && count[keys[1]] == 2 {
		return two_pair
	}

	if count[keys[0]]+jokers == 2 && count[keys[1]] == 1 {
		return one_pair
	}

	return high_card
}

func compareCard(card1 string, card2 string) int {
	return cardValue(card1) - cardValue(card2)
}

func cardValue(card string) int {
	switch card {
	case "T":
		return 10
	case "J":
		return 1
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	default:
		val, _ := strconv.Atoi(card)
		return val
	}
}

func countEls(els []string) map[string]int {
	m := make(map[string]int)

	for _, el := range els {
		_, ok := m[el]
		// If the key exists
		if ok {
			m[el]++
		} else {
			m[el] = 1
		}
	}

	return m
}

func parseLine(line string) game {
	parts := strings.Split(line, " ")

	bid, _ := strconv.Atoi(parts[1])
	game_ := game{
		hand: strings.Split(parts[0], ""),
		bid:  bid,
	}
	return game_
}
