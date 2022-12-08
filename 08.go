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

func visible(M [][]int, vis [][]bool, si, sj, di, dj int) {
	h := -1
	for i, j := si, sj; i >= 0 && i < len(M) && j >= 0 && j < len(M[i]); i, j = i+di, j+dj {
		if M[i][j] > h {
			h = M[i][j]
			vis[i][j] = true
		}
	}
}

func score(M [][]int, si, sj int) int {
	v := func(di, dj int) int {
		cnt := 0
		for i, j := si+di, sj+dj; i >= 0 && i < len(M) && j >= 0 && j < len(M[i]); i, j = i+di, j+dj {
			cnt++
			if M[i][j] >= M[si][sj] {
				break
			}
		}
		return cnt
	}

	// up
	upscore := v(-1, 0)

	// left
	leftscore := v(0, -1)

	// right
	rightscore := v(0, +1)

	// down
	downscore := v(+1, 0)

	return upscore * leftscore * rightscore * downscore
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	M := make([][]int, len(lines))
	vis := make([][]bool, len(M))
	for i, line := range lines {
		M[i] = Vatoi(Spac(line, "", -1))
		vis[i] = make([]bool, len(M[i]))
	}

	for i := range M {
		visible(M, vis, i, 0, 0, +1)
		visible(M, vis, i, len(M[i])-1, 0, -1)
	}

	for j := range M[0] {
		visible(M, vis, 0, j, +1, 0)
		visible(M, vis, len(M)-1, j, -1, 0)
	}

	cnt := 0
	for i := range M {
		for j := range M[i] {
			if vis[i][j] {
				cnt++
			}
		}
	}
	Sol(cnt)

	score(M, 3, 2)

	max := 0
	for i := range M {
		for j := range M[i] {
			s := score(M, i, j)
			if s > max {
				max = s
			}
		}
	}
	Sol(max)
}
