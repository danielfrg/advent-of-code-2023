package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput("input-test.txt")

	fmt.Println(input)
}

func readInput(fname string) []puzzle {
	file, _ := os.Open(fname)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := make([]puzzle, 0)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		row := strings.Split(parts[0], "")
		seqStr := strings.Split(parts[1], ",")

		seq := make([]int, len(seqStr))
		for i, n := range seqStr {
			num, _ := strconv.Atoi(n)
			seq[i] = num
		}

		p := puzzle{
			row: row,
			seq: seq,
		}

		input = append(input, p)
	}
	return input
}

type puzzle struct {
	row []string
	seq []int
}
