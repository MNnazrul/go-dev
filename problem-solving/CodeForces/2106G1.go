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

func ask(id, k int, u ...int) int {
	ans := 0
	if id == 1 {
		Fprintln(out, "?", id, k, u[0])
		out.Flush()
		Fscan(in, &ans)
	} else {
		Fprintln(out, "?", id, k)
		out.Flush()
	}
	return ans
}

func chk() int {
	ans := ask(1, 1, 1)
	if ans == 0 {
		ask(2, 1)
		ans = ask(1, 1, 1)
	}
	return ans
}

func solve() {
	var n int 
	Fscan(in, &n)
	for i := 0; i < n - 1; i++ {
		var u, v int 
		Fscan(in, &u, &v)
	}
	sm := chk()
	ans := make([]int, n)
	if sm == 1 || sm == -1 {
		ans[0] = sm 
		for i := 1; i < n; i++ {
			ans[i] = ask(1, 1, i + 1) - sm
		}
	} else {
		ans[0] = -(sm >> 1)
		ask(2, 1)
		for i := 1; i < n; i++ {
			p := ask(1, 1, i + 1)
			if p != -1 && p != 1 {
				p = sm >> 1
			}
			ans[i] = p 
		}
	}
	Fprint(out, "! ")
	for _, val := range ans {
		Fprint(out, val, " ")
	}
	Fprintln(out, )
	out.Flush()
}

func main() {

	defer out.Flush()

	var T int 
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}