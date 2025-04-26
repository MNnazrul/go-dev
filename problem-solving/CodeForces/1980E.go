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

func input2DSlice(ar[][] int, n, m int) {
	flg := -1
	for i := 0; i < n; i++ {
		ar[i] = make([]int, m)
		for j := 0; j < m; j++ {
			Fscan(in, &ar[i][j])
			if ar[i][j] == 1 {
				flg = j
			}
		}
	}
	if flg != 0 {
		for j := 0; j < n; j++ {
			ar[j][0], ar[j][flg] = ar[j][flg], ar[j][0]
		}
	}
	sort.Slice(ar, func(i, j int) bool {
		return ar[i][0] <= ar[j][0]
	})
}

func solve() {
	var n, m int 
	Fscan(in, &n, &m)
	ar := make([][]int, n)
	br := make([][]int, n)
	
	input2DSlice(ar, n, m)
	input2DSlice(br, n, m)

	mp := make(map[int]int)
	for i := 0; i < m; i++ {
		mp[ar[0][i]] = i
	}

	for i := 0; i < m; i++ {
		id, ok := mp[br[0][i]];
		if !ok {
			Fprintln(out, "NO")
			return
		}
		for j := 0; j < n; j++ {
			if ar[j][id] != br[j][i] {
				Fprintln(out, "NO")
				return
			} 
		}
	}
	Fprintln(out, "YES")
}

func main() {

	defer out.Flush()

	var T int 
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}