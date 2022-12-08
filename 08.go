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

func visible(v []int, vis []bool, s, d int) {
	h := -1
	for i := s; i >= 0 && i < len(v); i += d {
		if v[i] > h {
			h = v[i]
			vis[i] = true
		}
	}
}

func visibleCol(M [][]int, i int, vis [][]bool, s, d int) {
	h := -1
	for j := s; j >= 0 && j < len(M); j += d {
		if M[j][i] > h {
			h = M[j][i]
			vis[j][i] = true
		}
	}
}

func score(M [][]int, si, sj int) int {
	// up
	upscore := si
	for i := si - 1; i >= 0; i-- {
		if M[i][sj] >= M[si][sj] {
			upscore = si - i
			break
		}
	}

	// left
	leftscore := sj
	for j := sj - 1; j >= 0; j-- {
		if M[si][j] >= M[si][sj] {
			leftscore = sj - j
			break
		}
	}

	// right
	rightscore := len(M[si]) - sj - 1
	for j := sj + 1; j < len(M[si]); j++ {
		if M[si][j] >= M[si][sj] {
			rightscore = j - sj
			break
		}
	}

	// down
	downscore := len(M) - si - 1
	for i := si + 1; i < len(M); i++ {
		if M[i][sj] >= M[si][sj] {
			downscore = i - si
			break
		}
	}

	//pln(upscore, leftscore, rightscore, downscore)
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
	pln(M)

	for i := range M {
		visible(M[i], vis[i], 0, +1)
		visible(M[i], vis[i], len(M[i])-1, -1)
	}

	for j := range M[0] {
		visibleCol(M, j, vis, 0, +1)
		visibleCol(M, j, vis, len(M)-1, -1)
	}

	pln(vis)

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
