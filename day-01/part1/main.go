package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	calibrationValues := make([]int, 0) // or slice := make([]int, elems)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		first := firstDigit(line)
		last := lastDigit(line)
		lineValue := 10*first + last
		calibrationValues = append(calibrationValues, lineValue)

		fmt.Printf("%v -> %v\n", line, lineValue)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println("----------")
	fmt.Printf("Total: %v\n", sum(calibrationValues))
}

func firstDigit(line string) int {
	for _, r := range line {
		if unicode.IsDigit(r) {
			return int(r - '0') // - '0' to get the int value of the ascii value
		}
	}
	return -1
}

func lastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i -= 1 {
		r := line[i]
		if num, err := strconv.Atoi(string(r)); err == nil {
			return num
		}
	}
	return -1
}

func sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}
