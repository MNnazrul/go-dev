package main 

import (
	."fmt"
	"os"
	"bufio"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	mod = int64(998244353)
)

func binPow(a, b int64) int64 {
	a %= mod
	res := int64(1)
	for b > 0 {
		if b % 2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod 
		b /= 2
	}
	return res
}

func modInverse(a int64) int64 {
	return binPow(a, mod - 2) % mod
}

func minusOne(a int64) int64 {
	return (1 - a + mod) % mod
}

type Pair struct {
	l int64
	exist int64
	notExist int64
}

func main() {
	defer out.Flush()
	
	var n, m int64
	Fscan(in, &n, &m)

	adj := make([][]Pair, m + 2)
	mul := make([]int64, m + 2)
	dp := make([]int64, m + 2)

	for _, i := range mul {
		mul[i] = int64(1)
		adj[i] = nil
	}
	dp[0] = 1
	for i := int64(0); i < n; i++ {
		var p, q, l, r int64
		Fscan(in, &l, &r, &p, &q)
		exist := p * modInverse(q) % mod
		notExist := minusOne(exist)

		adj[r] = append(adj[r], Pair{l, exist, notExist})
		mul[r] = mul[r] * notExist % mod
		dp[0] = dp[0] * notExist % mod
	}

	for i, vec := range adj {
		for _, pair := range vec {
			l := pair.l  
			exist := pair.exist
			notExist := pair.notExist
			val := dp[l - 1] * modInverse(notExist) % mod 
			val = val * exist % mod 
			dp[i] = (dp[i] + val) % mod 
			dp[i] = (dp[i] + mod) % mod
		}
	}


	Fprintln(out, dp[m])

}
