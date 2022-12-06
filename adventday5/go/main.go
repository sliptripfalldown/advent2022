package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("../input.txt")
	buckets, commands := separateItems(string(data))

	part2(buckets, commands)
}

func part2(buckets, commands string) {

	letterSlice := parse(buckets)

	process(commands, letterSlice)

}

func process(commandList string, letterSlice [9][]byte) {

	commands := strings.Split(commandList, "\n")

	for _, letter := range letterSlice {
		fmt.Printf("LetterSlice: %#v\n", string(letter))
	}

	n := regexp.MustCompile(`(\d+)|(\d+)|(\d+)`)
	for _, command := range commands {

		values := n.FindAllString(command, -1)

		var total int
		total, _ = strconv.Atoi(values[0])
		from, _ := strconv.Atoi(values[1])
		to, _ := strconv.Atoi(values[2])

		// compensate for 0=>1 on the list
		from--
		to--

		if total > len(letterSlice[from]) {
			total = len(letterSlice[from])
			fmt.Println("Total too big")
		}

		valuesToMove := letterSlice[from][:total]

		if len(valuesToMove) == 0 {
			continue
		}

		res := bytes.Replace(letterSlice[from], valuesToMove, []byte{}, 1)
		letterSlice[from] = res

		newSlice := make([]byte, len(letterSlice[to])+total)
		copy(newSlice, valuesToMove)
		copy(newSlice[len(valuesToMove):], letterSlice[to])

		letterSlice[to] = newSlice

	}

	for i := range letterSlice {
		fmt.Printf("part 2 letters: %#v %#v\n", string(letterSlice[i][0]), i)
	}

}

func parse(list string) [9][]byte {
	lines := strings.Split(list, "\n")

	var rebuiltList [9][]byte
	for index := 0; index < 8; index++ {
		for charNum, item := range lines[index] {
			switch charNum {
			case 0, 2, 3, 4, 6, 7, 8, 10, 11, 12, 14, 15, 16, 18, 19, 20, 22, 23, 24, 26, 27, 28, 30, 31, 32, 34, 35:
				continue
			case 1:
				if item != ' ' {
					rebuiltList[0] = append(rebuiltList[0], byte(item))
				}
			case 5:
				if item != ' ' {
					rebuiltList[1] = append(rebuiltList[1], byte(item))
				}
			case 9:
				if item != ' ' {
					rebuiltList[2] = append(rebuiltList[2], byte(item))
				}
			case 13:
				if item != ' ' {
					rebuiltList[3] = append(rebuiltList[3], byte(item))
				}
			case 17:
				if item != ' ' {
					rebuiltList[4] = append(rebuiltList[4], byte(item))
				}
			case 21:
				if item != ' ' {
					rebuiltList[5] = append(rebuiltList[5], byte(item))
				}
			case 25:
				if item != ' ' {
					rebuiltList[6] = append(rebuiltList[6], byte(item))
				}
			case 29:
				if item != ' ' {
					rebuiltList[7] = append(rebuiltList[7], byte(item))
				}
			case 33:
				if item != ' ' {
					rebuiltList[8] = append(rebuiltList[8], byte(item))
				}
			}
		}
	}

	return rebuiltList
}

func separateItems(text string) (string, string) {
	// split the text into lines
	lines := strings.Split(text, "\n\n")

	return lines[0], lines[1]
}
