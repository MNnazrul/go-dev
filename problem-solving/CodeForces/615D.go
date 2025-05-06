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
const mod = int64(1e9 + 7)

func bigmod(a, b int64) int64 {
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

func inverseMod(n int64) int64 {
	return bigmod(n, mod - 2) % mod 
}

func solve() {
	var n int
	Fscan(in, &n)
	mp := make(map[int64]int64) 
	for i := 0; i < n; i++ {
		var val int64 
		Fscan(in, &val)
		mp[val]++
	}
	_, ans := int64(1), int64(1)
	keys := make([]int64, 0)
	for key := range mp {
		keys = append(keys, key)
	}
	suf, pre := make([]int64, len(keys) + 1), make([]int64, len(keys))
	
	prev := int64(1)
	for i := 0; i < len(keys); i++ {
		prev = prev * (mp[keys[i]] + 1) % (mod - 1)
		pre[i] = prev 
	}
	
	suf[len(keys)] = int64(1)
	for i := len(keys) - 1; i >= 0; i-- {
		suf[i] = suf[i + 1] * (mp[keys[i]] + 1) % (mod - 1)
	}

	prev = int64(1)
	for id, key := range keys {
		pwr := prev * suf[id + 1] % (mod - 1)
		val := bigmod(key, mp[key] * (mp[key] + 1) / 2) % mod 
		ans = ans * bigmod(val, pwr) % mod
		prev = pre[id]
	}

	Fprintln(out, ans)
}

func main() {

	defer out.Flush()

	// var T int 
	// for Fscan(in, &T); T > 0; T-- {
		solve()
	// }
}