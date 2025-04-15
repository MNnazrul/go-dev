package main

import "fmt"

func solve() {
	var l, r, d, u int
	fmt.Scan(&l, &r, &d, &u)

	if l == r && d == u && u == l {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")
}

func main() {
	var t int
	fmt.Scan(&t)

	for t > 0 {
		t--
		solve()
	}

}