package main

import (
	"fmt"
)

func solve() {
	var n int
	fmt.Scan(&n)
	ans := -n + 1
	for n > 0 {
		n--
		var val int
		fmt.Scan(&val)
		ans += val
	}

	fmt.Println(ans)
}

func main() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		t--
		solve()
	}
}