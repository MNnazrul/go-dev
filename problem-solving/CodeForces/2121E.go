package main

import (
	"bufio"
	. "fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}


func makeArray(a int) []int {
	ans := make([]int, 0)
	for a > 0 {
		ans = append(ans, a % 10)
		a /= 10
	}
	reverseSlice(ans)
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a 
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a 
	}
	return b 
}

func sol(id int, prev int, ar []int, br []int, dp [][]int) int {
	if id == len(ar) {
		return 0 
	}
	ret := dp[prev][id]
	if ret != -1 {
		return ret 
	}
	ret = 0 
	i := id
	dp[prev][id] = ret
	if br[id] - ar[id] > 1 {
		return 0
	}

	if prev == 0 {
		if ar[i] == br[i] {
			ret = sol(id + 1, 0, ar, br, dp) + 2
		} else {
			ret = 1 + min(sol(id + 1, 1, ar, br, dp), sol(id + 1, 2, ar, br, dp))
		}
	} else if prev == 1 {
		if ar[i] != 9 {
			if br[i] == 9 {
				dp[prev][id] = 1
				return 1;
			}
			return 0;
		}
		if ar[i] == 9 {
			ret = sol(id + 1, prev, ar, br, dp) + 1
		}
		if br[i] == 9 {
			ret++
		}
	} else {
		if br[i] != 0 {
			if ar[i] == 0 {
				dp[prev][id] = 1
				return 1;
			}
			return 0;
		}
		if br[i] == 0 {
			ret = sol(id + 1, prev, ar, br, dp) + 1
		}
		if ar[i] == 0 {
			ret++
		}
	}
	dp[prev][id] = ret
	return ret 
}

func solve() {
	var a, b int
	Fscan(in, &a, &b)
	ar := makeArray(a)
	br := makeArray(b)
	dp := make([][]int, 5)

	for i, _ := range dp {
		dp[i] = make([]int, len(ar) + 2)
		for j, _ := range dp[i] {
			dp[i][j] = -1
		}
	}

	Fprintln(out, sol(0, 0, ar, br, dp))

}

func main() {
	defer out.Flush()

	var T int
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}
