package main

import (
	"bufio"
	. "fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func max(a, b int) int {
	if a > b {
		return a 
	}
	return b 
}

func solve() {
	var n int 
	Fscan(in, &n)
	ar := make([]int, n)
	br := make([]int, n)

	posAr := make([]int, n + 1)
	posBr := make([]int, n + 1)

	for i := 0; i < n; i++ {
		Fscan(in, &ar[i])
	}
	for i := 0; i < n; i++ {
		Fscan(in, &br[i])
	}
	ans := 0 
	for i := n - 1; i >= 0; i-- {
		if ar[i] == br[i] {
			ans = i + 1
			break
		}

		if posAr[ar[i]] - i > 0 || posBr[ar[i]] - i > 1 {
			ans = i + 1
			break
		} 

		if posBr[br[i]] - i > 0 || posAr[br[i]] - i > 1 {
			ans = i + 1
			break
		}

		posAr[ar[i]] = max(posAr[ar[i]], i)
		posBr[br[i]] = max(posBr[br[i]], i)
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
