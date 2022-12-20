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
	val  int
	prev *node
	next *node
}

func show(root *node) {
	cur := root
	for {
		pf("%d ", cur.val)
		cur = cur.next
		if cur == root {
			break
		}
	}
	pln()
}

func getnum(cur *node, n int) int {
	for i := 0; i < n; i++ {
		cur = cur.next
	}
	return cur.val
}

func pickposfwd(cur *node, n int) *node {
	ins := cur.next
	for i := 0; i < n; i++ {
		ins = ins.next
	}
	return ins
}

const part2 = true
const key = 811589153

func shuffle(v []*node) {
	for _, cur := range v {
		cur.next.prev = cur.prev
		cur.prev.next = cur.next

		if cur.val > 0 {
			ins := pickposfwd(cur, (cur.val-1)%(len(v)-1))
			cur.next = ins.next
			cur.prev = ins
			ins.next = cur
			cur.next.prev = cur
		} else {
			ins := cur.prev
			for i := 0; i < (-cur.val)%(len(v)-1); i++ {
				ins = ins.prev
			}
			cur.next = ins.next
			cur.prev = ins
			ins.next = cur
			cur.next.prev = cur
		}

		//show(v[0])
	}
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))

	var cur *node
	v := []*node{}
	for _, k := range Vatoi(lines) {
		n := &node{k, nil, nil}
		if cur == nil {
			cur = n
		} else {
			cur.next = n
			n.prev = cur
			cur = n
		}
		v = append(v, n)
	}
	cur.next = v[0]
	v[0].prev = cur

	if part2 {
		for _, cur := range v {
			cur.val *= key
		}
	}

	shuffle(v)
	if part2 {
		for i := 0; i < 9; i++ {
			shuffle(v)
		}
	}

	show(v[0])

	for _, cur := range v {
		if cur.val == 0 {
			a := getnum(cur, 1000%len(v))
			b := getnum(cur, 2000%len(v))
			c := getnum(cur, 3000%len(v))
			pln(a, b, c, a+b+c)
			Sol(a + b + c)
			break
		}
	}
}

/*


811589153 0 -2434767459 3246356612 -1623178306 2434767459 1623178306


Initial arrangement:
811589153, 1623178306, -2434767459, 2434767459, -1623178306, 0, 3246356612

After 1 round of mixing:
0, -2434767459, 3246356612, -1623178306, 2434767459, 1623178306, 811589153

After 10 rounds of mixing:
0, -2434767459, 1623178306, 3246356612, -1623178306, 2434767459, 811589153
*/
