// link : https://codeforces.com/problemset/problem/2043/D

package main

import (
	"bufio"
	. "fmt"
	"os"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func solve() {
	
	var l, r, G int64

	Fscan(in, &l, &r, &G)

	p1 := (l + G - 1) / G 
	p2 := r / G 

	var mx, ans1, ans2 int64 = -1, 0, 0
	
	for i := p1; i < p1 + 10; i++ {
		for j := p2; j >= p2 - 10; j-- {
			if i * G <= r && j * G >= l && i <= j {
				if gcd(i, j) == 1 {
					if mx < j * G - i * G {
						mx = j * G - i * G
						ans1 = i * G
						ans2 = j * G
					}
				}
			}
		}
	}

	if mx > -1 {
		Fprintln(out, ans1, ans2)
		return
	} 
	Fprintln(out, "-1 -1")
	
}

func main() {
	defer out.Flush()
	var T int
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}