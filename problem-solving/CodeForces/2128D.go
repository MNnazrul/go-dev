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

const mod = int64(998244353)

func solve() {
	var n int64  
	Fscan(in, &n)
	ar := make([]int64, n)
	ans := int64(0)
	for i := int64(0); i < n; i++ {
		Fscan(in, &ar[i])
		ans += (i + 1) * (n - i)
		if i > int64(0) && ar[i-1] < ar[i] {
			ans -= i * (n - i)
		}
		// Fprintln(out, ans)
		// Fprintln(out, i + 1, n - i)
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
