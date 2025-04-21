package main

import (
	."fmt"
	"os"
	"bufio"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	dp[][] int64
)
const INF int64 = 1e18

func min(nums ...int64) int64 {
	ans := INF 
	for _, value := range nums {
		if value < ans {
			ans = value
		}
	}
	return ans
}

func chk(ar[] int, hr[][] int, id, op int) int64 {
	if id == len(ar) {
		return 0
	}
	ans := dp[id][op]
	if ans != -1 {
		return ans
	}
	ans = INF 
	flg1, flg2 := true, true
	for idx, value := range hr[id] {
		if hr[id-1][idx] + op == value {
			flg1 = false
		} 
		if hr[id-1][idx] + op == value + 1 {
			flg2 = false
		}
	}

	if flg1 {
		ans = min(ans, chk(ar, hr, id + 1, 0))
	} 
	if flg2 {
		ans = min(ans, chk(ar, hr, id + 1, 1) + int64(ar[id]))
	}
	dp[id][op] = ans
	
	return ans
}

func initiate(n int) {
	dp = make([][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int64, 3)
		for j := 0; j < 3; j++ {
			dp[i][j] = -1
		}
	}
}

func transpose(hr[][] int) {
	dummy := make([][]int, len(hr)) 
	for i := range hr {
		dummy[i] = make([]int, len(hr[i]))
		copy(dummy[i], hr[i])
	}
	for i := 0; i < len(hr); i++ {
		for j := 0; j < len(hr); j++ {
			hr[i][j] = dummy[j][i]
		}
	}
}

func solve() {
	var n int 
	Fscan(in, &n)
	hr := make([][]int, n)
	for i := 0; i < n; i++ {
		hr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			Fscan(in, &hr[i][j])
		}
	}
	ar := make([]int, n)
	br := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &ar[i])
	}
	for j := 0; j < n; j++ {
		Fscan(in, &br[j])
	}
	initiate(n)
	ans1 := min(chk(ar, hr, 1, 0), chk(ar, hr, 1, 1) + int64(ar[0]))
	initiate(n)
	transpose(hr)
	ans2 := min(chk(br, hr, 1, 0), chk(br, hr, 1, 1) + int64(br[0]))

	if ans1 >= INF || ans2 >= INF {
		Fprintln(out, -1)
		return
	}
	Fprintln(out, ans1 + ans2)
}

func main() {

	defer out.Flush()

	var T int 
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}