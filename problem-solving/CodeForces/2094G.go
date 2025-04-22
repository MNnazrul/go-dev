package main

import (
	."fmt"
	"os"
	"bufio"
	"container/list"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func solve() {
	var q, id, val int 
	n, sm := 0, int64(0)
	ans1, ans2 := int64(0), int64(0)
	lst := list.New()
	flg := true

	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &id)

		switch id {
		case 1:
			if flg {
				val = lst.Back().Value.(int)
				lst.Remove(lst.Back())
				lst.PushFront(val)
			} else {
				val = lst.Front().Value.(int)
				lst.Remove(lst.Front())
				lst.PushBack(val)
			}
			ans1, ans2 = ans1 - int64(n * val) + sm, ans2 - sm + int64(n * val)
		case 2:
			ans1, ans2 = ans2, ans1
			flg = !flg
		case 3:
			Fscan(in, &val)
			n++
			ans1 += int64(n * val)
			ans2 += sm + int64(val)
			sm += int64(val)
			if flg {
				lst.PushBack(val)
			} else {
				lst.PushFront(val)
			}
		}

		Fprintln(out, ans1)
	}
}

func main() {

	defer out.Flush()

	var T int 
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}