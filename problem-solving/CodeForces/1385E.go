package main 

import (
	."fmt"
	"os"
	"bufio"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	isCycle bool
)

func dfs(gr [][]int, clr []int, node int) {
	clr[node] = 1
	for _, child := range gr[node] {
		if clr[child] == 0 {
			dfs(gr, clr, child)
		}
		if clr[child] == 1 {
			isCycle = true
		}
	}
	clr[node] = 2
}

func solve() {
	var k, u, v, n, m int 
	Fscan(in, &n, &m)
	gr := make([][]int, n + 1)
	gr1 := make([][]int, n + 1)
	clr := make([]int, n + 1)
	indeg := make([]int, n + 1)
	ans := make([][]int, n + 1)
	for i := 0; i < m; i++ {
		Fscan(in, &k, &u, &v)
		if k == 1 {
			ans[u] = append(ans[u], v)
			gr[u] = append(gr[u], v)
			indeg[v]++
		} else {
			gr1[u] = append(gr1[u], v)
			gr1[v] = append(gr1[v], u)
		}
	}
	isCycle = false
	for i := 1; i <= n; i++ {
		if clr[i] == 0 {
			dfs(gr, clr, 1) 
		}
	}
	if isCycle {
		Fprintln(out, "NO")
		return
	}

	clr = make([]int, n + 1)
	

	queue := make([]int, 0)
	for i := 1; i <= n; i++ {
		if indeg[i] == 0 && clr[i] == 0 {
			queue = append(queue, i)
			clr[i] = 1
		}
	
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			clr[cur] = 2
			for _, val := range gr[cur] {
				if clr[val] == 0 {
					clr[val] = 1
					queue = append(queue, val)
				} 
			}
			for _, val := range gr1[cur] {
				if clr[val] == 2 {
					ans[val] = append(ans[val], cur)
				} else if clr[val] == 0 {
					clr[val] = 1
					queue = append(queue, val)
				}
			}
		}
	}

	Fprintln(out, "YES")
	for i := 1; i <= n; i++ {
		for _, val := range ans[i] {
			Fprintln(out, i, val)
		}
	}

}

func main() {
	defer out.Flush()
	var T int
	Fscan(in, &T)
	for T > 0 {
		T--
		solve()
	}
}