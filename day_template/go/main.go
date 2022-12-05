package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("../input.txt")
	commands := strings.Split(string(data), "\n")

	part1Results := part1(commands)
	part2Results := part2(commands)

	fmt.Printf("\npart 1 results: %v\n", part1Results)
	fmt.Printf("part 2 results: %v\n", part2Results)
}

func part1(list []string) int {

	return len(list)
}

func part2(list []string) int {

	return len(list)
}
