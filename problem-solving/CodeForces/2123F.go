package main 

import (
	."fmt"
	"os"
	"bufio"
	"strings"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	prime []bool
)

func solve() {
	var n int 
	Fscan(in, &n)
	ans := make([]int, n)
	clr := make([]bool, n)
	for i := 0; i < n; i++ {
		ans[i] = i + 1
	}
	for i := n; i >= 2; i-- {
		if prime[i] {
			continue
		}
		dc := make([]int, 0)
		ps := make([]int, 0)
		ps = append(ps, i - 1)
		for j := i + i; j <= n; j += i {
			if !clr[j - 1] {
				clr[j - 1] = true
				dc = append(dc, j)
				ps = append(ps, j - 1)
			}
		}
		dc = append(dc, i)
		for i1 := 0; i1 < len(ps); i1++ {
			ans[ps[i1]] = dc[i1]
		}
	}
	Fprintln(out, strings.Trim(Sprint(ans), "[]"))
}

func main() {
	defer out.Flush()
	var T int
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}

func init() {
	prime = make([]bool, 100005)
	for i := 2; i < 100005; i++ {
		if prime[i] {
			continue
		}
		for j := i + i; j < 100005; j+= i {
			prime[j] = true
		}
	}
}