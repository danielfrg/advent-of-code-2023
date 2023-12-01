package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
		coord := 10*first + last
		calibrationValues = append(calibrationValues, coord)

		fmt.Printf("%v -> %v\n", line, coord)
		// break
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println("----------")
	fmt.Printf("Total: %v\n", sum(calibrationValues))
}

func sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func firstDigit(line string) int {
	// First get the value and index of the FIRST numeric digit
	val := -1
	firstIndex := -1
	for i, r := range line {
		if unicode.IsDigit(r) {
			val = int(r - '0') // - '0' to get the int value of the ascii value
			firstIndex = i
			break
		}
	}

	// Now iterate the string digits and compare if they are BEFORE the numeric digit
	stringDigits := getStringsDigits()

	for _, digit := range stringDigits {
		index := strings.Index(line, digit)
		if index != -1 {
			if index < firstIndex {
				firstIndex = index
				val = stringToNum(digit)
			}
		}
	}

	return val
}

func lastDigit(line string) int {
	// First get the value and index of the LAST numeric digit
	val := -1
	lastIndex := -1

	for i := len(line) - 1; i >= 0; i -= 1 {
		r := line[i]
		if num, err := strconv.Atoi(string(r)); err == nil {
			lastIndex = i
			val = num
			break
		}
	}

	// Now iterate the string digits and compare if they are AFTER the numeric digit

	stringDigits := getStringsDigits()

	for _, digit := range stringDigits {
		index := strings.LastIndex(line, digit)
		if index != -1 {
			if index > lastIndex {
				lastIndex = index
				val = stringToNum(digit)
			}
		}
	}

	return val
}

func getStringsDigits() [10]string {
	digits := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	return digits
}

func stringToNum(str string) int {
	switch str {
	case "zero":
		return 0
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return -1
	}
}
