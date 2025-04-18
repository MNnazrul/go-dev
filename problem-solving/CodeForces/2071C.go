package main

import (
    ."fmt"
    "os"
    "bufio"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func bfs() {

}

func solve() {
    var n, en, st int 
    Fscan(in, &n, &st, &en)
    gr := make([][]int, n + 1)
    for i := 0; i < n - 1; i++ {
        var u, v int 
        Fscan(in, &v, &u)
        gr[u] = append(gr[u], v)
        gr[v] = append(gr[v], u)
    }

    clr := make([]bool, n + 2)
    queue := []int{en}
    clr[en] = true
    id := 0

    for id < len(queue) {
        nd := queue[id]
        id++
        for _, child := range gr[nd] {
            if clr[child] == false {
                clr[child] = true
                queue = append(queue, child)
            }
        }
    }

    for i := len(queue) - 1 ; i >= 0; i-- {
        Fprint(out, queue[i], " ")
    }
    Fprintln(out, )

}

func main() {
    defer out.Flush()

    var T int 
    for Fscan(in, &T); T > 0; T-- {
        solve()
    }
}