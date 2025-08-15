package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func check(mp map[int64]bool, sm int64) bool {
	if sm < 0 {
		if _, ok := mp[-sm]; ok {
			return false
		} 
		return true
	}
	if _, ok := mp[sm]; ok {
		return false
	}
	return true
}

func solve() {
	var n int 
	Fscan(in, &n)
	ar := make([]int64, 2*n)
	sm := int64(0)
	for i := 0; i < 2*n; i++ {
		Fscan(in, &ar[i])
	}
	
	sort.Slice(ar, func(i, j int) bool {
		return ar[i] > ar[j]
	})

	ans := make([]int64, 2 * n + 1)
	ans[0] = ar[0]
	l, r := 2, 3
	for i := 1; i < 2 * n; i++ {
		if i <= n {
			sm -= ar[i]
			ans[l] = ar[i]
			l += 2
		} else {
			sm += ar[i]
			ans[r] = ar[i]
			r += 2
		}
	}
	ans[1] = ans[0] - sm

	for _, val := range ans {
		Fprint(out, val, " ")
	}
	Fprintln(out, )
}

func main() {
	defer out.Flush()

	var T int
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}
