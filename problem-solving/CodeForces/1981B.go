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

func max(a, b int64) int64 {
	if a < b {
		return b 
	}
	return a 
}

func solve() {
	var n, m int64 
	Fscan(in, &n, &m)
	l, r := max(0, n - m), n + m  
	ans := n 
	for bt := 30; bt >= 0; bt-- {
		cn := int64(1 << bt)
		for i := 30; i >= 0; i-- {
			if (cn | int64(1 << i)) <= r {
				cn |= int64(1 << i)
			}
		}
		if cn >= l && cn <= r {
			ans |= cn 
		}
	}
	Fprintln(out, ans)
}

func main() {

	defer out.Flush()

	var T int 
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}