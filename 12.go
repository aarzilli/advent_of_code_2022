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

// finds the closest node in the fringe, lastmin is an optimization, if we find a node that is at that distance we return it immediately (there can be nothing that's closer)
func minimum(fringe map[point]int, lastmin int) point {
	var mink point
	first := true
	for k, d := range fringe {
		if first {
			mink = k
			first = false
		}
		if d == lastmin {
			return k
		}
		if d < fringe[mink] {
			mink = k
		}
	}
	return mink
}

func conv(ch byte) byte {
	if ch == 'S' {
		return 'a'
	}
	if ch == 'E' {
		return 'z'
	}
	return ch
}

func search(start point, impossible func(ch, ch2 byte) bool, finish func(point) bool) {
	djk := NewDijkstra[point](start)

	cnt := 0

	var cur point
	for djk.PopTo(&cur) {
		if cnt%1000 == 0 {
			//pf("fringe %d (min dist %d)\n", djk.Len(), djk.Dist[cur])
		}
		cnt++

		if finish(cur) {
			//pln(cur, djk.Dist[cur], djk.PathTo(cur))
			Sol(djk.Dist[cur])
			return
		}

		maybeadd := func(nb point) {
			if nb.i < 0 || nb.i >= len(M) || nb.j < 0 || nb.j >= len(M[nb.i]) {
				return
			}
			ch := conv(M[cur.i][cur.j])
			ch2 := conv(M[nb.i][nb.j])
			if impossible(ch, ch2) {
				return
			}
			if ok, nbdist := djk.Add(cur, nb, 1); ok {
				_ = nbdist
				//pln("from", cur, "to", nb, "dist", nbdist)
			}
		}

		// try to add all possible neighbors
		maybeadd(point{cur.i - 1, cur.j})
		maybeadd(point{cur.i, cur.j - 1})
		maybeadd(point{cur.i, cur.j + 1})
		maybeadd(point{cur.i + 1, cur.j})
	}
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	M = make([][]byte, len(lines))
	for i := range lines {
		M[i] = []byte(lines[i])
	}
	var start, end point
	for i := range M {
		for j := range M[i] {
			if M[i][j] == 'S' {
				start = point{i, j}
			}
			if M[i][j] == 'E' {
				end = point{i, j}
			}
		}
	}
	pln(start, end)
	search(start,
		func(ch, ch2 byte) bool {
			return ch2 > ch+1
		},
		func(cur point) bool {
			return M[cur.i][cur.j] == 'E'
		})
	search(end,
		func(ch, ch2 byte) bool {
			return ch2 < ch-1
		},
		func(cur point) bool {
			return conv(M[cur.i][cur.j]) == 'a'
		})
}
