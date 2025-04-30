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
	var n  int 
	Fscan(in, &n)
	ar := make([]int64, n + 2)
	for i := 0; i < n; i++ {
		Fscan(in, &ar[i + 1])
		ar[i + 1] += ar[i]
	}

	ans := int64(0)
	for i := 1; i <= n; i++ {
		if n % i != 0 {
			continue
		}
		mx, mn := int64(0), int64(1e17)
		for j := i; j <= n; j += i {
			mx = max(mx, ar[j] - ar[j - i])
			mn = min(mn, ar[j] - ar[j - i])
			// Fprintln(out, "i, j => ", i, j, ar[j] - ar[j - i])
		}
		// Fprintln(out, "=> ", mx, mn, mx - mn)
		ans = max(ans, mx - mn)
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