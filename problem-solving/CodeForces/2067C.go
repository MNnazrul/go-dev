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

func chk(P int64) bool {
	for P > 0 {
		k := P % 10
		if k == 7 {
			return true
		}
		P /= 10
	}
	return false
}

func sol(val, kc int64) int {
	if chk(val) {
		return 0;
	}
	ans := 9
	for i := 1; i < 9; i++ {
		ans = min(ans, sol(val + int64(i) * kc, kc * 10 + 9) + i)
	}
	return ans 
}

func min(a, b int) int {
	if a < b {
		return a 
	}
	return b
}

func solve() {
	var n int64
	Fscan(in, &n)
	ck := int64(9)
	if chk(n) {
		Fprintln(out, 0)
		return
	}
	ans := 7
	for i := 2; i <= 9; i++ {
		for j := 0; j < 10; j++ {
			if chk(n + ck * int64(j)) {
				ans = min(ans, j)
			}
			if chk(n + ck * int64(j) + ck*10+9) {
				ans = min(ans, j+1)
			}
		}
		ck = ck * 10 + 9
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