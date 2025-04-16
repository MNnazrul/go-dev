package main

import (
    "os"
    "bufio"
    ."fmt"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve() {
    var n int 
    Fscan(in, &n)
    gr := make([]int, n + 2)
    indeg := make([]int, n + 2)
    cnt := make([]int, n + 2)
    for i := 1; i <= n; i++ {
        Fscan(in, &gr[i])
        indeg[gr[i]]++
        cnt[i] = 1
    }
    queue := make([]int, 0)
    for i := 1; i <= n; i++ {
        if indeg[i] == 0 {
            queue = append(queue, i)
        }
    }
    ans := 2
    flg := false
    for len(queue) > 0 {
        nd := queue[0]
        // Fprint(out, nd, " ")
        queue = queue[1:]
        cnt[gr[nd]] += cnt[nd]
        indeg[gr[nd]]--
        if indeg[gr[nd]] == 0 {
            ans = max(ans, cnt[gr[nd]] + 1)
            queue = append(queue, gr[nd])
        }
        // ans++
        flg = true
    }
    if flg == true {
        ans++
    }
    // Fprintln(out, )
    Fprintln(out, ans)
}

func main() {
    defer out.Flush()
    var T int 
    for Fscan(in, &T); T > 0; T-- {
        solve()
    }
}