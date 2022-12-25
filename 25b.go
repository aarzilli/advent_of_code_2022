package main

import (
	. "aoc/util"
	"fmt"
	"os"
	_ "strings"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

func snafuadd(acc []int8, b []byte) {
	carry := int8(0)
	for i := range acc {
		var d int8
		if i < len(b) {
			switch b[i] {
			case '0', '1', '2':
				d = int8(b[i] - '0')
			case '-':
				d = -1
			case '=':
				d = -2
			}
		}
		acc[i] = acc[i] + d + carry
		carry = 0
		if acc[i] < -2 {
			acc[i] = 5 + acc[i]
			carry = -1
		} else if acc[i] > 2 {
			acc[i] = -5 + acc[i]
			carry = 1
		}
	}
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	tot := make([]int8, 20)
	for i := range tot {
		tot[i] = 0
	}
	for _, line := range lines {
		b := []byte(line)
		Reverse(b)
		snafuadd(tot, b)
	}
	Reverse(tot)
	for tot[0] == 0 && len(tot) > 1 {
		tot = tot[1:]
	}
	w := ""
	for i := range tot {
		switch tot[i] {
		case 0, 1, 2:
			w += string(tot[i] + '0')
		case -1:
			w += "-"
		case -2:
			w += "="
		}
	}
	Sol(w)
}
