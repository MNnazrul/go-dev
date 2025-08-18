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

func findPar(u int, par []int) int {
	if par[u] == -1 || par[u] == u {
		return u 
	}
	par[u] = findPar(par[u], par)
	return par[u]
}

func merge(u int, v int , par []int, sz []int) {
	paru := findPar(u, par)
	parv := findPar(v, par)
	if sz[parv] > sz[paru] {
		parv, paru = parv, paru
	}
	sz[paru] += sz[parv]
	par[parv] = paru
}

func solve() {

	var n, m1, m2 int 
	Fscan(in, &n, &m1, &m2)
	par1 := make([]int, n + 4)
	par2 := make([]int, n + 4)
	sz1 := make([]int, n + 4)
	sz2 := make([]int, n + 4)
	adj := make([][]int, n + 4)
	adj1 := make([][]int, n + 4)

	for i, _ := range par1 {
		par1[i] = -1
		par2[i] = -1
		adj[i] = nil
		adj1[i] = nil
		sz1[i] = 0
		sz2[i] = 0
	}

	for i := 0; i < m1; i++ {
		var u, v int 
		Fscan(in, &u, &v)
		if v < u {
			u, v = v, u
		}
		adj[u] = append(adj[u], v)
	}
	for i := 0; i < m2; i++ {
		var u, v int 
		Fscan(in, &u, &v)
		if v < u {
			u, v = v, u
		}
		adj1[u] = append(adj1[u], v)
		merge(u, v, par2, sz2)
	}
	ans := 0
	for i := 1; i <= n; i++ {
		parf1 := findPar(i, par2)
		for _, v := range adj[i] {
			parf2 := findPar(v, par2)
			if parf1 != parf2 {
				ans++
			} else {
				merge(i, v, par1, sz1)
			}
		}
	}

	for i := 1; i <= n; i++ {
		parf1 := findPar(i, par1)
		for _, v := range adj1[i] {
			parf2 := findPar(v, par1)
			if parf1 != parf2 {
				ans++
				merge(i, v, par1, sz1)
			}
		}
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
