package main

import (
	. "aoc/util"
	"fmt"
	"sort"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

func main() {
	groups := Input("01.txt", "\n\n", true)
	pln(len(groups))
	max := 0
	elves := []int{}
	for _, group := range groups {
		tot := 0
		for _, n := range Vatoi(Spac(group, "\n", -1)) {
			tot += n
		}
		if tot >= max {
			max = tot
		}
		elves = append(elves, tot)
	}
	Sol(max)
	sort.Ints(elves)
	pln(elves)
	Sol(elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3])
}
