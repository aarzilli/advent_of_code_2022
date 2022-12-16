package main

import (
	. "aoc/util"
	"fmt"
	"os"
	_ "sort"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

var next = map[string][]string{}
var rate = map[string]int{}
var nametonum = map[string]int{}
var numtoname = map[int]string{}
var maxrate int
var shortestpath = map[string]map[string]int{}

type possib struct {
	es     [2]entity
	opened uint64
	tot    int
	time   int
}

type entity struct {
	curnode string
	tgtnode string
	dist    int
}

func (p *possib) flow(m int) {
	currate := 0
	for n := range rate {
		if p.isopen(n) {
			currate += rate[n]
		}
	}
	p.tot += m * currate
	p.time += m
}

func (p *possib) isopen(pos string) bool {
	return (p.opened>>nametonum[pos])&0x1 != 0
}

func (p *possib) setopened(pos string) {
	p.opened = p.opened | (1 << nametonum[pos])
}

func updatefordist(cur *possib) {
	for i := range cur.es {
		if cur.es[i].dist == 0 {
			cur.es[i].curnode = cur.es[i].tgtnode
			cur.setopened(cur.es[i].curnode)
		}
	}
}

func checksol(cur possib) {
	if cur.tot > bestresult {
		//pln("solution", cur)
		bestresult = cur.tot
	}
}

var seen = map[possib]bool{}
var bestresult int

func search1(cur possib) {
	const maxtime = 30
	updatefordist(&cur)

	if cur.time == maxtime {
		checksol(cur)
		return
	}
	if cur.tot+(maxtime-cur.time)*maxrate < bestresult {
		return
	}

	if seen[cur] {
		return
	}
	seen[cur] = true

	// search all possible openable nodes
	tgts := []string{}
	for n := range rate {
		if rate[n] == 0 || cur.isopen(n) {
			continue
		}
		tgts = append(tgts, n)
	}

	if len(tgts) == 0 {
		var p possib
		p = cur
		p.flow(maxtime - cur.time)
		search1(p)
		return
	}

	var ps []possib
	if cur.es[0].dist != 0 {
		panic("unexpected")
	}
	for _, tgt := range tgts {
		var p possib = cur
		p.es[0].tgtnode = tgt
		p.es[0].dist = shortestpath[p.es[0].curnode][p.es[0].tgtnode] + 1
		ps = append(ps, p)
	}

	for _, p := range ps {
		m := p.es[0].dist
		if p.time+m > maxtime {
			p.flow(maxtime - p.time)
			checksol(p)
			continue
		}
		p.flow(m)
		p.es[0].dist -= m
		search1(p)
	}
}

func search2(cur possib) {
	const maxtime = 26
	updatefordist(&cur)

	if cur.time == maxtime {
		checksol(cur)
		return
	}
	if cur.tot+(maxtime-cur.time)*maxrate < bestresult {
		return
	}

	if cur.es[0].curnode > cur.es[1].curnode {
		cur.es[0], cur.es[1] = cur.es[1], cur.es[0]
	}

	if seen[cur] {
		return
	}
	seen[cur] = true

	// search all possible openable nodes
	tgts := []string{}
	for n := range rate {
		if rate[n] == 0 || cur.isopen(n) {
			continue
		}
		tgts = append(tgts, n)
	}

	if len(tgts) == 0 {
		var p possib
		p = cur
		p.flow(maxtime - cur.time)
		search2(p)
		return
	}

	var ps []possib
	if cur.es[0].dist == 0 && cur.es[1].dist > 0 {
		for _, tgt := range tgts {
			var p possib = cur
			p.es[0].tgtnode = tgt
			p.es[0].dist = shortestpath[p.es[0].curnode][p.es[0].tgtnode] + 1
			ps = append(ps, p)
		}
	} else if cur.es[1].dist == 0 && cur.es[0].dist > 0 {
		for _, tgt := range tgts {
			var p possib = cur
			p.es[1].tgtnode = tgt
			p.es[1].dist = shortestpath[p.es[1].curnode][p.es[1].tgtnode] + 1
			ps = append(ps, p)
		}
	} else if cur.es[0].dist == 0 && cur.es[1].dist == 0 {
		for _, tgt0 := range tgts {
			for _, tgt1 := range tgts {
				if tgt0 == tgt1 {
					continue
				}
				var p possib = cur
				p.es[0].tgtnode = tgt0
				p.es[0].dist = shortestpath[p.es[0].curnode][p.es[0].tgtnode] + 1
				p.es[1].tgtnode = tgt1
				p.es[1].dist = shortestpath[p.es[1].curnode][p.es[1].tgtnode] + 1
				ps = append(ps, p)
			}
		}
	}

	for _, p := range ps {
		m := Min([]int{p.es[0].dist, p.es[1].dist})
		if p.time+m > maxtime {
			p.flow(maxtime - p.time)
			checksol(p)
			continue
		}
		p.flow(m)
		p.es[0].dist -= m
		p.es[1].dist -= m
		search2(p)
	}
}

type pathnode struct {
	node string
	n    int
}

func searchpath(start, end string) int {
	fringe := []pathnode{pathnode{start, 0}}
	seen := map[string]bool{start: true}

	for len(fringe) > 0 {
		cur := fringe[0]
		fringe = fringe[1:]

		if cur.node == end {
			return cur.n
		}

		seen[cur.node] = true

		for _, nb := range next[cur.node] {
			if seen[nb] {
				continue
			}
			fringe = append(fringe, pathnode{nb, cur.n + 1})
		}
	}
	panic("unreachable")
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	for i, line := range lines {
		r := Getints(line, false)[0]
		v := Spac(line, " ", -1)
		cur := v[1]
		nametonum[cur] = i
		numtoname[i] = cur
		rate[cur] = r
		for i := 9; i < len(v); i++ {
			n := v[i]
			if n[len(n)-1] == ',' {
				n = n[:len(n)-1]
			}
			next[cur] = append(next[cur], n)
		}
	}
	pln(rate, next)

	for _, r := range rate {
		maxrate += r
	}

	shortestpath["AA"] = make(map[string]int)
	for n := range rate {
		if rate[n] > 0 {
			shortestpath["AA"][n] = searchpath("AA", n)
		}
	}

	for n1 := range rate {
		if rate[n1] == 0 {
			continue
		}
		if shortestpath[n1] == nil {
			shortestpath[n1] = make(map[string]int)
		}
		for n2 := range rate {
			if rate[n2] == 0 {
				continue
			}
			if n1 == n2 {
				continue
			}
			shortestpath[n1][n2] = searchpath(n1, n2)
		}
	}

	start := possib{es: [2]entity{entity{curnode: "AA", tgtnode: "AA", dist: 0}, entity{curnode: "AA", tgtnode: "AA", dist: 0}}}
	for n := range rate {
		if rate[n] == 0 {
			start.setopened(n)
		}
	}

	pln("==== PART 1 ====")

	search1(start)
	Sol(bestresult)

	pln("==== PART 2 ====")

	seen = map[possib]bool{}
	bestresult = 0
	search2(start)
	Sol(bestresult)
}
