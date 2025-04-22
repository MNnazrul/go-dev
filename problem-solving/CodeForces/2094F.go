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
	var n, m, k int 
	Fscan(in, &n, &m, &k)
	ar := make([]int, k)
	for i := 0; i < k; i++ {
		ar[i] = i + 1
	}
	idx := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			Fprint(out, ar[idx % k], " ")
			idx++
		}
		Fprintln(out, )
		if m % k == 0 {
			val := ar[0]
			ar = ar[1:]
			ar = append(ar, val)
		} 
	}
}

func main() {

	defer out.Flush()

	var T int 
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}