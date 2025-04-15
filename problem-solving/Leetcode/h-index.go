package main

import "fmt"

func hIndex(citations []int) int {
	l, r, ans := 0, len(citations), 0

	for l <= r {
		md := (l + r) / 2
		cnt := 0
		for _, value := range citations {
			if value >= md {
				cnt += 1
			}
		}
		if cnt >= md {
			l = md + 1
			if md > ans {
				ans = md
			}
		} else {
			r = md - 1
		}
	}

	return ans
}

func main() {
	fmt.Println(hIndex([]int{3,0,6,1,5}))
}