package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Assert(val int, expected int) {
	if val != expected {
		log.Fatalf("expected %v but got %v", expected, val)
	}
}

func Split(input string, sep rune) []string {
	// Split and remove empty items
	splitFn := func(c rune) bool {
		return c == sep
	}
	return strings.FieldsFunc(input, splitFn)
}

func Sum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func Multiply(arr []int) int {
	mult := 1
	for i := 0; i < len(arr); i++ {
		mult *= arr[i]
	}
	return mult
}

func ReadFile(fname string) string {
	content, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return string(content)
}

func ReadLines(fname string) []string {
	lines := []string{}

	file, err := os.Open(fname)

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

func ToIntSlice(input []string) []int {
	var s []int
	for _, number := range input {
		if number != "" {
			n, _ := strconv.Atoi(number)
			s = append(s, n)
		}
	}
	return s
}
