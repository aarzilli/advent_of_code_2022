package main

import (
	. "aoc/util"
	"fmt"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

func main() {
	lines := Input("XX.txt", "\n", true)
	pf("len %d\n", len(lines))
}
