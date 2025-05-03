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

func bigmod(a, b int64) int64 {
	a %= mod 
	res := int64(1)
	for b > 0 {
		if b % 2 == 1 {
			res = res * a % mod 
		}
		a = a * a % mod 
		b >>= 1;
	}
	return res
}

func inverseMod(n int64) int64 {
	return bigmod(n, mod - 2) % mod 
}

func solve() {
	sm := 0
	ar := make([]int, 1)
	for i := 0; i < 26; i++ {
		var val int 
		Fscan(in, &val)
		if val == 0 {
			continue
		}
		ar = append(ar, val)
		sm += val
	}
	fact := make([]int64, sm + 2)
	rfact := make([]int64, sm + 2)
	fact[0] = int64(1)
	rfact[0] = int64(1)

	for i := 1; i <= sm; i++ {
		fact[i] = fact[i - 1] * int64(i) % mod 
		rfact[i] = inverseMod(fact[i])
	}
	
	dp := make([]int64, sm + 2)

	dp[0] = 1
	for _, val := range ar {
		if val == 0 {
			continue
		}
		dp1 := make([]int64, sm + 2)
		for j := 0; j <= sm; j++ {
			if j < val {
				dp1[j] = dp[j]
				continue
			}
			dp1[j] = (dp[j] + dp[j - val]) % mod 
		}
		copy(dp, dp1)
	}

	p := int64(1)
	for _, val := range ar {
		p = p * rfact[val] % mod 
		p = (p + mod) % mod
	}
	ans := fact[sm/2] * fact[sm - (sm / 2)] % mod 
	ans = ans * (p * dp[sm - sm/2] % mod ) % mod 

	Fprintln(out, (ans + mod) % mod)
}

func main() {

	defer out.Flush()

	var T int 
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}