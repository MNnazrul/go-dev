package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

type pt struct {
	x int 
	y int 
	id int 
}

func solve() {
	var n int 
	Fscan(in, &n)
	ar := make([]pt, 0)
	for i := 0; i < n; i++ {
		var x, y int 
		Fscan(in, &x, &y)
		ar = append(ar, pt{x, y, i + 1})
	}

	sort.Slice(ar, func(i, j int) bool {
		if ar[i].x != ar[j].x {
			return ar[i].x < ar[j].x 
		}
		return ar[i].y < ar[j].y 
	})
	n /= 2 
	left := make([]pt, n)
	right := make([]pt, n)

	copy(left, ar[:n])
	copy(right, ar[n:])

	sort.Slice(left, func(i, j int) bool {
		return left[i].y < left[j].y
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i].y > right[j].y 
	})

	for i := 0; i < n; i++ {
		Fprintln(out, left[i].id, right[i].id)
	}
}

func main() {
	defer out.Flush()

	var T int
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}
