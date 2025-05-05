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
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a 
}

func solve() {
	var n, m, k int 
	Fscan(in, &n, &m, &k)

	ans := make([]int, n)
	if (n - (m + 1) * k) / (m + 1) > 0 {
		k += (n - (m + 1) * k) / (m + 1)
	}

	for i := 0; i < n; i++ {
		ans[i] = i % k 
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