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

type node struct {
	name string
	args [2]string
	op   string
	val  int
}

var nodes = map[string]*node{}

func (n *node) value() (r int) {
	if n.op == "" {
		return n.val
	}
	switch n.op {
	case "+":
		return nodes[n.args[0]].value() + nodes[n.args[1]].value()
	case "-":
		return nodes[n.args[0]].value() - nodes[n.args[1]].value()

	case "/":
		return nodes[n.args[0]].value() / nodes[n.args[1]].value()
	case "*":
		return nodes[n.args[0]].value() * nodes[n.args[1]].value()
	case "=":
		if nodes[n.args[0]].value() == nodes[n.args[1]].value() {
			return 1
		} else {
			return 0
		}
	}
	panic("blah")
}

func (n *node) simplify() (int, string) {
	if n.name == "humn" {
		return 0, "h"
	}
	if n.op == "" {
		return n.val, ""
	}

	a, astr := nodes[n.args[0]].simplify()
	b, bstr := nodes[n.args[1]].simplify()

	if astr == "" && bstr == "" {
		switch n.op {
		case "+":
			return a + b, ""
		case "-":
			return a - b, ""
		case "/":
			return a / b, ""
		case "*":
			return a * b, ""
		case "=":
			if a == b {
				return 1, ""
			} else {
				return 0, ""
			}
		}
	}

	if astr == "" {
		astr = fmt.Sprintf("%d", a)
	}
	if bstr == "" {
		bstr = fmt.Sprintf("%d", b)
	}

	return 0, fmt.Sprintf("(%s %s %s)", astr, n.op, bstr)
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	for _, line := range lines {
		v := Spac(line, " ", -1)
		name := v[0][:len(v[0])-1]
		if len(v) > 2 {
			n := &node{
				name: name,
				args: [2]string{
					v[1],
					v[3],
				},
				op: v[2],
			}
			nodes[name] = n
		} else {
			n := &node{
				name: name,
				val:  Atoi(v[1]),
			}
			nodes[name] = n
		}
	}
	Sol(nodes["root"].value())

	nodes["root"].op = "="
	_, str := nodes["root"].simplify()
	pln(str)
}
