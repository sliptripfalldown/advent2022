package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("../input.txt")

	part1Results := calc(data, 3)
	part2Results := calc(data, 13)

	fmt.Printf("\npart 1 results: %v\n", part1Results)
	fmt.Printf("part 2 results: %v\n", part2Results)
}

func calc(letters []byte, windowSize int) int {
	for index := 0; index < len(letters)-1; index++ {
		if index > 13 {
			window := letters[index-windowSize : index]

			hasDuplicate := hasDuplicate(window)
			if hasDuplicate {
				continue
			}

			currentLetter := letters[index]

			numMatchingLetters := 0
			for _, windowLetter := range window {
				if currentLetter == windowLetter {
					numMatchingLetters++
				}
			}

			if numMatchingLetters == 0 {
				fmt.Printf("Window: %v current lett: %v position: %v\n", string(window), string(currentLetter), index)
				return index + 1
			}
		}
	}

	return -999999999
}

func hasDuplicate(s []byte) bool {
	seenLetters := make(map[byte]bool)
	for _, b := range s {
		if seenLetters[b] {
			return true
		}
		seenLetters[b] = true
	}

	return false
}
