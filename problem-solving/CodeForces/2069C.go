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

const mod = int64(998244353)

func solve() {
	var n int
	Fscan(in, &n)
	ar := make([]int, n)
	pow := make([]int64, n + 1)
	pow[0] = int64(1)
	for i := 0; i < n; i++ {
		Fscan(in, &ar[i])
		if i > 0 {
			pow[i] = pow[i-1] * int64(2) % mod
		}
	}
	res := int64(0)
	ans := int64(0)
	one := int64(0)
	cnt := int64(0)
	for _, val := range ar {
		if val == 1 {
			one++
		} else if val == 2 {
			res = (res + one) * int64(2) % mod
			cnt += one
			one = 0
		} else {
			ans = (ans - cnt + mod) % mod
			ans = (ans + res) % mod
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
