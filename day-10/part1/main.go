package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

func main() {
	// input, starting := readInput("input.txt")
	// initialShape := "L" // Saw manually

	input, starting := readInput("input-test.txt")
	initialShape := "F" // Saw manually

	input[starting[0]][starting[1]] = initialShape

	fmt.Println(input)

	matrix := make([][]int, len(input))
	g := graph.New(graph.IntHash, graph.PreventCycles())
	for i := range matrix {
		matrix[i] = make([]int, len(input[0]))

		for j := 0; j < len(input[0]); j++ {
			myID := i*len(input[0]) + j
			_ = g.AddVertex(myID + 1)
			// fmt.Println(myID)
		}
	}

	link(input, g, &matrix, starting)

	file, _ := os.Create("./mygraph.gv")
	_ = draw.DOT(g, file)

	// v, _ := g.Vertex(startID)
	// fmt.Println(v)
	fmt.Println(g.Edges())
}

func link(input [][]string, g graph.Graph[int, int], matrix *[][]int, pos []int) {
	xLim := len(input[0])
	yLim := len(input)
	X := pos[0]
	Y := pos[1]
	shape := input[X][Y]

	xMult := 0
	if (X - 1) >= 0 {
		xMult = X
	}
	myID := xMult*xLim + Y + 1

	if (*matrix)[X][Y] == 1 {
		fmt.Println("Ignoring", myID)
		return
	}
	(*matrix)[X][Y] = 1

	// fmt.Println(X, Y, shape)

	switch shape {
	case "|":
		// North
		if X-1 > 0 {
			northID := (X-1)*xLim + Y + 1
			fmt.Println("|", "north", northID)
			_ = g.AddEdge(myID, northID)
			link(input, g, matrix, []int{X - 1, Y})
		}
		// South
		if X+1 < yLim {
			southID := (X+1)*xLim + Y + 1
			fmt.Println("|", "south", southID)
			_ = g.AddEdge(myID, southID)
			link(input, g, matrix, []int{X + 1, Y})
		}
	case "-":
		// West
		if Y-1 > 0 {
			westID := xMult*xLim + Y
			fmt.Println("-", "west", westID)
			_ = g.AddEdge(myID, westID)
			link(input, g, matrix, []int{X, Y - 1})
		}
		// East
		if Y+1 < xLim {
			eastID := xMult*xLim + Y + 2
			fmt.Println("-", "east", eastID)
			_ = g.AddEdge(myID, eastID)
			link(input, g, matrix, []int{X, Y + 1})
		}
	case "L":
		// North
		if X-1 > 0 {
			northID := (X-1)*xLim + Y + 1
			fmt.Println("L", "north", northID)
			_ = g.AddEdge(myID, northID)
			link(input, g, matrix, []int{X - 1, Y})
		}
		// East
		if Y+1 < xLim {
			eastID := xMult*xLim + Y + 2
			fmt.Println("L", "east", eastID)
			_ = g.AddEdge(myID, eastID)
			link(input, g, matrix, []int{X, Y + 1})
		}
	case "J":
		// North
		if X-1 > 0 {
			northID := (X-1)*xLim + Y + 1
			fmt.Println("J", "north", northID)
			_ = g.AddEdge(myID, northID)
			link(input, g, matrix, []int{X - 1, Y})
		}
		// West
		if Y-1 > 0 {
			westID := xMult*xLim + Y
			fmt.Println("J", "west", westID)
			_ = g.AddEdge(myID, westID)
			link(input, g, matrix, []int{X, Y - 1})
		}
	case "7":
		// South
		if X+1 < yLim {
			southID := (X+1)*xLim + Y + 1
			fmt.Println("7", "south", southID)
			_ = g.AddEdge(myID, southID)
			link(input, g, matrix, []int{X + 1, Y})
		}
		// West
		if Y-1 > 0 {
			westID := xMult*xLim + Y
			fmt.Println("7", "west", westID)
			_ = g.AddEdge(myID, westID)
			link(input, g, matrix, []int{X, Y - 1})
		}
	case "F":
		// South
		if X+1 < yLim {
			southID := (X+1)*xLim + Y + 1
			fmt.Println("F", "south", southID)
			_ = g.AddEdge(myID, southID)
			link(input, g, matrix, []int{X + 1, Y})
		}
		// East
		if Y+1 < xLim {
			eastID := xMult*xLim + Y + 2
			// fmt.Println("F", "east", eastID)
			_ = g.AddEdge(myID, eastID)
			link(input, g, matrix, []int{X, Y + 1})
		}
	}
}

func readInput(fname string) ([][]string, []int) {
	lines := utils.ReadLines(fname)

	rows := make([][]string, 0)
	starting := make([]int, 2)

	for i, line := range lines {
		cols := strings.Split(line, "")
		rows = append(rows, cols)

		idx := slices.IndexFunc(cols, func(c string) bool { return c == "S" })
		if idx > 0 {
			starting[0] = i
			starting[1] = idx
		}
	}

	return rows, starting
}
