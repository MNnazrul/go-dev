package main 

import (
	."fmt"
	"os"
	"bufio"
	"sort"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func solve() {
	var n, m int 
	Fscan(in, &n, &m)
	mp1 := make(map[int]int)
	ar := make([]int, 0)
	for i := 0; i < n; i++ {
		var val int 
		Fscan(in, &val)
		mp1[val]++
	}
	for i := 0; i < m; i++ {
		var val int 
		Fscan(in, &val)
		if mp1[val] > 0 {
			mp1[val]--
		} else {
			ar = append(ar, val)
		}
	}

	sort.Ints(ar)

	for _, val := range ar {
		queue := make([]int, 1)
		queue[0] = val 
		for len(queue) > 0 && len(queue) <= n {
			top := queue[0]
			queue = queue[1:]
			if mp1[top] > 0 {
				mp1[top]--
				continue
			}
			queue = append(queue, top/2)
			queue = append(queue, top - top/2)
			if top <= 1 {
				break	
			}
		}
		if len(queue) > 0 {
			Fprintln(out, "No")
			return
		}
	}
	for _, val := range mp1 {
		if val != 0 {
			Fprintln(out, "No")
			return
		}
	}
	Fprintln(out, "Yes")
}

func main() {

	defer out.Flush()

	var T int 
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}