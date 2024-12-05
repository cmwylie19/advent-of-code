package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func separateByNewLine(input string) ([]int, []int) {
	leftColumn := make([]int, 0)
	rightColumn := make([]int, 0)
	line := strings.Split(input, "\n")
	for _, l := range line {
		if l != "" {
			stringVals := strings.Split(l, "   ")
			leftValue, err := strconv.Atoi(stringVals[0])
			if err != nil {
				panic(err)
			}
			rightValue, err := strconv.Atoi(stringVals[1])
			if err != nil {
				panic(err)
			}
			leftColumn = append(leftColumn, leftValue)
			rightColumn = append(rightColumn, rightValue)
		}
	}
	return leftColumn, rightColumn
}

func readFile() string {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func distanceBetweenColumns(leftColumn, rightColumn []int) int {
	var difference int
	for i := 0; i < len(leftColumn); i++ {
		difference += abs(rightColumn[i] - leftColumn[i])
	}
	return difference
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	var input string
	var leftColumn, rightColumn []int

	// Get input
	input = readFile()

	// Sort input into left and right columns
	leftColumn, rightColumn = separateByNewLine(input)

	// Sort the columns
	sort.Ints(leftColumn)
	sort.Ints(rightColumn)

	difference := distanceBetweenColumns(leftColumn, rightColumn)
	println(difference)
}
