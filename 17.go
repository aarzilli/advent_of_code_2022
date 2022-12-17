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

var M = make([][]byte, 100000*4)
var pos point
var kind int

type point struct {
	i, j int
}

func render(ch byte) {
	for _, d := range offsets() {
		M[pos.i+d.i][pos.j+d.j] = ch
	}
}

func highest() int {
	for i := len(M) - 1; i > 0; i-- {
		found := false
		for j := range M[i] {
			if M[i][j] != 0 {
				found = true
				break
			}
		}
		if !found {
			return i + 1
		}
	}
	panic("full")
}

func offsets() []point {
	switch kind {
	case 0:
		return []point{{0, 0}, {0, 1}, {0, 2}, {0, 3}}
	case 1:
		return []point{{0, 1}, {-1, 1}, {-1, 0}, {-2, 1}, {-1, 2}}
	case 2:
		return []point{{0, 0}, {0, 1}, {0, 2}, {-1, 2}, {-2, 2}}
	case 3:
		return []point{{0, 0}, {-1, 0}, {-2, 0}, {-3, 0}}
	case 4:
		return []point{{0, 0}, {0, 1}, {-1, 0}, {-1, 1}}
	default:
		panic("blah")
	}
}

func colliding() bool {
	for _, d := range offsets() {
		i, j := pos.i+d.i, pos.j+d.j
		if i < 0 || i >= len(M) {
			return true
		}
		if j < 0 || j >= len(M[i]) {
			return true
		}
		if M[i][j] != 0 {
			return true
		}
	}
	return false

}

func show() {
	render('@')
	for i := highest() - 4; i < len(M); i++ {
		for j := 0; j < len(M[i]); j++ {
			if M[i][j] == 0 {
				pf(".")
			} else {
				pf("%c", M[i][j])
			}
		}
		pln()
	}
	render(0)
	pln()
}

func showlimit() {
	h := highest()
	for i := h - 4; i < Min([]int{h + 30, len(M)}); i++ {
		for j := 0; j < len(M[i]); j++ {
			if M[i][j] == 0 {
				pf(".")
			} else {
				pf("%c", M[i][j])
			}
		}
		pln()
	}
	pln()
}

func run(ch byte) {
	saved := pos
	switch ch {
	case '<':
		pos.j--
	case '>':
		pos.j++
	default:
		panic("blah")
	}
	if colliding() {
		pos = saved
	}
}

func drop() bool {
	saved := pos
	pos.i++
	if colliding() {
		pos = saved
		return false
	}
	return true
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	input := []byte(lines[0])

	for i := range M {
		M[i] = make([]byte, 7)
	}

	start := highest() - 4
	pos = point{start, 2}
	k := 0
	n := 0
	t := 0

	pln(len(input))

	target := 1000000000000

	var initialOffset int
	var increment int
	var initialOffsetN int
	var incrementN int
	var hoff = 0

	// does not work on example

	for {
		if k == 0 {
			switch t / len(input) {
			case 1:
				initialOffset = len(M) - highest()
				initialOffsetN = n
			case 2:
				increment = (len(M) - highest()) - initialOffset
				incrementN = n - initialOffsetN
				pf("h = %d + %d, n = %d + %d\n", initialOffset, increment, initialOffsetN, incrementN)

				c := (target - initialOffsetN) / incrementN
				n2 := c*incrementN + initialOffsetN
				pf("during iteration %d n = %d\n", c, n2)
				n = n2
				hoff = increment * (c - 1)

			default:
				// nothing
			}
		}

		run(input[k])
		k = (k + 1) % len(input)
		t++

		if !drop() {
			render('#')
			start = highest() - 4
			pos = point{start, 2}
			kind = (kind + 1) % 5
			n++
			h := len(M) - highest()
			if n == 2022 {
				Sol(h + hoff)
			}
			if n == target {
				Sol(h + hoff)
				break
			}
		}
	}
}
