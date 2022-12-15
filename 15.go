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

func filled(y, minx, maxx int) int {
	cnt := 0
	for x := minx; x <= maxx; x++ {
		inrange := false
		for i, s := range S {
			_ = i
			if s.inrange(x, y) {
				//pf("%d,%d covered by sensor %d (%d,%d) range %d (dist: %d)\n", x, y, i, s.x, s.y, dist(&s, s.bx, s.by), dist(&s, x, y))
				//pf("\t%d\n", s.rng() - Abs(s.y - y))

				inrange = true
				break
			}
		}
		if inrange {
			for _, s := range S {
				if s.by == y && s.bx == x {
					inrange = false
					break
				}
			}
		}
		if inrange {
			cnt++
		}
	}
	pln()
	return cnt
}

func (s *sensor) fullcoverage(y int) (min, max int) {
	rng := s.rng() - Abs(s.y-y)
	//pf("sensor %d has range %d at row y=%d covers x=%d..%d\n", 0, rng, y, s.x - rng, s.x + rng)
	return s.x - rng, s.x + rng
}

func filled2(y, minx, maxx int) int {
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
	pln("found:", cnt-len(beacons))
	return cnt
}

func searchy(y, minx, maxx int) bool {
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
			pln(x, y)
			Sol(x*4000000 + y)
			x++
			return true
		}
	}
	return false
}

func dist(s *sensor, x, y int) int {
	return Abs(s.x-x) + Abs(s.y-y)
}

func (s *sensor) inrange(x, y int) bool {
	return dist(s, x, y) <= dist(s, s.bx, s.by)
}

func (s *sensor) rng() int {
	return dist(s, s.bx, s.by)
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
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
	pln("minmax", minx-maxrng, maxx+maxrng)
	pln(S)
	if len(lines) < 20 {
		Sol(filled(10, minx-maxrng, maxx+maxrng))
		filled2(10, minx-maxrng, maxx+maxrng)
		for y := 0; y <= 20; y++ {
			if searchy(y, 0, 20) {
				break
			}
		}
	} else {
		Sol(filled(2000000, minx-maxrng, maxx+maxrng))
		filled2(2000000, minx-maxrng, maxx+maxrng)
		for y := 0; y <= 4000000; y++ {
			if searchy(y, 0, 4000000) {
				break
			}
		}
	}

}

// 3985370
// 3985369
