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

const part1 = false

func main() {
	buf, err := ioutil.ReadFile(os.Args[1])
	Must(err)
	groups := strings.SplitN(string(buf), "\n\n", -1)
	pf("len %d\n", len(groups))

	var stacks [][]string

	lines := strings.SplitN(groups[0], "\n", -1)
	lines = lines[:len(lines)-1]

	for k := len(lines) - 1; k >= 0; k-- {
		line := lines[k]
		if stacks == nil {
			stacks = make([][]string, len(line)/4+1)
		}
		for i, j := 1, 0; i < len(line); i, j = i+4, j+1 {
			if line[i] != ' ' {
				stacks[j] = append(stacks[j], string(line[i]))
			}
		}
	}
	pln(stacks)

	move := func(start, end int, n int) {
		tomove := stacks[start][len(stacks[start])-n : len(stacks[start])]
		stacks[start] = stacks[start][:len(stacks[start])-n]
		stacks[end] = append(stacks[end], tomove...)
	}

	for _, line := range Noempty(Spac(groups[1], "\n", -1)) {
		v := Getints(line, false)
		n, start, end := v[0], v[1], v[2]
		start--
		end--
		if part1 {
			for i := 0; i < n; i++ {
				move(start, end, 1)
			}
		} else {
			move(start, end, n)
		}
		pln(stacks)
	}

	r := ""
	for i := range stacks {
		r += stacks[i][len(stacks[i])-1]
	}
	Sol(r)
}
