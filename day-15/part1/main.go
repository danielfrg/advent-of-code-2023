package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	utils.Assert(hash("HASH"), 52)
	utils.Assert(hash("rn=1"), 30)
	utils.Assert(hash("qp-"), 14)
	utils.Assert(hash("ot=7"), 231)

	sequences := readInput("input-test.txt")
	// fmt.Println(sequences)

	utils.Assert(calc(sequences), 1320)

	sequences = readInput("input.txt")
	fmt.Println(calc(sequences))
}

func calc(sequences []string) int {
	sum := 0

	for _, seq := range sequences {
		// fmt.Println(seq, hash(seq))
		sum = sum + hash(seq)
	}

	return sum
}

func hash(input string) int {
	val := 0

	C1 := 17
	C2 := 256

	for _, rune := range input {
		if rune == '\n' {
			continue
		}
		val = val + int(rune)
		// fmt.Println(i, val)
		val = val * C1
		// fmt.Println(i, val)
		val = val % C2
		// fmt.Println(i, val)
	}

	return val
}

func readInput(fname string) []string {
	content, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	str := string(content)
	parts := strings.Split(str, ",")
	return parts
}
