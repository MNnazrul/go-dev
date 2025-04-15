/*
link : https://codeforces.com/problemset/problem/2044/F
*/

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	defer out.Flush()

	var n, m, q int

	Fscan(in, &n, &m, &q)
	
	s1 , s2 := 0, 0
	
	ar1 := make([]int, 0)
	ar2 := make([]int, 0)

	for i := 0; i < n; i++ {
		var val int
		Fscan(in, &val)
		s1 += val
		ar1 = append(ar1, val)
	}
	for i := 0; i < m; i++ {
		var val int
		Fscan(in, &val)
		s2 += val
		ar2 = append(ar2, val)
	}

	mp1 := make(map[int]int)
	mp2 := make(map[int]int)

	for i := 0; i < n; i++ {
		mp1[s1 - ar1[i]] = 1
	}
	
	for i := 0; i < m; i++ {
		mp2[s2 - ar2[i]] = 1
	}

	chk := func(x, val int) bool {
		_, ok1 := mp1[val]
		_, ok2 := mp2[x / val]
		_, ok3 := mp1[x / val]
		_, ok4 := mp2[val]
		return (ok1 && ok2) || (ok3 && ok4)
	}

	for ; q > 0; q-- {
		var x int
		Fscan(in, &x)
		flg := false
		for i := 1; i * i <= absInt(x); i++ {
			if x % i == 0 && (chk(x, i) || chk(x, -i))  {
				Fprintln(out, "YES")
				flg = true
				break
			}
		}
		if flg == false {
			Fprintln(out, "NO")
		}
	}
}
