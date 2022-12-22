package main

import (
	. "aoc/util"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

type point struct {
	i, j int
}

func startj(M [][]byte, i int) point {
	for j := 0; j < len(M[i]); j++ {
		if M[i][j] != ' ' {
			return point{i, j}
		}
	}
	panic("not found")
}

func endj(M [][]byte, i int) point {
	for j := len(M[i]) - 1; j >= 0; j-- {
		if M[i][j] != ' ' {
			return point{i, j}
		}
	}
	panic("not found")
}

func starti(M [][]byte, j int) point {
	for i := 0; i < len(M); i++ {
		if j <= len(M[i]) && M[i][j] != ' ' {
			return point{i, j}
		}
	}
	panic("not found")
}

func endi(M [][]byte, j int) point {
	for i := len(M) - 1; i >= 0; i-- {
		if j <= len(M[i]) && M[i][j] != ' ' {
			//pln("endi(%d) finds %d,%d\n", j, i, j)
			return point{i, j}
		}
	}
	panic("not found")
}

func advance1(M [][]byte, dir int, pos point) point {
	next := pos

	adji := func() {
		if next.i < 0 || next.i >= len(M) || next.j < 0 || next.j >= len(M[next.i]) || M[next.i][next.j] == ' ' {
			switch dir {
			case 1:
				next = starti(M, next.j)
			case 3:
				next = endi(M, next.j)
			}
		}
	}

	adjj := func() {
		if next.i < 0 || next.i >= len(M) || next.j < 0 || next.j >= len(M[next.i]) || M[next.i][next.j] == ' ' {
			switch dir {
			case 0:
				next = startj(M, next.i)
			case 2:
				next = endj(M, next.i)
			}
		}
	}

	switch dir {
	case 0:
		next.j++
		adjj()
	case 1:
		next.i++
		adji()
	case 2:
		next.j--
		adjj()
	case 3:
		next.i--
		adji()
	default:
		panic("blah")
	}

	return next
}

func face(i, j int) (face int, reali, realj int) {
	const sz = 50
	if i < sz {
		if j >= sz && j < sz*2 {
			return 1, i, j - sz
		}
		if j >= sz*2 && j < sz*3 {
			return 2, i, j - (sz * 2)
		}
		panic("unreachable")
	}
	if i >= sz && i < sz*2 {
		if j >= sz && j < sz*2 {
			return 3, i - sz, j - sz
		}
		panic("unreachable")
	}
	if i >= sz*2 && i < sz*3 {
		if j < sz {
			return 4, i - (sz * 2), j
		}
		if j >= sz && j < sz*2 {
			return 5, i - (sz * 2), j - sz
		}
		panic("unreachable")
	}
	if i >= sz*3 && i < sz*4 {
		if j < sz {
			return 6, i - (sz * 3), j
		}
		panic("unreachable")
	}
	panic("unreachable")
}

func intoface(f int, facei, facej int) point {
	const facesz = 50
	if facei < 0 {
		facei = facesz - 1
	}
	if facej < 0 {
		facej = facesz - 1
	}
	switch f {
	case 1:
		return point{facei, facej + facesz}
	case 2:
		return point{facei, facej + 2*facesz}
	case 3:
		return point{facei + facesz, facej + facesz}
	case 4:
		return point{facei + 2*facesz, facej}
	case 5:
		return point{facei + 2*facesz, facej + facesz}
	case 6:
		return point{facei + 3*facesz, facej}
	default:
		panic("blah")
	}
}

const (
	RIGHT = iota
	DOWN
	LEFT
	UP
)

func advance2(M [][]byte, pos point, dir int) (point, int) {
	const facesz = 50

	next, nextdir := pos, dir
	switch dir {
	case RIGHT:
		next.j++
		if next.j >= len(M[pos.i]) {
			face, facei, _ := face(pos.i, pos.j)
			switch face {
			case 2:
				next = intoface(5, facesz-facei-1, -1)
				nextdir = LEFT
			case 3:
				next = intoface(2, -1, facei)
				nextdir = UP

			case 5:
				next = intoface(2, facesz-facei-1, -1)
				nextdir = LEFT
			case 6:
				next = intoface(5, -1, facei)
				nextdir = UP
			default:
				panic("blah")
			}
		}
	case DOWN:
		next.i++
		if next.i >= len(M) || next.j >= len(M[next.i]) {
			face, _, facej := face(pos.i, pos.j)
			switch face {
			case 2:
				next = intoface(3, facej, -1)
				nextdir = LEFT
			case 5:
				next = intoface(6, facej, -1)
				nextdir = LEFT
			case 6:
				next = intoface(2, 0, facej)
				nextdir = DOWN
			default:
				panic("blah")
			}
		}
	case LEFT:
		next.j--
		if next.j < 0 || M[next.i][next.j] == ' ' {
			face, facei, _ := face(pos.i, pos.j)
			switch face {
			case 1:
				next = intoface(4, facesz-facei-1, 0)
				nextdir = RIGHT
			case 3:
				next = intoface(4, 0, facei)
				nextdir = DOWN
			case 4:
				next = intoface(1, facesz-facei-1, 0)
				nextdir = RIGHT
			case 6:
				next = intoface(1, 0, facei)
				nextdir = DOWN
			default:
				panic("blah")
			}

		}
	case UP:
		next.i--
		if next.i < 0 || M[next.i][next.j] == ' ' {
			face, _, facej := face(pos.i, pos.j)
			_ = facej
			switch face {
			case 1:
				next = intoface(6, facej, 0)
				nextdir = RIGHT
			case 2:
				next = intoface(6, facesz-1, facej)
				nextdir = UP
			case 4:
				next = intoface(3, facej, 0)
				nextdir = RIGHT
			default:
				panic("blah")
			}
		}
	default:
		panic("blah")
	}
	return next, nextdir
}

type instrIter struct {
	pathstr string
	steps   int
	rot     byte
}

func (it *instrIter) next() bool {
	if len(it.pathstr) <= 0 {
		return false
	}
	end := 0
	for end < len(it.pathstr) {
		if it.pathstr[end] < '0' || it.pathstr[end] > '9' {
			break
		}
		end++
	}

	it.steps = Atoi(it.pathstr[:end])
	it.pathstr = it.pathstr[end:]

	it.rot = byte(0)
	if len(it.pathstr) > 0 {
		it.rot = it.pathstr[0]
		it.pathstr = it.pathstr[1:]
	}
	return true
}

func main() {
	buf, err := ioutil.ReadFile(os.Args[1])
	Must(err)
	lines := strings.SplitN(string(buf), "\n", -1)
	M := make([][]byte, len(lines))
	for i, line := range lines {
		M[i] = []byte(line)
	}

	M = M[:len(M)-1]
	pathstr := string(M[len(M)-1])
	M = M[:len(M)-2]

	pos := startj(M, 0)
	dir := 0

	it := instrIter{pathstr, 0, 0}

	for it.next() {
		for k := 0; k < it.steps; k++ {
			next := advance1(M, dir, pos)

			if M[next.i][next.j] == '#' {
				break
			}
			if M[next.i][next.j] != '.' {
				panic("blah")
			}

			pos = next
		}

		switch it.rot {
		case 'L':
			dir = (dir - 1)
			if dir == -1 {
				dir = 3
			}
		case 'R':
			dir = (dir + 1) % 4
		case 0:
			// end
		default:
			panic("blah")
		}
	}

	Sol(1000*(pos.i+1) + 4*(pos.j+1) + dir)

	if len(M) < 50 {
		os.Exit(1)
	}

	pos = startj(M, 0)
	dir = 0
	it = instrIter{pathstr, 0, 0}

	for it.next() {
		for k := 0; k < it.steps; k++ {
			next, nextdir := advance2(M, pos, dir)

			if M[next.i][next.j] == '#' {
				break
			}
			if M[next.i][next.j] != '.' {
				panic("blah")
			}

			pos = next
			dir = nextdir
		}

		switch it.rot {
		case 'L':
			dir = (dir - 1)
			if dir == -1 {
				dir = 3
			}
		case 'R':
			dir = (dir + 1) % 4
		case 0:
			// end
		default:
			panic("blah")
		}
	}

	Sol(1000*(pos.i+1) + 4*(pos.j+1) + dir)
}
