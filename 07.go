package main

import (
	. "aoc/util"
	"fmt"
	"os"
	"strings"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

type entry struct {
	name   string
	size   int
	child  map[string]*entry
	parent *entry
	isdir  bool
}

var root = &entry{name: "/", isdir: true, child: make(map[string]*entry)}

func (e *entry) calcDirSize() {
	if !e.isdir {
		return
	}
	for _, child := range e.child {
		child.calcDirSize()
	}
	for _, child := range e.child {
		e.size += child.size
	}
	//pln("dir", e.name, "size:", e.size)
}

func (e *entry) part1() int {
	if !e.isdir {
		return 0
	}
	r := 0
	for _, child := range e.child {
		r += child.part1()
	}
	if e.size <= 100000 {
		//pln(e.name, e.size)
		r += e.size
	}
	return r
}

var part2cur int = 1000000000000000000

func (e *entry) part2(tofree int) {
	if !e.isdir {
		return
	}
	for _, child := range e.child {
		child.part2(tofree)
	}
	if e.size >= tofree {
		if e.size < part2cur {
			pln(e.name, e.size)
			part2cur = e.size
		}
	}
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	cwd := root
	listing := false
	for _, line := range lines {
		if strings.HasPrefix(line, "$ cd ") {
			listing = false
			arg := line[5:]
			switch arg {
			case "/":
				cwd = root
			case "..":
				cwd = cwd.parent
			default:
				if cwd.child[arg] == nil {
					panic("blah")
				}
				cwd = cwd.child[arg]
			}
		} else if line == "$ ls" {
			listing = true
		} else {
			if !listing {
				panic("error")
			}
			v := Spac(line, " ", 2)
			if v[0] == "dir" {
				cwd.child[v[1]] = &entry{name: v[1], isdir: true, child: make(map[string]*entry), parent: cwd}
			} else {
				sz := Atoi(v[0])
				cwd.child[v[1]] = &entry{name: v[1], size: sz}
			}
		}
	}
	root.calcDirSize()
	Sol(root.part1())

	freespace := 70000000 - root.size
	tofree := 30000000 - freespace
	root.part2(tofree)
	Sol(part2cur)
}
