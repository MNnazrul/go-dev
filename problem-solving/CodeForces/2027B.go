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
	var n, val int 
	Fscan(in, &n) 
	mp := make(map[int]int)
	firstPos := make(map[int]int)
	for i := 0; i < n; i++ {
		Fscan(in, &val)
		mp[val]++
		_, ok := firstPos[val]
		if !ok {
			firstPos[val] = i
		}
	}
	keys := make([]int, 0)
	for key := range mp {
		keys = append(keys, key)
	} 
	sort.Slice(keys, func(i, j int) bool {
		if keys[i] > keys[j] {
			return true
		}
		return false
	})
	ans, tot := n, 0
	for _, val := range keys {
		ans = min(ans, tot + firstPos[val])
		tot += mp[val]
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