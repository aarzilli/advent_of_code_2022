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

type elf struct {
	i, j   int
	di, dj int
	move   string
}

type point struct {
	i, j int
}

var M = map[point]byte{}
var E []elf
var moves = []string{"N", "S", "W", "E"}

func tochar(ch byte) byte {
	if ch == 0 {
		return '.'
	}
	return '#'
}

func bounds() (point, point) {
	var min, max point
	for p := range M {
		if p.i < min.i {
			min.i = p.i
		}
		if p.i > max.i {
			max.i = p.i
		}
		if p.j < min.j {
			min.j = p.j
		}
		if p.j > max.j {
			max.j = p.j
		}
	}
	return min, max
}

func show() {
	min, max := bounds()
	for i := min.i; i <= max.i; i++ {
		for j := min.j; j <= max.j; j++ {
			pf("%c", tochar(M[point{i, j}]))
		}
		pln()
	}
	pln()
}

func part1() int {
	min, max := bounds()
	pln(min, max)
	return (max.i-min.i+1)*(max.j-min.j+1) - len(E)
}

func round() int {
	m := make(map[point]int)
	for i := range E {
		if !E[i].shouldmove() {
			E[i].di = E[i].i
			E[i].dj = E[i].j
			continue
		}
		E[i].pickmove()
		m[point{E[i].di, E[i].dj}]++
		//pf("elf %d (%d,%d) wants to move to %d,%d (%s)\n", i, E[i].i, E[i].j, E[i].di, E[i].dj, E[i].move)
	}

	for p := range M {
		delete(M, p)
	}

	cnt := 0
	for k := range E {
		if m[point{E[k].di, E[k].dj}] == 1 && (E[k].i != E[k].di || E[k].j != E[k].dj) {
			E[k].i = E[k].di
			E[k].j = E[k].dj
			cnt++
		}
		M[point{E[k].i, E[k].j}] = '#'
	}

	move := moves[0]
	copy(moves, moves[1:])
	moves[len(moves)-1] = move
	return cnt
}

func (e *elf) shouldmove() bool {
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}
			i, j := e.i+di, e.j+dj
			if M[point{i, j}] == '#' {
				return true
			}
		}
	}
	return false
}

func (e *elf) pickmove() {
	for _, move := range moves {
		if dst, ok := e.checkmove(move); ok {
			e.di = dst.i
			e.dj = dst.j
			e.move = move
			return
		}
	}
	e.di = e.i
	e.dj = e.j
	e.move = "-"
}

func (e *elf) checkmove(move string) (point, bool) {
	switch move {
	case "N":
		return point{e.i - 1, e.j}, allempty(point{e.i - 1, e.j}, point{e.i - 1, e.j - 1}, point{e.i - 1, e.j + 1})
	case "S":
		return point{e.i + 1, e.j}, allempty(point{e.i + 1, e.j}, point{e.i + 1, e.j - 1}, point{e.i + 1, e.j + 1})
	case "W":
		return point{e.i, e.j - 1}, allempty(point{e.i, e.j - 1}, point{e.i + 1, e.j - 1}, point{e.i - 1, e.j - 1})
	case "E":
		return point{e.i, e.j + 1}, allempty(point{e.i, e.j + 1}, point{e.i + 1, e.j + 1}, point{e.i - 1, e.j + 1})
	default:
		panic("blah")
	}
}

func allempty(ps ...point) bool {
	for _, p := range ps {
		if !empty(p) {
			return false
		}
	}
	return true
}

func empty(p point) bool {
	return M[point{p.i, p.j}] != '#'
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	M = make(map[point]byte)
	for i := range lines {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == '#' {
				E = append(E, elf{i, j, 0, 0, ""})
				M[point{i, j}] = '#'

			}
		}
	}

	var p int
	for p = 0; p < 10000000; p++ {
		pln("round", p+1)
		cnt := round()
		if cnt == 0 {
			break
		}
		if p == 9 {
			Sol(part1())
		}
		/*show()
		pln(part1())
		pln()*/
	}

	Sol(p + 1)
}

/*

.......#......
...........#..
..#.#..#......
......#.......
...#.....#..#.
.#......##....
.....##.......
..#........#..
....#.#..#....
..............
....#..#..#...
..............
*/
