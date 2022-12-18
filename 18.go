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
	x, y, z int
}

var points = []point{}

var M = make(map[point]bool)

func neighbors(p point) []point {
	sides := []point{}
	for dx := -1; dx <= 1; dx++ {
		if dx == 0 {
			continue
		}
		p2 := p
		p2.x += dx
		sides = append(sides, p2)
	}
	for dy := -1; dy <= 1; dy++ {
		if dy == 0 {
			continue
		}
		p2 := p
		p2.y += dy
		sides = append(sides, p2)
	}
	for dz := -1; dz <= 1; dz++ {
		if dz == 0 {
			continue
		}
		p2 := p
		p2.z += dz
		sides = append(sides, p2)
	}
	return sides
}

func part1(count func(point) bool) int {
	cnt := 0
	for _, p := range points {
		for _, p2 := range neighbors(p) {
			if count(p2) {
				cnt++
			}
		}

	}
	return cnt
}

const limit = 25

func main() {
	lines := Input(os.Args[1], "\n", true)
	for _, line := range lines {
		v := Vatoi(Spac(line, ",", -1))
		points = append(points, point{x: v[0], y: v[1], z: v[2]})
		M[points[len(points)-1]] = true
	}

	Sol(part1(func(p2 point) bool {
		return !M[p2]
	}))

	djk := NewDijkstra[point](point{-limit, -limit, -limit})
	var p point
	for djk.PopTo(&p) {
		sides := neighbors(p)

		for _, p2 := range sides {
			if p2.x > limit || p2.y > limit || p2.z > limit || p2.x < -limit || p2.y < -limit || p2.z < -limit {
				continue
			}
			if M[p2] {
				continue
			}
			djk.Add(p, p2, 1)
		}
	}

	djk.Dist[point{0, 0, 0}] = 1

	Sol(part1(djk.Seen))
}
