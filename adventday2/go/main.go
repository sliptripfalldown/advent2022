package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Item struct {
	points int
	right  string
	left   string
}

/* Part 1
 X  Y  Z
{4, 8, 3} A
{1, 5, 9} B
{7, 2, 6} C

A X // Rock vs Rock = 4
A Y // Rock vs Paper = 8
A Z // Rock vs Scissors = 3

B X // Paper vs Rock = 1
B Y // Paper vs Paper = 5
B Z // Paper vs Scissors = 9

C X // Scissor vs Rock = 7
C Y // Scissor vs Paper = 2
C Z // Scissor vs Scissors = 6
*/

/* Part 2
 X  Y  Z
{3, 4, 8} A
{1, 5, 9} B
{2, 6, 7} C

A X // Rock vs Rock = 3
A Y // Rock vs Paper = 4
A Z // Rock vs Scissors = 8

B X // Paper vs Rock = 1
B Y // Paper vs Paper = 5
B Z // Paper vs Scissors = 9

C X // Scissor vs Rock = 2
C Y // Scissor vs Paper = 6
C Z // Scissor vs Scissors = 7
*/

var (
	r = &Item{
		right:  "X", // part 2 lose
		left:   "A",
		points: 1,
	}

	p = &Item{
		right:  "Y", // part 2 draw
		left:   "B",
		points: 2,
	}

	s = &Item{
		right:  "Z", // part 2 win
		left:   "C",
		points: 3,
	}
)

func main() {
	strategyGuide, _ := os.Open("../input.txt")
	defer strategyGuide.Close()

	fScan := bufio.NewScanner(strategyGuide)

	var myScore int

	for fScan.Scan() {
		chunk := strings.Split(fScan.Text(), " ")
		opponent, player := chunk[0], chunk[1]

		score := resolve(opponent, player)
		myScore += score
	}

	fmt.Printf("My score: %v\n", myScore)
}

func resolve(opponent, player string) int {
	switch opponent {
	case r.left:
		switch player {
		case r.right:
			// Tie (but need to lose per part 2)
			//return r.points + 3 // part 1
			return s.points
		case p.right:
			// Win (but needs to draw per part 2)
			//return p.points + 6 // part 1
			return r.points + 3
		case s.right:
			// Lose (but needs to win per part 2)
			//return s.points // part 1
			return p.points + 6
		}
	case p.left:
		switch player {
		case r.right:
			// Lose (works for part 2)
			return r.points
		case p.right:
			// Tie (works for part 2)
			return p.points + 3
		case s.right:
			// Win (works for part 2)
			return s.points + 6
		}
	case s.left:
		switch player {
		case r.right:
			// Win (but needs to to lose per part 2)
			//return r.points + 6 // part 1
			return p.points
		case p.right:
			// Lose (but needs to draw per part 2)
			//return p.points // part 1
			return s.points + 3
		case s.right:
			// Tie (but needs to win per part 2)
			//return s.points + 3 // part 1
			return r.points + 6
		}
	}

	return -9999999999999999
}
