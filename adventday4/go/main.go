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
	part2Result := part2(lines)

	fmt.Printf("\npart 1 result: %v\n", part1Result)
	fmt.Printf("part 2 result: %v\n", part2Result)
}

func part1(lines []string) int {
	var overlappingWork int

	for _, line := range lines {
		elf := strings.Split(line, ",")

		left := makeRange(elf[0])
		right := makeRange(elf[1])

		if containsEither(left, right) {
			overlappingWork++
		}
	}

	return overlappingWork
}

func part2(lines []string) int {
	var overlappingWork int

	for _, line := range lines {
		elf := strings.Split(line, ",")

		left := makeRange(elf[0])
		right := makeRange(elf[1])

		if containsAny(left, right) {
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

func containsEither(left, right []int) bool {
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

func containsAny(left, right []int) bool {
	for _, leftValue := range left {
		if slices.Contains(right, leftValue) {
			return true
		}
	}
	for _, rightValue := range right {
		if slices.Contains(left, rightValue) {
			return true
		}
	}

	return false
}
