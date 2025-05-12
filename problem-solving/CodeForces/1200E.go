package main 

import (
	."fmt"
	"os"
	"bufio"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

type Node struct {
	val1 int64 
	val2 int64
	const base[2] int64 = {1231, 1511}
	const mod[2] int64 = {1e9 + 7, 998244353}
}

type Hash struct {
	ar []int64
}

func (h *Hash) init(st string) {
	n := len(st)
	h.ar = make([]int64, n + 2)
	h.ar[0].val1 = 0
	h.ar[0].val2 = 0

}

func main() {
	defer out.Flush()




}
