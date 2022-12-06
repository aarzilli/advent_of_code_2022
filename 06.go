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

func solve(line string, n int) {
	for i := range line {
		if i+n > len(line) {
			break
		}
		if len(Histo([]byte(line[i:i+n]))) == n {
			Sol(i+n)
			break
		}
	}
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	solve(lines[0], 4)
	solve(lines[0], 14)
}
