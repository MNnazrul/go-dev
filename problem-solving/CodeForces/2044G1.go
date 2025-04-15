// link : https://codeforces.com/problemset/problem/2044/G1

package main

import (
    ."fmt"
    "os"
    "bufio"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve() {
    var n int
    Fscan(in, &n)
    ar := make([]int, n)
    cnt := make([]int, n + 2)
    for i := 0; i < n; i++ {
        Fscan(in, &ar[i])
        cnt[ar[i]]++
    }
    var ar1 []int 
    for i := 1; i <= n; i++ {
        if cnt[i] == 0 {
            ar1 = append(ar1, i)
        }
    }
    ans := 2
    for len(ar1) > 0 {
        var ar2 []int
        for _, id := range ar1 {
            cnt[ar[id - 1]]--
            if cnt[ar[id - 1]] == 0 {
                ar2 = append(ar2, ar[id - 1])
            }
        }
        ar1 = ar2
        ans++
    }
    Fprintln(out, ans)
}

func main() {
    defer out.Flush()
    var T int
    for Fscan(in, &T); T > 0; T-- {
        solve()
    }
}
