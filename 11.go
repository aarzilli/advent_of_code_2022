package main

import (
	. "aoc/util"
	"fmt"
	"os"
	"sort"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

type titem struct {
	op      string
	arg1    *titem
	arg2    *titem
	val     int
	memomod map[int]int
}

type monkey struct {
	items           []*titem
	op              op
	divby           int
	iftrue, iffalse int
	activity        int
}

type op struct {
	arg1, op, arg2 string
}

const debugmonkey = false
const debuground = false

var monkeys = []monkey{}

func domonkey(i int) {
	if debugmonkey {
		pln("Monkey ", i)
	}
	m := &monkeys[i]
	for _, it := range m.items {
		if debugmonkey {
			pln(" Item", it)
		}
		it = m.op.do(it)
		if debugmonkey {
			pln("  New worry level", it)
		}
		dst := 0
		if it.mod(m.divby) == 0 {
			if debugmonkey {
				pln("  throw to", m.iftrue, it)
			}
			dst = m.iftrue
		} else {
			if debugmonkey {
				pln("  throw to", m.iffalse, it)
			}
			dst = m.iffalse
		}
		monkeys[dst].items = append(monkeys[dst].items, it)
		m.activity++
	}
	m.items = m.items[:0]
}

func (it *titem) mod(c int) int {
	switch it.op {
	case "":
		return it.val % c

	case "*":
		if it.memomod == nil {
			it.memomod = make(map[int]int)
		}
		if r, ok := it.memomod[c]; ok {
			return r
		}
		r := (it.arg1.mod(c) * it.arg2.mod(c)) % c
		it.memomod[c] = r
		return r

	case "+":
		if it.memomod == nil {
			it.memomod = make(map[int]int)
		}
		if r, ok := it.memomod[c]; ok {
			return r
		}
		r := (it.arg1.mod(c) + it.arg2.mod(c)) % c
		it.memomod[c] = r
		return r

	default:
		panic("blah")
	}
}

func (op *op) do(worry *titem) *titem {
	arg := func(arg string) *titem {
		if arg == "old" {
			return worry
		}
		return &titem{val: Atoi(arg)}
	}
	return &titem{op: op.op, arg1: arg(op.arg1), arg2: arg(op.arg2)}
}

func doround(n int) {
	for i := range monkeys {
		domonkey(i)
	}
	if debuground {
		pln("After round", n)
		for i := range monkeys {
			pf(" monkey %d ", i)
			for _, it := range monkeys[i].items {
				pf("%d, ", it)
			}
			pln()
		}
	}
}

func main() {
	groups := Input(os.Args[1], "\n\n", true)
	for _, group := range groups {
		var m monkey
		lines := Spac(group, "\n", -1)

		items := Vatoi(Spac(Spac(lines[1], ":", -1)[1], ",", -1))
		for _, it := range items {
			m.items = append(m.items, &titem{val: it})
		}

		opv := Spac(Spac(lines[2], ":", -1)[1], " ", -1)
		m.op.arg1 = opv[2]
		m.op.op = opv[3]
		m.op.arg2 = opv[4]

		m.divby = Getints(lines[3], false)[0]
		m.iftrue = Getints(lines[4], false)[0]
		m.iffalse = Getints(lines[5], false)[0]

		monkeys = append(monkeys, m)
	}

	for i := 1; i <= 10000; i++ {
		doround(i)
	}

	act := []int{}
	for i := range monkeys {
		act = append(act, monkeys[i].activity)
		pln(monkeys[i].activity)
	}

	sort.Ints(act)
	Sol(act[len(act)-1] * act[len(act)-2])
}

