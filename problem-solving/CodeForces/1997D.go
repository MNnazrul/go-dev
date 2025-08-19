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

func dfs(root int, adj[][]int, ar []int) (int, int) {
	if len(adj[root]) == 0 {
		return ar[root], ar[root]
	}
	mn := 1000000005
	for _, child := range adj[root] {
		_, mn1 := dfs(child, adj, ar)
		if mn > mn1 {
			mn = mn1
		}
	}
	mx := (ar[root] + mn)
	if ar[root] < mn {
		mn = (ar[root] + mn ) / 2
	}
	return mx, mn 
}

func solve() {
	var n int 
	Fscan(in, &n)
	ar := make([]int, n + 1)
	adj := make([][]int, n + 1)
	for i := 1; i <= n; i++ {
		Fscan(in, &ar[i])
		adj[i] = nil
	}
	for i := 2; i <= n; i++ {
		var p int 
		Fscan(in, &p)
		adj[p] = append(adj[p], i)
	}

	ans, _ := dfs(1, adj, ar)

	Fprintln(out, ans)
}

func main() {
	defer out.Flush()

	var T int
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}
