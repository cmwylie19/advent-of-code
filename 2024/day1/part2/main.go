package main

import (
	"fmt"
	"os"
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

func countOccurences(left []int, right []int) int {
	freq := make(map[int]int)
	for _, num := range right {
		freq[num]++
	}

	total := 0
	for _, num := range left {
		total += num * freq[num]
	}

	return total
}
func readFile() string {
	bytes, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
func main() {
	var input string
	var leftColumn, rightColumn []int

	// Get input
	input = readFile()

	// Sort input into left and right columns
	leftColumn, rightColumn = separateByNewLine(input)
	fmt.Println(countOccurences(leftColumn, rightColumn))

}
