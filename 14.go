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

var M [][]byte

const SZ = 1000

func draw(si, sj, ei, ej int) {
	if ej != sj && si != ei {
		pln(sj, ej, ei != ej)
		panic("unhandled")
	}
	if si > ei {
		si, ei = ei, si
	}
	if sj > ej {
		sj, ej = ej, sj
	}
	for i := si; i <= ei; i++ {
		for j := sj; j <= ej; j++ {
			M[i][j] = '#'
		}
	}
}

func show() {
	for i := 0; i < 10; i++ {
		for j := 493; j <= 503; j++ {
			if M[i][j] == 0 {
				pf(".")
			} else {
				pf("%c", M[i][j])
			}
		}
		pln()
	}
	pln()
}

func drip() bool {
	j := 500
	for i := 0; i < len(M); i++ {
		if i+1 >= len(M) {
			return false
		}
		if M[i+1][j] != 0 {
			if M[i+1][j-1] == 0 {
				j--
				continue
			}
			if M[i+1][j+1] == 0 {
				j++
				continue
			}
			M[i][j] = 'o'
			return true
		}
	}
	return false
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	M = make([][]byte, SZ)
	for i := range M {
		M[i] = make([]byte, SZ)
	}
	maxi := 0
	for _, line := range lines {
		v := Getints(line, false)

		first := true
		previ, prevj := 0, 0

		for k := 0; k < len(v); k += 2 {
			i := v[k+1]
			j := v[k]
			if i > maxi {
				maxi = i
			}
			if !first {
				//pf("%d,%d -> %d,%d\n", previ, prevj, i, j)
				draw(previ, prevj, i, j)
			}
			previ, prevj = i, j
			first = false
		}
	}

	pln(maxi)
	maxi += 2

	for k := 0; k < 100000; k++ {
		if !drip() {
			pln("overflow", k)
			Sol(k)
			break
		}
	}

	for j := 0; j < len(M[maxi]); j++ {
		if M[maxi][j] != 0 {
			panic("unhandled")
		}
		M[maxi][j] = '#'
	}

	for k := 0; k < 100000; k++ {
		drip()
		if M[0][500] == 'o' {
			pln("blocked")
			break
		}
	}

	cnt := 0
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[i]); j++ {
			if M[i][j] == 'o' {
				cnt++
			}
		}
	}
	Sol(cnt)
}
