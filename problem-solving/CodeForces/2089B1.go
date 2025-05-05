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

func takeInput(n int) []int64 {
	ar := make([]int64, 2 * n)
	for i := 0; i < n; i++ {
		Fscan(in, &ar[i])
		ar[i + n] = ar[i]
	}
	return ar
}

func solve() {
	var n, k int 
	Fscan(in, &n, &k)

	ar := takeInput(n)
	br := takeInput(n)
	for i := 0; i < 2 * n; i++ {
		ar[i] = ar[i] - br[i]
	}
	ans, cnt, sm := 1, 1, int64(0)
	for _, val := range ar {
		sm += val
		if sm > 0 {
			cnt++
		} else {
			if cnt > ans {
				ans = cnt 
			}
			sm = int64(0)
			cnt = 1
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