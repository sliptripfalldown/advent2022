package main

import (
	"fmt"
	"os"
	"strconv"
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
	signal := []int{1}
	for _, line := range list {
		if strings.Contains(line, "noop") {
			signal = append(signal, signal[len(signal)-1])
		} else {
			part := strings.Split(line, " ")
			value, _ := strconv.Atoi(part[1])
			signal = append(signal, signal[len(signal)-1])
			signal = append(signal, signal[len(signal)-1]+value)
		}
	}

	sum := signal[19]*20 + signal[59]*60 + signal[99]*100 + signal[139]*140 + signal[179]*180 + signal[219]*220
	fmt.Println("Answer 1:", sum)

	return sum
}

func part2(list []string) int {

	return 0
}
