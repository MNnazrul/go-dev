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

func min(a, b int ) int {
	if a < b {
		return a
	}
	return b
}

func solve() {
	var n int 
	Fscan(in, &n)
	ar := make([]int, n)
	ps := make([]int, n + 1)
	for i := 0; i < n; i++ {
		Fscan(in, &ar[i])
		ps[ar[i]] = i 
	}
	ans := 0
	for i := 1; i < n; i++ {
		mn1, mn2 := 0, 0  
		for j := 0; j < ps[i]; j++ {
			if ar[j] < i {
				continue
			}
			if ar[j] > i {
				mn1++
			}
		}
		for j := ps[i] + 1; j < n; j++ {
			if ar[j] < i {
				continue
			}
			if ar[j] > i {
				mn2++
			}
		}
		ans += min(mn1, mn2)
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
