package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type item struct{ calorie int }
type elf struct {
	inventory     []item
	totalCalories int
}
type Elves []elf

func main() {
	ledger, _ := os.ReadFile("../input.txt")
	elfLedger := strings.Split(string(ledger), "\n\n")

	var best, second, third elf
	var elves Elves
	for _, elfInventory := range elfLedger {
		elf := &elf{}
		eInventory := strings.Split(elfInventory, "\n")
		for _, eItem := range eInventory {
			inventoryItem, _ := strconv.Atoi(eItem)
			elf.totalCalories += inventoryItem
			elf.inventory = append(elf.inventory, item{calorie: inventoryItem})
		}

		elves = append(elves, *elf)
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[j].totalCalories < elves[i].totalCalories
	})

	for _, elf := range elves {
		switch {
		case elf.totalCalories > best.totalCalories:
			best = elf
		case elf.totalCalories < best.totalCalories && elf.totalCalories > second.totalCalories:
			second = elf
		case elf.totalCalories < second.totalCalories && elf.totalCalories > third.totalCalories:
			third = elf
		}
	}

	fmt.Printf("best: %v\n", best.totalCalories)
	fmt.Printf("second: %v\n", second.totalCalories)
	fmt.Printf("third: %v\n", third.totalCalories)
	fmt.Printf("Top 3: %v", best.totalCalories+second.totalCalories+third.totalCalories)
}
