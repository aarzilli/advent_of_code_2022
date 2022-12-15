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

type sensor struct {
	x, y   int
	bx, by int
}

var S []sensor

func dist(s *sensor, x, y int) int {
	return Abs(s.x-x) + Abs(s.y-y)
}

func (s *sensor) inrange(x, y int) bool {
	return dist(s, x, y) <= dist(s, s.bx, s.by)
}

func (s *sensor) rng() int {
	return dist(s, s.bx, s.by)
}

func (s *sensor) fullcoverage(y int) (min, max int) {
	rng := s.rng() - Abs(s.y-y)
	return s.x - rng, s.x + rng
}

func filled(y, minx, maxx int) int {
	cnt := 0
	x := minx
	for {
		if x > maxx {
			break
		}

		var inrange *sensor
		for i, _ := range S {
			_ = i
			s := &S[i]
			if s.inrange(x, y) {
				inrange = s
				break
			}
		}
		if inrange != nil {
			_, b := inrange.fullcoverage(y)
			cnt += (b + 1) - x
			x = b + 1
		} else {
			x++
		}
	}
	beacons := map[int]bool{}
	for x := minx; x <= maxx; x++ {
		for _, s := range S {
			if s.by == y {
				beacons[s.bx] = true
			}
		}
	}
	return cnt - len(beacons)
}

func searchy(y, minx, maxx int) (int, bool) {
	x := minx
	for {
		if x > maxx {
			break
		}

		var inrange *sensor
		for i, _ := range S {
			_ = i
			s := &S[i]
			if s.inrange(x, y) {
				inrange = s
				break
			}
		}
		if inrange != nil {
			_, b := inrange.fullcoverage(y)
			x = b + 1
		} else {
			return x*4000000 + y, true
		}
	}
	return 0, false
}

func part2(sz int) int {
	for y := 0; y <= sz; y++ {
		sol, ok := searchy(y, 0, sz)
		if ok {
			return sol
		}
	}
	panic("not found")
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	first := true
	minx, maxx := 0, 0
	for _, line := range lines {
		v := Getints(line, true)
		S = append(S, sensor{x: v[0], y: v[1], bx: v[2], by: v[3]})
		for _, x := range []int{v[0], v[2]} {
			if x < minx || first {
				minx = x
			}
			if x > maxx {
				maxx = x
			}
			first = false
		}
	}
	maxrng := 0
	for _, s := range S {
		if s.rng() > maxrng {
			maxrng = s.rng()
		}
	}
	if len(lines) < 20 {
		Expect(26)
		Sol(filled(10, minx-maxrng, maxx+maxrng))
		Expect(56000011)
		Sol(part2(20))
	} else {
		Expect(5525990)
		Sol(filled(2000000, minx-maxrng, maxx+maxrng))
		Expect(11756174628223)
		Sol(part2(4000000))
	}
}
