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
	ar := make([]int, n + 1)
	br := make([]int, n + 1)
	for i := 0; i < n; i++ {
		Fscan(in, &ar[i])
	}
	for i := 0; i < n; i++ {
		Fscan(in, &br[i])
	}
	ans := 0 
	for i := n-1; i >= 0; i-- {
		if ar[i] == br[i] || ar[i] == ar[i + 1] || br[i] == br[i + 1] {
			ans = i + 1
			break
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