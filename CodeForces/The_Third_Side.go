package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func solve() {
	var n int
	fmt.Fscanf(reader, "%d", &n)
	ans := -n + 1
	for n > 0 {
		n--
		var val int
		fmt.Fscanf(reader, "%d", &val)
		ans += val
	}
	
	fmt.Println(ans)
}

func main() {
	var t int
	fmt.Fscanf(reader, "%d", &t)
	for t > 0 {
		t--
		solve()
	}
}