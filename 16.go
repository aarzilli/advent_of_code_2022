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

var next = map[string][]string{}
var rate = map[string]int{}
var nametonum = map[string]int{}
var numtoname = map[int]string{}
var maxrate int
var shortestpath = map[string]map[string]int{}

type possib struct {
	es       [2]entity
	opened   [55]bool
	bestrate int
	time     int
}

type entity struct {
	curnode string
	tgtnode string
	dist    int
	steps   int
}

func (p *possib) flow() {
	for i := range p.opened {
		if !p.opened[i] {
			continue
		}
		p.bestrate += rate[numtoname[i]]
	}
	p.time++
}

func (p *possib) isopen(pos string) bool {
	return p.opened[nametonum[pos]]
}

func (p *possib) sign() possib {
	var s possib = *p
	if s.es[0].curnode > s.es[1].curnode {
		s.es[0], s.es[1] = s.es[1], s.es[0]
	}
	return s
}

var seen = map[possib]bool{}
var bestresult int

/*func search1(cur possib, time int) {
	if seen[cur] {
		return
	}
	seen[cur] = true
	if time == 0 {
		if cur.bestrate > bestresult {
			pln(cur)
			bestresult = cur.bestrate
		}
		return
	}
	if rate[cur.node] > 0 && !cur.isopen(cur.node) {
		search1(cur.open(cur.node), time-1)
	}
	for _, n := range next[cur.node] {
		if !cur.isopen(n) && rate[n] > 0 && n != cur.node {
			search1(cur.moveto(n), time-1)
		}
	}
	for _, n := range next[cur.node] {
		if n == cur.node {
			continue
		}
		if cur.isopen(n) || rate[n] == 0 {
			search1(cur.moveto(n), time-1)
		}
	}
}*/

func updatefordist(cur *possib) {
	for i := range cur.es {
		if cur.es[i].dist == 0 {
			cur.es[i].steps = 0
			cur.es[i].curnode = cur.es[i].tgtnode
			cur.opened[nametonum[cur.es[i].curnode]] = true
		}
	}
}

/*func reschedule(cur *possib, tgts []string) {
	if cur.es[0].dist == 0 || cur.es[1].dist == 0 || cur.es[0].tgtnode != cur.es[1].tgtnode {
		return
	}

	if cur.es[0].dist > cur.es[1].dist {
		cur.es[0], cur.es[1] = cur.es[1], cur.es[0]
	}

	for _, tgt := range tgts {
		d := shortestpath[cur.es[1].curnode][tgt] + 1
		if d >= cur.es[1].steps {

		}
	}
}*/

func checksol(cur possib) {
	if cur.bestrate > bestresult {
		pln("solution", cur)
		bestresult = cur.bestrate
	}
}

func search2better(cur possib, time int) {
	//pln("search2better", cur, time)
	/*if seen[cur] {
		return
	}
	seen[cur] = true*/
	if 26-time != cur.time {
		panic("corrupted")
	}

	updatefordist(&cur)

	if time == 0 {
		checksol(cur)
		return
	}
	if cur.bestrate+time*maxrate < bestresult {
		return
	}

	/*s := cur.sign()
	if seen[s] {
		return
	}
	seen[s] = true*/

	// search all possible openable nodes
	tgts := []string{}
	for n := range rate {
		if rate[n] == 0 || cur.isopen(n) {
			continue
		}
		tgts = append(tgts, n)
	}

	if len(tgts) == 0 {
		// no openable nodes, just run the simulation to the end
		pf("no openable nodes at time %d so far %d (%d %d)\n", 26-time, cur.bestrate, cur.es[0].dist, cur.es[1].dist)
		var p possib
		p = cur
		t := time
		for t > 0 {
			p.flow()
			p.es[0].dist--
			p.es[1].dist--
			t--
		}
		pf("finished %d\n", p.bestrate)
		search2better(p, 0)
		return
	}

	var ps []possib
	if cur.es[0].dist == 0 && cur.es[1].dist > 0 {
		for _, tgt := range tgts {
			var p possib = cur
			p.es[0].tgtnode = tgt
			p.es[0].dist = shortestpath[p.es[0].curnode][p.es[0].tgtnode] + 1
			if p.es[0].dist == 1 {
				panic("bad")
			}
			//reschedule(&p, tgts)
			ps = append(ps, p)
		}
	} else if cur.es[1].dist == 0 && cur.es[0].dist > 0 {
		for _, tgt := range tgts {
			var p possib = cur
			p.es[1].tgtnode = tgt
			p.es[1].dist = shortestpath[p.es[1].curnode][p.es[1].tgtnode] + 1
			if p.es[1].dist == 1 {
				panic("bad")
			}
			//reschedule(&p, tgts)
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
				if p.es[0].dist == 0 || p.es[1].dist == 1 {
					panic("bad")
				}
				ps = append(ps, p)
			}
		}
	} else {
		pf("%d %d\n", cur.es[0].dist, cur.es[1].dist)
		panic("wtf?")
	}

	for _, p := range ps {
		if p.time+Min([]int{p.es[0].dist, p.es[1].dist}) > 26 {
			t := time
			for t > 0 {
				p.flow()
				p.es[0].dist--
				p.es[1].dist--
				t--
			}
			checksol(p)
			continue
		}
		t := time
		for {
			p.flow()
			t--
			p.es[0].dist--
			p.es[0].steps++
			p.es[1].dist--
			p.es[1].steps++
			if p.es[0].dist == 0 || p.es[1].dist == 0 {
				search2better(p, t)
				break
			}
		}
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
	pf("len %d\n", len(lines))
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

	//search(possib{ node: "AA" }, 30)
	//search2(possib{ node: "AA", node1: "AA" }, 26)
	shortestpath["AA"] = make(map[string]int)
	for n := range rate {
		if rate[n] > 0 {
			shortestpath["AA"][n] = searchpath("AA", n)
			pf("%s -> %s %d\n", "AA", n, shortestpath["AA"][n])
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
			pf("%s -> %s %d\n", n1, n2, shortestpath[n1][n2])
		}
	}

	start := possib{es: [2]entity{entity{curnode: "AA", tgtnode: "AA", dist: 0}, entity{curnode: "AA", tgtnode: "AA", dist: 0}}}
	for n := range rate {
		if rate[n] == 0 {
			start.opened[nametonum[n]] = true
		}
	}

	search2better(start, 26)
}

/*
2847 too low
2902 too low
4293 too high
2865 bad result
2871 bad result
*/
