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
	var n, m int 
	Fscan(in, &n, &m)
	str := make([]string, n)
	for i := 0; i < n; i++ {
		Fscan(in, &str[i])
	}
	mp := make(map[int]int)
	mx := 0
	for i := 0; i < m; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			var s string 
			Fscan(in, &s)
			if s == str[j] {
				cnt++
				mp[j] = 1
			}
		}
		if mx < cnt {
			mx = cnt 
		}
	}
	if len(mp) != n {
		Fprintln(out, -1)
		return
	}
	Fprintln(out, 3 * n - 2 * mx)
}

func main() {

	defer out.Flush()
	
	var T int 
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}