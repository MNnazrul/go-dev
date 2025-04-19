package main

import (
	."fmt"
	"os"
	"bufio"
	"sort"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func max_min(a, b int64) (int64, int64) {
	if a > b {
		return a, b
	}
	return b, a
}


func solve() {
	var n, k int 
	Fscan(in, &n, &k)
	ar := make([]int64, n)

	for i := 0; i < n; i++ {
		Fscan(in, &ar[i])
	}
	ans := int64(0)
	for i := 0; i < n; i++ {
		var val int64
		Fscan(in, &val)
		_, mn := max_min(ar[i], val)
		ans += val + ar[i]
		ar[i] = mn
	}

	sort.Slice(ar, func(i, j int) bool {
		return ar[i] < ar[j]
	})

	for idx, value := range ar {
		if idx <= n - k {
			ans -= value
		}
	}
	ans++

	Fprintln(out, ans)
}

func main() {

	defer out.Flush()

	var T int 
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}