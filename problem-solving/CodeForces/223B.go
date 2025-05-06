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

	var s, t string 
	Fscan(in, &s, &t)

	n, m := len(s), len(t)

	j := 0 
	pre := make([]int, n + 3)
	for id, ch := range s {
		if id > 0 {
			pre[id] = pre[id - 1]
		}
		if j < m && ch == rune(t[j]) {
			pre[id]++
			j++
		}
	}

	mp := make(map[byte]int)
	j, curL := m - 1, 0
	flg := true
	for i := n - 1; i >= 0; i-- {
		preL := 0
		if i > 0 {
			preL = pre[i - 1]
		}
		if j >= 0 && t[j] == s[i] {
			curL++
			mp[s[i]] = curL
			j--
		}
		if len, ok := mp[s[i]]; !ok || len + preL < m {
			flg = false
			break
		} 
	}

	if flg {
		Fprintln(out, "Yes")
	} else {
		Fprintln(out, "No")
	}

}