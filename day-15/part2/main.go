package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	utils.Assert(hash("HASH"), 52)
	utils.Assert(hash("rn=1"), 30)
	utils.Assert(hash("qp-"), 14)
	utils.Assert(hash("ot=7"), 231)
	utils.Assert(hash("rn"), 0)
	utils.Assert(hash("cm"), 0)

	sequences := readInput("input-test.txt")
	// fmt.Println(sequences)
	utils.Assert(solve(sequences), 145)

	sequences = readInput("input.txt")
	fmt.Println(solve(sequences))
}

func solve(sequences []Info) int {
	boxes := map[int][]Info{}

	// Init
	for i := 0; i < 256; i++ {
		boxes[i] = []Info{}
	}

	printBoxes(boxes)

	for _, info := range sequences {
		id := hash(info.Label)
		box := boxes[id]

		fmt.Printf("---\nIter %v=%v -> %v\n", info.Label, info.Value, id)

		if info.Type == Equals {
			// Find a box with the same label
			found := false
			for i, info2 := range box {
				if info.Label == info2.Label {
					found = true
					box[i] = info
					break
				}
			}
			if !found {
				boxes[id] = append(boxes[id], info)
			}
		}
		if info.Type == Minus {
			for i, info2 := range box {
				if info.Label == info2.Label {
					boxes[id] = append(boxes[id][:i], boxes[id][i+1:]...)
					break
				}
			}
		}
		// printBoxes(boxes)
	}

	// printBoxes(boxes)

	output := 0
	for i := 0; i < 256; i++ {
		box := boxes[i]
		for j, slot := range box {
			v := (i + 1) * (j + 1) * slot.Value
			fmt.Println((i + 1), j+1, slot.Value, v)
			output += v
		}
	}

	return output
}

func printBoxes(boxes map[int][]Info) {
	for i := 0; i < 256; i++ {
		box := boxes[i]
		if (len(box)) > 0 {
			fmt.Printf("Box %v: %v\n", i, box)
		}
	}
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

func readInput(fname string) []Info {
	content, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	str := string(content)
	parts := strings.Split(strings.Trim(str, "\n"), ",")

	ret := []Info{}

	for _, part := range parts {
		if strings.Contains(part, "=") {
			infoArr := strings.Split(part, "=")
			val, _ := strconv.Atoi(infoArr[1])
			info := Info{
				Label: infoArr[0],
				Value: val,
				Type:  Equals,
			}
			ret = append(ret, info)
		} else {
			infoArr := strings.Split(part, "-")
			info := Info{
				Label: infoArr[0],
				Value: -1,
				Type:  Minus,
			}
			ret = append(ret, info)
		}
	}

	return ret
}

type Type int

const (
	Equals Type = iota
	Minus
)

type Info struct {
	Label string
	Value int
	Type  Type
}
