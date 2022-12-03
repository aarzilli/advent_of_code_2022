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

func conv(s string) int {
	r := 0
	for i := range s {
		ch := s[i]
		var z int
		if ch >= 'a' && ch <= 'z' {
			z = int(ch) - 'a' + 1
		} else {
			z = int(ch) - 'A' + 27
		}
		r += z
	}
	return r
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	var r string
	for _, line := range lines {
		m1 := Histo([]byte(line[len(line)/2:]))
		m2 := Histo([]byte(line[:len(line)/2]))
		m3 := Intersect(m1, m2)
		if len(m3) != 1 {
			panic("blah")
		}
		r += string(OneKey(m3))
	}
	pln(r)
	Sol(conv(r))

	r2 := ""
	for i := 0; i < len(lines); i += 3 {
		ma := Histo([]byte(lines[i+0]))
		mb := Histo([]byte(lines[i+1]))
		mc := Histo([]byte(lines[i+2]))
		m3 := Intersect(Intersect(ma, mb), mc)
		if len(m3) != 1 {
			panic("blah")
		}
		r2 += string(OneKey(m3))
	}
	pln(r2)
	Sol(conv(r2))
}
