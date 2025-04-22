package main

import (
	."fmt"
	"os"
	"bufio"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func solve() {
	var n int 
	Fscan(in, &n)
	mp1 := make(map[int]int)
	mp2 := make(map[int]int)
	for i := 0; i < n; i++ {
		var x, y int 
		Fscan(in, &x, &y)
		mp1[x] ^= 1
		mp2[x + y] ^= 1
	}
	x, y := 0, 0
	for key, value := range mp1 {
		if value == 1 {
			x = key
		}
	}
	for key, value := range mp2 {
		if value == 1 {
			y = key - x 
			break
		}
	}
	Fprintln(out, x, y)
}

func main() {

	defer out.Flush()

	var T int 
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}