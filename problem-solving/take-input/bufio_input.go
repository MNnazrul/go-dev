package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()


	var T, n int
	Fscan(in, &T)

	for ;T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		for i := range a {
			Fprint(out, a[i], " ")
		}
		Fprintln(out, )
	}

}