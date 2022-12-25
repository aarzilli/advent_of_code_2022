package main

import (
	. "aoc/util"
	"fmt"
	"os"
	"strconv"
	_ "strings"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

var conv = map[int][2]int{
	0: [2]int{0, 0},
	1: [2]int{0, 1},
	2: [2]int{0, 2},
	3: [2]int{1, -2},
	4: [2]int{1, -1},
}

func convnum(line string) int {
	r := 0
	for i := range line {
		var d int
		switch line[i] {
		case '-':
			d = -1
		case '=':
			d = -2
		default:
			d = Atoi(string(line[i]))
		}
		r *= 5
		r += d
	}
	return r
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	tot := 0
	for _, line := range lines {
		r := convnum(line)
		pf("%s\t%d\n", line, r)
		tot += r
	}
	pln(tot)

	pf("base 5 (inv):\n")
	weirdo := []int{}
	i := 0
	for tot > 0 {
		pf("%d %d\n", tot%5, conv[tot%5])
		if i < len(weirdo) {
			weirdo[i] += conv[tot%5][1]
			weirdo = append(weirdo, conv[tot%5][0])
		} else {
			weirdo = append(weirdo, conv[tot%5][1], conv[tot%5][0])

		}
		pf("-> %d\n", weirdo)
		tot /= 5
		i++
	}
	pln()

	carry := 0
	for i := range weirdo {
		weirdo[i] += carry
		carry = 0
		switch weirdo[i] {
		case 0, 1, 2, -1, -2:
			// ok
		case 3:
			carry = 1
			weirdo[i] = -2
		default:
			panic("bad")
		}
	}

	Reverse(weirdo)
	if weirdo[0] == 0 {
		weirdo = weirdo[1:]
	}
	pf("%d\n", weirdo)
	w := ""
	for i := range weirdo {
		switch weirdo[i] {
		case -2:
			w += "="
		case -1:
			w += "-"
		case 0, 1, 2:
			w += strconv.Itoa(weirdo[i])
		default:
			pf("%d\n", weirdo[i])
			panic("badnumber")
		}
	}
	pf("%q %d\n", w, convnum(w))
	//q := strings.Join(weirdo, "")
	//_ = q
	/*pf("%q %d\n", q, convnum(q))
	pf("%d\n", convnum("1-="))*/
	//pln(convnum("1=2"))
	Sol(w)
}

// 13-1=0
