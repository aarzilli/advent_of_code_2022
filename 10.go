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

func main() {
	lines := Input(os.Args[1], "\n", true)
	x := 1
	cc := 1
	ss := 0

	crt := make([][]byte, 6)
	for i := range crt {
		crt[i] = make([]byte, 40)
	}

	calc := func(cc int) {
		if (cc % 40) == 20 {
			pf("ciclo %d x=%d\n", cc, x)
			ss += cc * x
		}
		crth := (cc - 1) % 40
		crtv := (cc - 1) / 40
		if crth == x-1 || crth == x || crth == x+1 {
			crt[crtv][crth] = '#'
		} else {
			crt[crtv][crth] = '.'
		}
		//pf("h=%d v=%d x=%d\n", crth, crtv, x)
	}

	for _, line := range lines {
		v := Spac(line, " ", -1)
		//pf("%s\t[cc=%d x=%d]\n", line, cc, x)
		calc(cc)
		switch v[0] {
		case "addx":
			calc(cc + 1)
			x += Atoi(v[1])
			cc += 2
		case "noop":
			cc += 1
		default:
			panic("blah")
		}
	}
	Sol(ss)

	for i := range crt {
		pf("%s\n", string(crt[i]))
	}
}
