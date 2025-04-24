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
const INF int64 = 1e18

func max(a, b int64) int64 {
	if a > b {
		return a 
	}
	return b 
}

func solve() {
	var n int  
	Fscan(in, &n)
	ar := make([]int64, n)
	mp := make(map[int64][]int64)
	for i := 0; i < n; i++ {
		Fscan(in, &ar[i])
		l := ar[i] + int64(i)
		mp[l] = append(mp[l], int64(i))
	}
	keys := make([]int64, 0)
	for key := range mp {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		if keys[i] > keys[j] {
			return true
		}
		return false
	})
	ans := make(map[int64]int64)
	for _, l := range keys {
		for _, id := range mp[l] {
			ans[l] = max(ans[l], max(l + id, ans[l + id]))
		}
	}
	Fprintln(out, max(int64(n), ans[int64(n)]))
}

func main() {

	defer out.Flush()

	var T int 
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}