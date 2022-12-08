package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("../input.txt")

	part1Result := part1(string(data))
	part2Result := part2(string(data))

	fmt.Printf("\npart 1 result: %v\n", part1Result)
	fmt.Printf("part 2 result: %v\n", part2Result)

}

func part1(data string) int {
	folderSizes := parse(data)

	totalSize := 0
	for _, size := range folderSizes {
		if size < 100000 {
			totalSize += size
		}
	}

	return totalSize
}

func part2(data string) int {
	folderSizes := parse(data)

	neededSpace := 30000000 - (70000000 - folderSizes["/"])
	minSize := 0
	for _, size := range folderSizes {
		if size >= neededSpace && (minSize == 0 || size < minSize) {
			minSize = size
		}
	}

	return minSize
}

func parse(data string) map[string]int {
	folderPath := []string{}
	folderSizes := make(map[string]int)

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		commands := strings.Split(line, " ")

		if commands[0] == "$" {
			if commands[1] == "cd" {
				if commands[2] == ".." {
					folderPath = folderPath[:len(folderPath)-1]
				} else if commands[2] == "/" {
					folderPath = []string{"/"}
				} else {
					folderPath = append(folderPath, commands[2])
				}
			}
		} else {
			if commands[0] != "dir" {
				currentPath := ""
				for _, folder := range folderPath {
					if currentPath != "/" && folder != "/" {
						currentPath += "/"
					}
					currentPath += folder
					size, _ := strconv.Atoi(commands[0])
					folderSizes[currentPath] = folderSizes[currentPath] + size
				}
			}
		}
	}

	return folderSizes

}
