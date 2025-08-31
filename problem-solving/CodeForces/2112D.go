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

type pair struct {
	x int 
	y int 
}

func solve() {
	var n int
	Fscan(in, &n)
	adj := make([][]int, n + 1)
	ans := make([]pair, 0)
	for i := 0; i < n - 1; i++ {
		var u, v int 
		Fscan(in, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	id := -1
	for i := 1; i <= n; i++ {
		if len(adj[i]) == 2 {
			id = i 
			break
		}
	}
	if id == -1 {
		Fprintln(out, "NO")
		return
	}
	ans = append(ans, pair{adj[id][0], id}, pair{id, adj[id][1]})
	queue := make([]pair, 0)
	queue = append(queue, pair{adj[id][0], 1}, pair{adj[id][1], 0})
	clr := make([]bool, n + 1)
	clr[id] = true 
	clr[adj[id][0]] = true
	clr[adj[id][1]] = true

	for len(queue) > 0 {
		pc := queue[0]
		node := pc.x 
		id := pc.y
		queue = queue[1:]
		for _, cur := range adj[node] {
			if clr[cur] {
				continue
			}
			queue = append(queue, pair{cur, id ^ 1})

			if id == 1 {
				ans = append(ans, pair{node, cur})
			} else {
				ans = append(ans, pair{cur, node})
			}

			clr[cur] = true
		}
	}
	Fprintln(out, "YES")
	for _, pc := range ans {
		Fprintln(out, pc.x, pc.y)
	}

}

func main() {
	defer out.Flush()

	var T int
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}
