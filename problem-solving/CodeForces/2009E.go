package main 

import (
	."fmt"
	"os"
	"bufio"
	// "strconv"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	id1 = 0 
)

func chk(a, b int64) int64 {
	ans := (b * (b + 1) - a * (a - 1)) / 2
	return ans
}

func min(a, b int64) int64 {
	if a < b {
		return a 
	}
	return b 
}

func bin(n, k int64, flg bool) int64 {
	l, r := int64(n), int64(k - 1)
	ans := chk(n, k)
	for l <= r {
		md := ((l + r )>> 1)
		l1, r1 := chk(n, md), chk(md + 1, k)
		if flg {
			if l1 <= r1 {
				ans = min(ans, r1 - l1)
				l = md + 1
			} else {
				r = md - 1
			}
		} else {
			if l1 >= r1 {
				ans = min(ans, l1 - r1)
				r = md - 1
			} else {
				l = md + 1
			}
		}
	}
	return ans
}

func solve() {
	id1++
	var n, k int64 
	Fscan(in, &n, &k)
	if n < 3 {
		if n == 1 {
			Fprintln(out, k)
			return
		}
		Fprintln(out, 1)
		return
	}
	Fprintln(out, min(bin(k, k+n-1, true), bin(k, k+n-1, false)))
}

func main() {

	defer out.Flush()

	var T int 
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}