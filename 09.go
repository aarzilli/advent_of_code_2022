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

type point struct {
	i, j int
}

var head point
var tails [9]point

func dist(a, b point) int {
	return Abs(a.i-b.i) + Abs(a.j-b.j)
}

func neighbors8(a point) []point {
	r := []point{}
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}
			r = append(r, point{a.i + di, a.j + dj})
		}
	}
	return r
}

func touching(a, b point) bool {
	for _, p := range neighbors8(a) {
		if p == b {
			return true
		}
	}
	return false

}

func advance(head, tail *point) {
	if tail.i == head.i {
		if tail.j-head.j >= 2 {
			tail.j--
		} else if head.j-tail.j >= 2 {
			tail.j++
		}
	} else if tail.j == head.j {
		if tail.i-head.i >= 2 {
			tail.i--
		} else if head.i-tail.i >= 2 {
			tail.i++
		}
	} else if !touching(*head, *tail) {
		found := false
		for _, p := range []point{
			point{tail.i + 1, tail.j + 1},
			point{tail.i + 1, tail.j - 1},
			point{tail.i - 1, tail.j - 1},
			point{tail.i - 1, tail.j + 1}} {
			if touching(*head, p) {
				found = true
				*tail = p
				break
			}
		}
		if !found {
			panic("error")
		}
	}
}

var seen = map[point]bool{}
var seen2 = map[point]bool{}

func vis() {
	for i := -4; i <= 0; i++ {
	jloop:
		for j := 0; j <= 6; j++ {
			p := point{i, j}
			if p == head {
				pf("H")
				continue jloop
			}
			for i := range tails {
				if tails[i] == p {
					pf("%d", i+1)
					continue jloop
				}
			}
			pf(".")
		}
		pf("\n")
	}
	pf("\n")
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))

	for _, line := range lines {
		v := Spac(line, " ", -1)
		n := Atoi(v[1])
		pln(line)
		for k := 0; k < n; k++ {
			switch v[0] {
			case "R":
				head.j++
			case "L":
				head.j--
			case "U":
				head.i--
			case "D":
				head.i++
			default:
				panic("blah")
			}

			curhead := &head
			for i := range tails {
				advance(curhead, &tails[i])
				curhead = &tails[i]
			}

			seen[tails[0]] = true
			seen2[tails[8]] = true

			vis()
		}
	}

	Sol(len(seen))
	Sol(len(seen2))
}
