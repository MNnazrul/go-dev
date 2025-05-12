package main 

import (
	."fmt"
	"os"
	"bufio"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)
const mod = int64(998244353)
var n, k int 

func sol(id, flg int) int64 {
	if id == n + 1 {
		return 1 
	}
	if dp[id][flg] != -1  {
		return dp[id][flg]
	}
	dp[id][flg] = 0

	if flg == 1 {
		if ar[id] == ar[id - 1] && ar[id] != -1 {
			return 0
		}
		
	} else {

	}
	return dp[id][flg]
}

func main() {
	defer out.FLush()

	Fscan(in, &n, &k)
	ar := make([]int64, n + 2)
	dp := make([][]int64, n + 2)
	for i := 1; i <= n; i++ {
		Fscan(in, &ar[i])
		dp[i] = make([]int, 2)
		dp[i][0] = 1
		dp[i][2] = 1
	}

	ans := sol(1, 0)
}