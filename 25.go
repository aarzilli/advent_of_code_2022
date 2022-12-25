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

func snafu2int(line string) int {
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

func int2snafu(tot int) string {
	carry := 0
	snafu := []byte{}
	for tot > 0 {
		switch d := (tot % 5) + carry; d {
		case 0, 1, 2:
			snafu = append(snafu, byte(d+'0'))
			carry = 0
		case 3:
			snafu = append(snafu, '=')
			carry = 1
		case 4:
			snafu = append(snafu, '-')
			carry = 1
		case 5:
			snafu = append(snafu, '0')
			carry = 1
		}
		tot /= 5
	}
	if carry == 1 {
		snafu = append(snafu, '1')
	}
	Reverse(snafu)
	return string(snafu)
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	tot := 0
	for _, line := range lines {
		r := snafu2int(line)
		pf("%s\t%d\n", line, r)
		if int2snafu(r) != line {
			panic("conv")
		}
		tot += r
	}
	pln(tot)
	w := int2snafu(tot)
	Sol(w)
}

// 13-1=0
