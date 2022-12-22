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

func show(M [][]byte, pos point, dir int) {
	pln("MAPSTART")
	for i := range M {
		for j := range M[i] {
			if pos.i == i && pos.j == j {
				switch dir {
				case 0:
					pf(">")
				case 1:
					pf("v")
				case 2:
					pf("<")
				case 3:
					pf("^")
				}
			} else {
				pf("%c", M[i][j])
			}
		}
		pln()
	}
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

func assertface(pos point, wantface int) {
	f, _, _ := face(pos.i, pos.j)
	if f != wantface {
		pf("want %d got %d\n", wantface, f)
		panic("wrong")
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
				assertface(next, 5)
			case 3:
				next = intoface(2, -1, facei)
				nextdir = UP
				assertface(next, 2)

			case 5:
				next = intoface(2, facesz-facei-1, -1)
				nextdir = LEFT
				assertface(next, 2)
			case 6:
				next = intoface(5, -1, facei)
				nextdir = UP
				assertface(next, 5)
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
				assertface(next, 3)
			case 5:
				next = intoface(6, facej, -1)
				nextdir = LEFT
				assertface(next, 6)
			case 6:
				next = intoface(2, 0, facej)
				nextdir = DOWN
				assertface(next, 2)
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
				assertface(next, 4)
			case 3:
				next = intoface(4, 0, facei)
				nextdir = DOWN
				assertface(next, 4)
			case 4:
				next = intoface(1, facesz-facei-1, 0)
				nextdir = RIGHT
				assertface(next, 1)
			case 6:
				next = intoface(1, 0, facei)
				nextdir = DOWN
				assertface(next, 1)
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
				assertface(next, 6)
			case 2:
				next = intoface(6, facesz-1, facej)
				nextdir = UP
				assertface(next, 6)
			case 4:
				next = intoface(3, facej, 0)
				nextdir = RIGHT
				assertface(next, 3)
			default:
				panic("blah")
			}
		}
	default:
		panic("blah")
	}
	return next, nextdir
}

func dirstr(dir int) string {
	switch dir {
	case 0:
		return "RIGHT"
	case 1:
		return "DOWN"
	case 2:
		return "LEFT"
	case 3:
		return "UP"
	default:
		panic("blah")
	}
}

func main() {
	/*
		lines := Input(os.Args[1], "\n", true)
		pf("len %d\n", len(lines))*/
	buf, err := ioutil.ReadFile(os.Args[1])
	Must(err)
	lines := strings.SplitN(string(buf), "\n", -1)
	M := make([][]byte, len(lines))
	for i, line := range lines {
		M[i] = []byte(line)
	}

	M = M[:len(M)-1]
	pathstr := string(M[len(M)-1])
	startpathstr := pathstr
	//pln(pathstr)
	M = M[:len(M)-2]

	pos := startj(M, 0)
	dir := 0

	//pln(pos)

	for len(pathstr) > 0 {
		end := 0
		for end < len(pathstr) {
			if pathstr[end] < '0' || pathstr[end] > '9' {
				break
			}
			end++
		}

		n := Atoi(pathstr[:end])
		pathstr = pathstr[end:]
		//pln(n)

		for k := 0; k < n; k++ {
			next := advance1(M, dir, pos)

			if M[next.i][next.j] == '#' {
				break
			}
			if M[next.i][next.j] != '.' {
				panic("blah")
			}

			pos = next

			//show(M, pos, dir)
		}

		//pln(pos)

		if len(pathstr) > 0 {
			//pln(string(pathstr[0]))
			switch pathstr[0] {
			case 'L':
				dir = (dir - 1)
				if dir == -1 {
					dir = 3
				}
			case 'R':
				dir = (dir + 1) % 4
			default:
				panic("blah")
			}
			pathstr = pathstr[1:]
		}
	}

	//Sol(1000 * (pos.i+1) + 4 * (pos.j+1) + dir)

	if len(M) < 50 {
		os.Exit(1)
	}

	/*for i := range M {
		for j := range M[i] {
			if M[i][j] == ' ' {
				pf(" ")
				continue
			}
			f, _, _ := face(i, j)
			pf("%d", f)
		}
		pln()
	}*/

	pos = startj(M, 0)
	dir = 0
	pathstr = startpathstr

	for len(pathstr) > 0 {
		end := 0
		for end < len(pathstr) {
			if pathstr[end] < '0' || pathstr[end] > '9' {
				break
			}
			end++
		}

		n := Atoi(pathstr[:end])
		pathstr = pathstr[end:]

		for k := 0; k < n; k++ {
			face, _, _ := face(pos.i, pos.j)
			//pf("%d %d %s (%d)\n", pos.j+1, pos.i+1, dirstr(dir), face)

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

		if len(pathstr) > 0 {
			//pln(string(pathstr[0]))
			switch pathstr[0] {
			case 'L':
				dir = (dir - 1)
				if dir == -1 {
					dir = 3
				}
			case 'R':
				dir = (dir + 1) % 4
			default:
				panic("blah")
			}
			pathstr = pathstr[1:]
		}
	}

	Sol(1000*(pos.i+1) + 4*(pos.j+1) + dir)
}

// 1204
// 35594
// 124018
