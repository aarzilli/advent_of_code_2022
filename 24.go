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

var M [][]byte

type point struct {
	i, j int
}

type state struct {
	e    point
	time int
}

type blizz struct {
	point
	dir byte
}

var B []blizz

type occukey struct {
	i, j, t int
}

var occumemo = map[occukey]byte{}

func occupied(i, j, t int) byte {
	if M[i][j] == '#' {
		return '#'
	}
	if ch, ok := occumemo[occukey{i, j, t}]; ok {
		return ch
	}

	for _, b := range B {
		bi, bj := b.posat(t)
		if bi == i && bj == j {
			occumemo[occukey{i, j, t}] = b.dir
			return b.dir
		}
	}

	occumemo[occukey{i, j, t}] = '.'
	return '.'
}

func (b *blizz) posat(t int) (int, int) {
	normi := func(deltai int) int {
		L := len(M) - 2
		return (((((b.i - 1) + deltai) % L) + L) % L) + 1
	}

	normj := func(deltaj int) int {
		L := len(M[0]) - 2
		return (((((b.j - 1) + deltaj) % L) + L) % L) + 1
	}

	switch b.dir {
	case '<':
		return b.i, normj(-t)
	case '>':
		return b.i, normj(+t)
	case 'v':
		return normi(+t), b.j
	case '^':
		return normi(-t), b.j
	default:
		panic("blah")
	}
}

func show(cur state) {
	for i := range M {
		for j := range M[i] {
			if cur.e.i == i && cur.e.j == j {
				pf("E")
			} else if M[i][j] == '#' {
				pf("#")
			} else {
				ch := occupied(i, j, cur.time)
				pf("%c", ch)
			}
		}
		pln()
	}
}

func search(start state, end point) state {
	djk := NewDijkstra[state](start)
	cnt := 0
	var cur state
	for djk.PopTo(&cur) {
		if cnt%1000 == 0 {
			pf("fringe %d (min dist %d)\n", djk.Len(), djk.Dist[cur])
		}
		cnt++
		if cur.e == end {
			pf("found %v\n", cur)
			//show(cur)
			return cur
		}

		maybeadd := func(nb state) {
			if nb.e.i < 0 || nb.e.i >= len(M) || nb.e.j < 0 || nb.e.j >= len(M[nb.e.i]) {
				return
			}
			if M[nb.e.i][nb.e.j] == '#' {
				return
			}
			ch := occupied(nb.e.i, nb.e.j, nb.time)
			if ch == '.' {
				djk.Add(cur, nb, 1)
			}
		}

		maybeadd(state{e: cur.e, time: cur.time + 1})
		maybeadd(state{e: point{cur.e.i - 1, cur.e.j}, time: cur.time + 1})
		maybeadd(state{e: point{cur.e.i + 1, cur.e.j}, time: cur.time + 1})
		maybeadd(state{e: point{cur.e.i, cur.e.j + 1}, time: cur.time + 1})
		maybeadd(state{e: point{cur.e.i, cur.e.j - 1}, time: cur.time + 1})
	}
	panic("not found")
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	M = make([][]byte, len(lines))
	for i := range lines {
		M[i] = []byte(lines[i])
		for j := range M[i] {
			if M[i][j] == '>' || M[i][j] == '<' || M[i][j] == 'v' || M[i][j] == '^' {
				B = append(B, blizz{point: point{i: i, j: j}, dir: M[i][j]})
			}
		}
	}

	startpos := point{0, 1}
	endpos := point{len(M) - 1, len(M[0]) - 2}
	p0 := state{e: startpos, time: 0}
	p1 := search(p0, endpos)
	Sol(p1.time)
	p2 := search(p1, startpos)
	p3 := search(p2, endpos)
	Sol(p3.time)
}

// 182
