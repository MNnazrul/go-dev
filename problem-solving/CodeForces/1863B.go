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

func solve() {
	var n int 
	Fscan(in, &n)
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		var val int 
		Fscan(in, &val)
		pos[val-1] = i 
	}
	ans := 0
 	for i := 1; i < n; i++ {
		if pos[i-1] > pos[i] {
			ans++
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