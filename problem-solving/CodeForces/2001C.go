package main

import (
	. "fmt"
)

func ask(a, b int) int {
	if a == b {
		return a
	}
	Println("?", a, b)
	var p int 
	Scan(&p)
	return p 
}

func solve() {
	var n int 
	Scan(&n)
	ans := []any{"!"}
	ok := make([]bool, n + 1)

	for i := 2; n > 1; n-- {
		for ok[i] {
			i++
		}
		v, w := 1, i
		for {
			m := ask(v, w)
			if m == v {
				ok[w] = true
				ans = append(ans, v, w)
				break
			}
			if ok[m] {
				v = m
			} else {
				w = m
			}
		}
	}
	Println(ans...)
}

func main() {
	var T int
	for Scan(&T); T > 0; T-- {
		solve()
	}
}
