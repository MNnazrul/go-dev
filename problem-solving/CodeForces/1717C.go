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
const INF int64 = 1e18

func max(a, b int) int {
	if a > b {
		return a 
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a 
	}
	return b
}

func solve() {
	var n int 
	Fscan(in, &n)
	ar := make([]int, n)
	br := make([]int, n)
	mx, id := 0, -1
	for i := 0; i < n; i++ {
		Fscan(in, &ar[i])
		if mx < ar[i] {
			mx = ar[i]
			id = i
		}
	}
	flg := false
	for i := 0; i < n; i++ {
		Fscan(in, &br[i])
		if ar[i] > br[i] {
			flg = true
		}
	}
	if flg {
		Fprintln(out, "NO")
		return
	}
	for cnt := 2 * n; cnt > 0; cnt-- {
		id = (id + n) % n 
		id1 := (id + 1) % n 
		ar[id] = max(ar[id], min(max(ar[id], br[id]), ar[id1] + 1))
		id--
	}
	for i := 0; i < n; i++ {
		ii := (i + 1) % n 
		if !(ar[i] == br[i] && ar[ii] == br[ii]) && br[i] - br[ii] > 1 {
			Fprintln(out, "NO")
			return
		}
	}
	Fprintln(out, "YES")
}

func main() {

	defer out.Flush()

	var T int 
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}