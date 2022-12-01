package main

import (
	. "aoc/util"
	"fmt"
	"os"
	"sort"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

func main() {
	groups := Input(os.Args[1], "\n\n", true)
	elves := []int{}
	for _, group := range groups {
		tot := 0
		for _, n := range Vatoi(Spac(group, "\n", -1)) {
			tot += n
		}
		elves = append(elves, tot)
	}
	sort.Ints(elves)
	Sol(elves[len(elves)-1])
	Sol(elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3])
}
