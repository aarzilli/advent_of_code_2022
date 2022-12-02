package main

import (
	. "aoc/util"
	"fmt"
	"os"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

var wins = map[string]string{
	"paper":    "rock",
	"scissors": "paper",
	"rock":     "scissors",
}

var convert = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
}

func win(a, b string) int {
	a = convert[a]
	b = convert[b]
	if a == b {
		return 3
	}
	if wins[a] == b {
		return 0
	}
	return 6
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	scores := []int{}
	for _, line := range lines {
		fields := Spac(line, " ", -1)
		score := int(fields[1][0]-'X') + 1
		score += win(fields[0], fields[1])
		scores = append(scores, score)
	}
	Sol(Sum(scores))

	part2 := 0
	for _, line := range lines {
		fields := Spac(line, " ", -1)
		score := 0
		switch fields[1] {
		case "X":
			// lose
		case "Y":
			score += 3
		case "Z":
			score += 6
		}
		for _, b := range []string{"X", "Y", "Z"} {
			if win(fields[0], b) == score {
				score += int(b[0]-'X') + 1
				break
			}
		}
		part2 += score
	}
	Sol(part2)
}
