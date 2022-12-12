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

func search(start point, cond func(ch, ch2 byte) bool, finish func(point) bool) {
	fringe := map[point]int{start: 0}   // nodes discovered but not visited (start at node 0 with distance 0)
	seen := map[point]bool{start: true} // nodes already visited (we know the minimum distance of those)

	lastmin := 0

	cnt := 0

	for len(fringe) > 0 {
		cur := minimum(fringe, lastmin)

		if cnt%1000 == 0 {
			fmt.Printf("fringe %d (min dist %d)\n", len(fringe), fringe[cur])
		}
		cnt++

		if finish(cur) {
			pln(cur, fringe[cur])
			Sol(fringe[cur])
			return
		}

		distcur := fringe[cur]
		lastmin = distcur
		delete(fringe, cur)
		seen[cur] = true

		maybeadd := func(nb point) {
			if seen[nb] {
				return
			}
			if nb.i < 0 || nb.i >= len(M) || nb.j < 0 || nb.j >= len(M[nb.i]) {
				return
			}
			ch := conv(M[cur.i][cur.j])
			ch2 := conv(M[nb.i][nb.j])
			if cond(ch, ch2) {
				return
			}
			d, ok := fringe[nb]
			if !ok || distcur+1 < d {
				fringe[nb] = distcur + 1
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
