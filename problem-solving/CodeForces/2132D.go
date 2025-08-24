package main 

import (
	."fmt"
	"os"
	"bufio"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	pw []int64
)

func sol(n int64) []int64 {
	ar := make([]int64, 0)
	for n > 0 {
		ar = append(ar, n % 10)
		n /= 10
	}
	return ar
}

func size(i, ps int64) int64 {
	mn := i - 1
	if i <= ps {
		mn++
	}
	return mn 
}

func solution(id, ps, val int64, ar[] int64) int64 {
	if id < 0 {
		return int64(1)
	}
	mn := size(id, ps)
	if id == ps {
		if val == ar[id] {
			return solution(id - 1, ps, val, ar)
		} else if val < ar[id] {
			return pw[mn]
		} else {
			return int64(0)
		}
	}  
	if ar[id] == 0 {
		return solution(id - 1, ps, val, ar)
	}
	return ar[id] * pw[mn] + solution(id - 1, ps, val, ar)
}

func calculateValue(n int64) (int64, int64) {
	i1 := int64(1)
	mul := int64(9)
	tc := int64(0)
	for {
		if n < mul * i1 {
			tc += n / i1 
			return tc, n % i1
		}
		n -= mul * i1
		tc += mul 
		mul *= 10
		i1++
	}
}

func solve() {
	var n int64
	Fscan(in, &n)
	
	num, mod := calculateValue(n)

	ar := sol(num)
	ans := int64(0)

	for i := int64(len(ar) - 1); i >= 0; i-- {
		for j := int64(1); j < 10 ; j++ {
			ans += solution(int64(len(ar) - 1), i, j, ar) * j
		}
	}

	if mod != 0 {
		ar1 := sol(num + 1)
		for mod > 0 {
			ans += ar1[len(ar1) - int(mod)]
			mod--
		}
	}

	Fprintln(out, ans)
}

func main() {
	defer out.Flush()

	var T int
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}

func init() {
	pw = make([]int64, 18)
	pw[0] = 1
	for i := 1; i < 18; i++ {
		pw[i] = pw[i-1] * 10
	}
}

/*

6
5
10
13
29
1000000000
1000000000000000

*/