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

const (
	RIGHT = iota
	DOWN
	LEFT
	UP
)

var dirstr = map[int]string{RIGHT: "right", DOWN: "down", LEFT: "left", UP: "up"}

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

type face struct {
	name           int
	starti, startj int
	M              [][]byte
	edges          [4]edge
}

type edge struct {
	toface *face
	toedge int
}

func (e *edge) empty() bool {
	return e.toface == nil
}

func debugconnectivity(faces []*face) {
	for i := range faces {
		fmt.Printf("face %d\n", faces[i].name)
		for j, edge := range faces[i].edges {
			if edge.toface != nil {
				fmt.Printf("\t%s -> %d,%s\n", dirstr[j], edge.toface.name, dirstr[edge.toedge])
			}
		}
	}
}

func mod4(n int) int {
	return ((n % 4) + 4) % 4
}

func splitfaces(M [][]byte) []*face {
	facesz := 50
	if len(M) < 50 {
		facesz = 4
	}

	faces := []*face{}

	for starti := 0; starti < len(M); starti += facesz {
		for startj := 0; startj < len(M[starti]); startj += facesz {
			face := &face{starti: starti, startj: startj}
			face.M = make([][]byte, facesz)
			for i := 0; i < facesz; i++ {
				face.M[i] = make([]byte, facesz)
				for j := 0; j < facesz; j++ {
					face.M[i][j] = M[starti+i][startj+j]
				}
			}
			if !face.isempty() {
				face.name = len(faces) + 1
				faces = append(faces, face)
			}
		}
	}

	if len(faces) != 6 {
		panic("bad number of faces")
	}

	getface := func(starti, startj int) *face {
		for _, f := range faces {
			if f.starti == starti && f.startj == startj {
				return f
			}
		}
		return nil
	}

	// connections on foldout
	for _, f := range faces {
		f.edges[RIGHT] = edge{getface(f.starti, f.startj+facesz), LEFT}
		f.edges[LEFT] = edge{getface(f.starti, f.startj-facesz), RIGHT}
		f.edges[UP] = edge{getface(f.starti-facesz, f.startj), DOWN}
		f.edges[DOWN] = edge{getface(f.starti+facesz, f.startj), UP}
	}

	// angles
	for k := 0; k < 2; k++ {
		for _, f := range faces {
			for dir := RIGHT; dir <= UP; dir++ {
				a := f.edges[dir].toface
				aedge := f.edges[dir].toedge
				b := f.edges[mod4(dir+1)].toface
				bedge := f.edges[mod4(dir+1)].toedge
				if a != nil && b != nil {
					a.edges[mod4(aedge-1)] = edge{b, mod4(bedge + 1)}
					b.edges[mod4(bedge+1)] = edge{a, mod4(aedge - 1)}
				}
			}
		}
	}

	//debugconnectivity(faces)
	return faces
}

func (f *face) isempty() bool {
	for i := range f.M {
		for j := range f.M[i] {
			if f.M[i][j] != ' ' {
				return false
			}
		}
	}
	return true
}

func mapcoord(pos point, curedge int, nextface *face, nextedge int) point {
	var good int
	switch curedge {
	case RIGHT, LEFT:
		good = pos.i
	case UP, DOWN:
		good = pos.j
	}
	sz := len(nextface.M)
	var next point
	var p *int
	switch nextedge {
	case UP:
		p = &next.j
		next.i = 0
	case DOWN:
		p = &next.j
		next.i = sz - 1
	case LEFT:
		p = &next.i
		next.j = 0
	case RIGHT:
		p = &next.i
		next.j = sz - 1
	}

	if nextedge == curedge {
		*p = sz - good - 1
	} else if mod4(nextedge+2) == curedge {
		*p = good
	} else {
		switch curedge {
		case RIGHT, LEFT:
			switch nextedge {
			case mod4(curedge + 3):
				*p = sz - good - 1
			case mod4(curedge + 1):
				*p = good
			}

		case UP, DOWN:
			switch nextedge {
			case mod4(curedge + 3):
				*p = good
			case mod4(curedge + 1):
				*p = sz - good - 1
			}
		}
	}

	return next
}

func advance2(curface *face, pos point, dir int) (*face, point, int) {
	nextface := curface
	next := pos
	nextdir := dir

	switch dir {
	case RIGHT:
		next.j++
	case DOWN:
		next.i++
	case LEFT:
		next.j--
	case UP:
		next.i--
	default:
		panic("blah")
	}

	cross := func(edge int) {
		nextface = curface.edges[edge].toface
		next = mapcoord(next, edge, nextface, curface.edges[edge].toedge)
		nextdir = mod4(curface.edges[edge].toedge + 2)
	}

	if next.i < 0 {
		cross(UP)
	}
	if next.i >= len(curface.M) {
		cross(DOWN)
	}
	if next.j < 0 {
		cross(LEFT)
	}
	if next.j >= len(curface.M[next.i]) {
		cross(RIGHT)
	}

	return nextface, next, nextdir
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

	{
		// PART 1
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
	}

	{
		// PART 2
		faces := splitfaces(M)

		curface := faces[0]
		pos := point{0, 0}
		dir := RIGHT
		it := instrIter{pathstr, 0, 0}

		for it.next() {
			for k := 0; k < it.steps; k++ {
				//pln(curface.name, curface.starti+pos.i, curface.startj+pos.j, dirstr[dir])
				nextface, next, nextdir := advance2(curface, pos, dir)

				if nextface.M[next.i][next.j] == '#' {
					break
				}
				if nextface.M[next.i][next.j] != '.' {
					panic("blah")
				}

				curface = nextface
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
		Sol(1000*(curface.starti+pos.i+1) + 4*(curface.startj+pos.j+1) + dir)
	}
}
