package main

import (
    ."fmt"
    "bufio"
    "os"
)

var (
    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)
    n, m int64
)

func solve() {
    Fscan(in, &n, &m)
    if n == m {
        Fprintln(out, -1)
        return
    }
    if n < m {
        n = m
    }
    Fprintln(out, (1 << 35) - n)
}

func main() {

    defer out.Flush()

    var T int
    for Fscan(in, &T); T > 0; T-- {
        solve()
    }

}