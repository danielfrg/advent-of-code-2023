package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const LIMIT_RED int = 12
const LIMIT_GREEN int = 13
const LIMIT_BLUE int = 14

func main() {
	lines := readLines("input.txt")
	games := parseGames(lines)

	sum := 0
	for _, roundGames := range games {
		gameWithFewest := Game{}
		for i := 0; i < len(roundGames); i++ {
			game := roundGames[i]
			if game.red > gameWithFewest.red {
				gameWithFewest.red = game.red
			}
			if game.blue > gameWithFewest.blue {
				gameWithFewest.blue = game.blue
			}
			if game.green > gameWithFewest.green {
				gameWithFewest.green = game.green
			}
		}

		power := gameWithFewest.red * gameWithFewest.blue * gameWithFewest.green
		sum = sum + power
	}

	fmt.Println("----------")
	fmt.Println(sum)
}

type Game struct {
	red   int
	blue  int
	green int
}

func sum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func parseGames(lines []string) map[int][]Game {
	games := make(map[int][]Game)

	for _, line := range lines {
		parts := strings.Split(line, ":")
		gameParts := strings.Split(parts[0], " ")

		gameID, err := strconv.Atoi(gameParts[1])
		if err != nil {
			log.Panic("Error parsing Game ID")
			os.Exit(1)
		}

		gamesStr := strings.Split(parts[1], ";")

		thisLineGames := []Game{}
		for i := 0; i < len(gamesStr); i++ {
			game := Game{}

			singleGames := strings.Split(gamesStr[i], ",")

			for j := 0; j < len(singleGames); j++ {
				gameString := strings.Trim(singleGames[j], " ")
				draw := strings.Split(gameString, " ")
				value, err := strconv.Atoi(draw[0])

				if err != nil {
					log.Panic("Error parsing Game draw")
					os.Exit(1)
				}
				color := draw[1]

				if color == "red" {
					game.red = value
				}
				if color == "blue" {
					game.blue = value
				}
				if color == "green" {
					game.green = value
				}
			}

			thisLineGames = append(thisLineGames, game)
		}

		fmt.Printf("%v -> %v\n", line, thisLineGames)
		games[gameID] = thisLineGames
	}

	return games
}

func readLines(fname string) []string {
	lines := []string{}

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}
