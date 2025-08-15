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

const mod = int64(998244353)

func gcd(a int, b int) int {
	if b == 0 {
		return a 
	}
	return gcd(b, a % b)
}

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
	var n int 
	Fscan(in, &n)
	ar := make([]int, n)
	gc := 0
	dp := make(map[int]int)
	for i := 0; i < n; i++ {
		Fscan(in, &ar[i])
		gc = gcd(gc, ar[i])
	}
	cnt := 0
	flg := false
	for _, val := range ar {
		if val != gc {
			cnt++
		}
		if val == gc {
			flg = true
		}
	}
	if flg {
		Fprintln(out, cnt)
		return
	}
	sort.Ints(ar)
	for i, _ := range ar {
		dp[ar[i]] = 0
		for key := range dp {
			x := gcd(ar[i], key)
			if _, ok := dp[x]; ok {
				dp[x] = min(dp[x], dp[key] + 1)
			} else {
				dp[x] = dp[key] + 1
			}
		}
	}
	Fprintln(out, dp[gc] + n - 1)
}

func main() {
	defer out.Flush()

	var T int
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}
