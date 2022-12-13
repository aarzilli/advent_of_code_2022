package main

import (
	. "aoc/util"
	"fmt"
	"os"
	"sort"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

func readlist(s string) ([]any, string) {
	r := []any{}
	if s[0] == ']' {
		return r, s[1:]
	}
	for {
		if s[0] == '[' {
			var x []any
			x, s = readlist(s[1:])
			r = append(r, x)
		} else {
			for i := 0; i < len(s); i++ {
				if s[i] < '0' || s[i] > '9' {
					r = append(r, Atoi(s[:i]))
					s = s[i:]
					break
				}
			}
		}
		if s[0] == ',' {
			s = s[1:]
			continue
		}
		if s[0] == ']' {
			return r, s[1:]
		}
		panic(fmt.Errorf("bad %q", s))
	}
}

func totree(s string) []any {
	if s[0] == '[' {
		r, s := readlist(s[1:])
		if s != "" {
			panic("blah")
		}
		return r
	}
	panic("blah")
}

func compare(a, b []any) int {
	i, j := 0, 0

	//pf("compare %v %v\n", a, b)

	for {
		if i >= len(a) {
			if j >= len(b) {
				return 0
			}
			return -1
		}
		if j >= len(b) {
			return 1
		}

		switch ax := a[i].(type) {
		case int:
			switch bx := b[i].(type) {
			case int:
				if ax < bx {
					//pln("left first")
					return -1
				}
				if ax > bx {
					//pln("right first")
					return 1
				}
			case []any:
				r := compare([]any{ax}, bx)
				if r != 0 {
					return r
				}
			}
		case []any:
			switch bx := b[i].(type) {
			case int:
				r := compare(ax, []any{bx})
				if r != 0 {
					return r
				}
			case []any:
				r := compare(ax, bx)
				if r != 0 {
					return r
				}
			}
		}

		i++
		j++
	}
}

func isdivider(a []any) int {
	if len(a) != 1 {
		return 0
	}
	ax, ok := a[0].([]any)
	if !ok {
		return 0
	}
	if len(ax) != 1 {
		return 0
	}
	axx, ok := ax[0].(int)
	if !ok {
		return 0
	}
	return axx
}

func main() {
	groups := Input(os.Args[1], "\n\n", true)
	rightorder := []int{}
	all := [][]any{}
	for i, group := range groups {
		lines := Spac(group, "\n", -1)
		a := totree(lines[0])
		b := totree(lines[1])
		r := compare(a, b)
		if r == -1 {
			rightorder = append(rightorder, i+1)
		}
		all = append(all, a, b)
	}
	pln(rightorder)
	Sol(Sum(rightorder))

	all = append(all, totree("[[2]]"), totree("[[6]]"))
	sort.Slice(all, func(i, j int) bool {
		r := compare(all[i], all[j])
		if r == -1 {
			return true
		}
		return false
	})

	dividers := 1
	for i := range all {
		n := isdivider(all[i])
		if n == 2 || n == 6 {
			pln(all[i], i+1)
			dividers *= (i + 1)
		}
	}
	Sol(dividers)
}
