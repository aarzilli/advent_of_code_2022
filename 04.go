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

func contains(a, b []int) bool {
	return a[0] <= b[0] && a[1] >= b[1]
}

func overlaps(a, b []int) bool {
	return contains1(a, b[0]) || contains1(a, b[1]) || contains1(b, a[0]) || contains1(b, a[1])
}

func contains1(a []int, n int) bool {
	return n >= a[0] && n <= a[1]
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	cnt := 0
	cnt2 := 0
	for _, line := range lines {
		v := Getints(line, false)
		if contains(v[:2], v[2:]) || contains(v[2:], v[:2]) {
			cnt++
		}
		if overlaps(v[:2], v[2:]) {
			cnt2++
		}
	}
	Sol(cnt)
	Sol(cnt2)
}
