package main 

import (
	."fmt"
	"os"
	"bufio"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	pw []int64
)

func min(a, b int64) int64 {
	if a < b {
		return a 
	}
	return b 
}

func check(i int64) int64 {
	if i == 0 {
		return int64(3)
	}
	return pw[i + 1] + i * pw[i - 1]
}

func solve() {
	var n, k int64
	Fscan(in, &n, &k)
	i1 := int64(len(pw) - 1)
	mp := make(map[int64]int64)
	for i1 >= 0 {
		if n >= pw[i1] {
			n -= pw[i1]
			k--
			mp[i1]++
		} else {
			i1--
		}
	}
	if k < 0 {
		Fprintln(out, -1)
		return
	}
	for i := int64(len(pw) - 1); i > 0; i-- {
		val, ok := mp[i]
		if !ok {
			continue
		}
		mn := min(k / 2, val)
		mp[i] -= mn 
		mp[i - 1] += mn * 3
		k -= mn * 2
	}
	ans := int64(0)
	for key, value := range mp {
		ans += check(key) * value
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

func init() {
	pw = make([]int64, 20)
	pw[0] = 1
	for i := 1; i < 20; i++ {
		pw[i] = pw[i-1] * 3
	}
}
