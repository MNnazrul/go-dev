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


func main() {

	defer out.Flush()

	var n, c int 
	Fscan(in, &n, &c)
	ar := make([]int, n + 1)
	suf := make([]int, n + 2)
	for i := 1; i <= n; i++ {
		Fscan(in, &ar[i])
	}
	for i := n; i >= 0; i-- {
		suf[i] = suf[i + 1]
		if ar[i] == c {
			suf[i]++
		}
	}

	mp := make(map[int]int)
	cnt, ans := 0, suf[0]

	for i := 1; i <= n; i++ {

		mp[ar[i]]++
		if mp[ar[i]] < cnt + 1 {
			mp[ar[i]] = cnt + 1
		} 

		if ans < mp[ar[i]] + suf[i + 1] {
			ans = mp[ar[i]] + suf[i + 1]
		}

		if ar[i] == c {
			cnt++
		}
	}

	Fprintln(out, ans)

}