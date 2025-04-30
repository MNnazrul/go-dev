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
	if a > b {
		return a 
	}
	return b 
}

func min(a, b int64) int64 {
	if a < b {
		return a 
	}
	return b
}

func solve() {
	var n, m int 
	Fscan(in, &n, &m)
	ar := make([]int, n)
	id := -1
	for i := 0; i < n; i++ {
		Fscan(in, &ar[i])
	}
	br := make([]int, m)
	for i := 0; i < m; i++ {
		Fscan(in, &br[i])
	}
	sort.Ints(br)
	i1, prev := 0, 0 
	for i := 0; i < n; i++ {
		for i1 < m && br[i1] - ar[i] < prev {
			i1++
		}
		if i1 < m {
			ar[i] = min(br[i1] - ar[i], ar[i])
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