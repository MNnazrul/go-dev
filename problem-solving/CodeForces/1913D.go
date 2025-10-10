package main

import (
	"bufio"
	. "fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	mod = int(998244353)
)

type TestCase struct {
	n, ans  int
	ar      []int
	minLeft []int
}

func NewTest() *TestCase {
	return &TestCase{}
}

func (t *TestCase) takeInput() {
	Fscan(in, &t.n)
	t.ar = make([]int, t.n)
	t.minLeft = make([]int, t.n*4+5)
	for i := 0; i < t.n; i++ {
		Fscan(in, &t.ar[i])
	}
}

func (t *TestCase) sol() {

}

func (t *TestCase) produceAnswer() {

}

func (t *TestCase) showAnswer() {
	Fprintln(out, t.ans)
}

func solve() {

}

func main() {
	defer out.Flush()

	var T int
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}
