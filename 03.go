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

func chars(s string) map[byte]bool {
	m := make(map[byte]bool)
	for i := range s {
		m[s[i]] = true
	}
	return m
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
		m1 := chars(line[len(line)/2:])
		m2 := chars(line[:len(line)/2])
		m3 := make(map[byte]int)
		for k := range m1 {
			m3[k]++
		}
		for k := range m2 {
			m3[k]++
		}
		for k := range m3 {
			if m3[k] == 2 {
				r += string(k)
				found = true
			}
		}
	}
	pln(r)
	Sol(conv(r))

	r2 := ""
	for i := 0; i < len(lines); i += 3 {
		ma := chars(lines[i+0])
		mb := chars(lines[i+1])
		mc := chars(lines[i+2])
		m3 := make(map[byte]int)
		for k := range ma {
			m3[k]++
		}
		for k := range mb {
			m3[k]++
		}
		for k := range mc {
			m3[k]++
		}
		for k := range m3 {
			if m3[k] == 3 {
				r2 += string(k)
				break
			}
		}
	}
	pln(r2)
	Sol(conv(r2))
}

// 7826

/*
WWvTzgzzgR
DdbGdLZLttl

vJrwpWtwJgWr
csFMMfFFhFp
*/
