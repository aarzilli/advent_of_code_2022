package main

import (
	. "aoc/util"
	"fmt"
	"math/rand"
	"os"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

const (
	oreCostIndex = iota
	clayCostIndex
	obsidianCostIndex
	geodeCostIndex
)

type blueprint [4]robot
type robot struct {
	output int
	cost   [3]int
}

func convertMaterial(s string) int {
	switch s {
	case "ore":
		return oreCostIndex
	case "clay":
		return clayCostIndex
	case "obsidian":
		return obsidianCostIndex
	case "geode":
		return geodeCostIndex
	default:
		panic("blah")
	}
}

func (r *robot) canbuild(s state) bool {
	for j := range r.cost {
		if r.cost[j] > s.materials[j] {
			return false
		}
	}
	return true
}

func (r *robot) build(s *state) {
	for j := range r.cost {
		s.materials[j] -= r.cost[j]
	}
}

type state struct {
	materials [4]int
	robots    [4]int
	time      int
}

func incmaterials(s *state) {
	for i := range s.robots {
		s.materials[i] += s.robots[i]
	}
}

func maxgeodes(numgeodes int, remtime int) int {
	L := numgeodes
	R := numgeodes + remtime
	return ((L + R) * (R - L + 1)) / 2
}

func shuffle(robots []robot) []robot {
	out := make([]robot, len(robots))
	copy(out, robots)
	rand.Shuffle(len(out), func(i, j int) {
		out[i], out[j] = out[j], out[i]
	})
	return out
}

func search(bp *blueprint, maxtime int) int {
	var bestgeodes = 0
	seen := map[state]bool{}

	var search func(s state)
	search = func(s state) {
		if s.time == maxtime {
			//pln(s)
			if s.materials[geodeCostIndex] > bestgeodes {
				pln("geodes", s.materials[geodeCostIndex])
				bestgeodes = s.materials[geodeCostIndex]
			}
			return
		}

		M := maxgeodes(s.robots[geodeCostIndex], maxtime-s.time-1)

		if s.materials[geodeCostIndex]+M <= bestgeodes {
			return
		}

		if seen[s] {
			return
		}
		seen[s] = true

		bc := 0

		for _, robot := range shuffle(bp[:]) {
			if robot.canbuild(s) {
				s2 := s
				robot.build(&s2)
				incmaterials(&s2)
				s2.robots[robot.output]++
				s2.time++
				search(s2)
				if robot.output == geodeCostIndex {
					break
				}
				bc++
			}
			if s.materials[geodeCostIndex]+M <= bestgeodes {
				return
			}
		}

		if bc >= 4 {
			return
		}

		s2 := s
		incmaterials(&s2)
		s2.time++
		search(s2)
	}
	search(state{robots: [4]int{1, 0, 0, 0}})
	return bestgeodes
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	blueprints := []blueprint{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		var bp blueprint
		line = Spac(line, ":", -1)[1]
		robotstrs := Spac(line, ".", -1)
		for i, robotstr := range robotstrs {
			if robotstr == "" {
				continue
			}
			v := Spac(robotstr, " ", -1)
			bp[i].output = convertMaterial(v[1])
			bp[i].cost[convertMaterial(v[5])] = Atoi(v[4])
			if len(v) > 7 {
				bp[i].cost[convertMaterial(v[8])] = Atoi(v[7])
			}
		}
		blueprints = append(blueprints, bp)
	}
	pln(blueprints)

	tot := 0
	for i, bp := range blueprints {
		b := search(&bp, 24)
		pln("blueprint", i+1, b)
		tot += b * (i + 1)
	}
	Sol(tot)

	a := search(&blueprints[0], 32)
	pln("first", a)
	b := search(&blueprints[1], 32)
	pln("second", b)
	c := search(&blueprints[2], 32)
	pln("third", c)
	Sol(a * b * c)
}
