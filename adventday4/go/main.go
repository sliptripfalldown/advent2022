package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	coordinates, _ := os.ReadFile("../input.txt")
	lines := strings.Split(string(coordinates), "\n")

	part1Result := part1(lines)

	fmt.Printf("part 1 result: %v", part1Result)
}

func part1(lines []string) int {

	return calculateOverlap(lines)
}

func calculateOverlap(lines []string) int {
	var overlappingWork int

	for _, line := range lines {
		elf := strings.Split(line, ",")

		left := makeRange(elf[0])
		right := makeRange(elf[1])

		if deepContains(left, right) {
			overlappingWork++
		}
	}

	return overlappingWork
}

func makeRange(elf string) []int {
	section := strings.Split(elf, "-")
	min, _ := strconv.Atoi(section[0])
	max, _ := strconv.Atoi(section[1])

	var a []int
	for i := min; i <= max; i++ {
		a = append(a, i)
	}

	return a
}

func deepContains(left, right []int) bool {
	leftResult := true
	rightResult := true
	for _, leftValue := range left {
		if !slices.Contains(right, leftValue) {
			leftResult = false
		}
	}

	for _, rightValue := range right {
		if !slices.Contains(left, rightValue) {
			rightResult = false
		}
	}
	return leftResult || rightResult
}
