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
	divs[][] int
	pos map[int][]int
)
const N = 1e5 + 5

type Query struct {
	L int
	R int 
	K int
	Prev int
}

func generateDivs() {
	divs = make([][]int, N)
	for i := 2; i < N; i++ {
		for j := i; j < N; j += i {
			divs[j] = append(divs[j], i)
		}
	}
}

func lowerBound(div, L , R int) int {
	l, r := 1, len(pos[div])
	ans := int(N + 2)
	for l <= r {
		mid :=  (l + r) >> 1
		if pos[div][mid-1] < L {
			l = mid + 1
		} else {
			if ans > pos[div][mid-1] {
				ans = pos[div][mid-1]
			}
			r = mid - 1
		}
	}
	return ans
}

func solve() {
	var n, q int 
	Fscan(in, &n, &q)
	ar := make([] int, n)
	pos = make(map[int][]int)
	for i := 0; i < n; i++ {
		Fscan(in, &ar[i])
		pos[ar[i]] = append(pos[ar[i]], i + 1)
	}
	for i := 0; i < q; i++ {
		var l, r, k int 
		Fscan(in, &k, &l, &r)
		pox := make([]int, 0)
		for _, div := range divs[k] {
			nb := lowerBound(div, l, r)
			if nb <= r {
				pox = append(pox, nb)
			}
		}
		sort.Ints(pox)
		prev, ans := l-1, 0
		for _, idx := range pox {
			ans += (idx - prev - 1) * k
			for k % ar[idx - 1] == 0 {
				k /= ar[idx - 1]
			}
			ans += k 
			prev = idx 
		}
		ans += (r - prev) * k 
		Fprintln(out, ans)
	}
	
}

func main() {

	defer out.Flush()

	generateDivs()

	var T int 
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}