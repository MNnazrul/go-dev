package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"sort"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a , b int) int {
	if a < b {
		return a
	}
	return b
}

func chk(mp map[int]int, ar []int, i1, i2 int) {
	sm, mn, mx := 0, math.MaxInt, math.MinInt
	mn1 , mx1 := 0, 0
	mn, mx = 0, 0
	
	for i := i1; i < i2; i++ {
		sm += ar[i]
		mn = min(mn, sm)
		mx = max(mx, sm)
	
		mn1 = min(mn1, sm - mx)
		mx1 = max(mx1, sm - mn)
	}
	
	for i := mn1; i <= mx1; i++ {
		mp[i] = 1
	}
}

func PrintAns(ans map[int]int) {
	// Fprintln(out, len(ans))
	// for key, _ := range ans {
	// 	Fprint(out, key, " ")
	// }

	keys := make([]int, 0, len(ans))
	for key := range ans {
		keys = append(keys, key)
	}

	sort.Ints(keys)
	Fprintln(out, len(keys))
	for _, val := range keys {
		Fprint(out, val, " ")
	}
	Fprintln(out, )
}

func solve() {

	var n int
	Fscan(in, &n)
	ar := make([]int, n)
	ans := make(map[int]int)

	for i := 0; i < n; i++ {
		Fscan(in, &ar[i])
	}

	idx := n

	for index, value := range ar {
		if value != -1 && value != 1 {
			idx = index
			break
		}
	}

	// Fprintln(out, idx)

	chk(ans, ar, 0, idx)
	
	if idx == n {
		PrintAns(ans)
		return
	}
	
	chk(ans, ar, idx + 1, n)

	// sm := 0
	// for i := idx; i >= 0; i-- {
	// 	sm += ar[i]
	// 	ans[sm] = 1
	// }
	// sm = 0
	// for i := idx; i < n; i++ {
	// 	sm += ar[i]
	// 	ans[sm] = 1
	// }

	mn, mx, sm := 0, 0, 0
	
	for i := idx - 1; i >= 0; i-- {
		sm += ar[i]
		mx = max(mx, sm)
		mn = min(mn, sm)
	}
	// Fprintln(out, mx, mn)
	mn1, mx1 := mn, mx
	sm = 0
	for i := idx + 1; i < n; i++ {
		sm += ar[i]
		mx1 = max(sm + mx, mx1)
		mn1 = min(sm + mn, mn1)
	}

	// Fprintln(out, mx1, mn1)

	for i := mn1; i <= mx1; i++ {
		ans[i + ar[idx]] = 1
	}

	PrintAns(ans)
}

func main() {
	defer out.Flush()
	var T int
	for Fscan(in, &T) ; T > 0; T-- {
		solve()
	}
}

