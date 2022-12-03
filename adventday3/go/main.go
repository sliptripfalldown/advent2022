package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/maps"
)

func main() {
	ruckSack, _ := os.ReadFile("../input.txt")
	lines := strings.Split(string(ruckSack), "\n")

	part1Results := part1(lines)
	part2Results := part2(lines)

	fmt.Printf("\npart 1 result: %v\n", calc(part1Results))
	fmt.Printf("part 2 result: %v\n", calc(part2Results))
}

func part1(lines []string) []string {
	var results []string
	for _, line := range lines {
		half := len(line) / 2
		left, right := line[:half], line[half:]

		collector := make(map[string]bool)

		for _, leftLetter := range left {
			for _, rightLetter := range right {
				if leftLetter == rightLetter {
					collector[string(leftLetter)] = true
				}
			}
		}

		results = append(results, maps.Keys(collector)...)
	}

	return results
}

func calc(list []string) int {
	priority := generatePriority()

	var score int
	for _, item := range list {
		score += priority[item]
	}

	return score
}

func generatePriority() map[string]int {
	a := make(map[string]int, 52)

	counter := 1
	for m := 'a'; m <= 'z'; m++ {
		a[string(m)] = counter
		counter++
	}
	for m := 'A'; m <= 'Z'; m++ {
		a[string(m)] = counter
		counter++
	}

	return a
}

func part2(lines []string) []string {
	var results []string

	groupsOfThree := chunk(lines, 3)

	for _, group := range groupsOfThree {
		collector := make(map[string]bool)

		for _, leftLetter := range group[0] {
			for _, rightLetter := range group[1] {
				if leftLetter == rightLetter {
					for _, lastLetter := range group[2] {
						if leftLetter == lastLetter {
							collector[string(leftLetter)] = true
						}
					}
				}
			}
		}

		results = append(results, maps.Keys(collector)...)
	}

	return results
}

func chunk(lines []string, chunkSize int) [][]string {
	var group [][]string

	for i := 0; i < len(lines); i += chunkSize {
		end := i + chunkSize

		if end > len(lines) {
			end = len(lines)
		}

		group = append(group, lines[i:end])
	}

	return group
}
