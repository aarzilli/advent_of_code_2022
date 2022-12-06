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

var part1 = true

func main() {
	lines := Input(os.Args[1], "\n", true)
	line := lines[0]
	for i := range line {
		if part1 {
			if i >= 3 {
				m := make(map[byte]int)
				for j := 0; j < 4; j++ {
					m[line[i-j]]++
				}
				if len(m) == 4 {
					Sol(i + 1)
					part1 = false
				}
			}
		} else {
			if i >= 13 {
				m := make(map[byte]int)
				for j := 0; j < 14; j++ {
					m[line[i-j]]++
				}
				if len(m) == 14 {
					Sol(i + 1)
					break
				}
			}
		}
	}
}
