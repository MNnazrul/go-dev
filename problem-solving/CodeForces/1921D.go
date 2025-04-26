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

func abs(a int) int64 {
	if a >= 0 {
		return int64(a)
	}
	return int64(-a) 
}

func max(a, b int64) int64 {
	if a < b {
		return b 
	}
	return a 
}

func solve() {
	var n, m int 
	Fscan(in, &n, &m)
	ar := make([]int, n)
	br := make([]int, m)
	cr := make([]int, m)
	for i := 0; i < n; i++ {
		Fscan(in, &ar[i])
	}
	sort.Ints(ar)
	for i := 0; i < m; i++ {
		Fscan(in, &br[i])
	}
	sort.Ints(cr)
	sort.Ints(br)
	ans1 := int64(0)
	
	l, r := 0, m - 1
	i1, i2 := 0, n - 1

	fun := func(i1, l1, r1 int) int64 {
		c, d := abs(br[l1] - ar[i1]), abs(br[r1] - ar[i1])
		return max(c, d)
	}

	fun2 := func(i1, l1, r1 int) (int, int) {
		c, d := abs(br[l1] - ar[i1]), abs(br[r1] - ar[i1])
		if c > d {
			return l1 + 1, r1
		}
		return l1, r1 - 1
	}

	for i := 0; i < n; i++ {
		a, b := fun(i1, l, r), fun(i2, l, r)
		if a > b {
			l, r = fun2(i1, l, r)
			ans1 += a
			i1++
		} else {
			l, r = fun2(i2, l, r)
			ans1 += b
			i2--
		}
	}
	
	Fprintln(out, ans1)
}

func main() {

	defer out.Flush()

	var T int 
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}