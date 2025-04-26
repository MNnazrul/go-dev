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
	ans := make([]int, n)
	str := make([]string, n)

	for i := 0; i < n; i++ {
		Fscan(in, &str[i])
	}

	for i := n - 1; i >= 0; i-- {
		cnt := 0
		for j := 0; j < i; j++ {
			if str[i][j] == '1' {
				cnt++
			}
		}
		i1 := 0
		for j := 0; j < cnt; j++ {
			for ans[i1] != 0 {
				i1++
			}
			i1++
		} 
		for ans[i1] != 0 {
			i1++
		}
		ans[i1] = i + 1
	}

	for _, val := range ans {
		Fprint(out, val, " ")
	}
	Fprintln(out, )
}

func main() {

	defer out.Flush()

	var T int 
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}