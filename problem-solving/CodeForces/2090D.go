package main

import (
    ."fmt"
    "os"
    "bufio"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func is_prime(n int) bool {

    for i := int64(2); i * i <= int64(n); i++ {
        if int64(n) % i == 0 {
            return false
        }
    }

    return true
}

func solve() {

    var n int 
    Fscan(in, &n)
    pr := -1 
    for i := n; i >= 1; i-- {
        if is_prime((i + 2) / 2) {
            pr = i
            break
        }
    }

    mp := make([]bool, n + 1)
    ans := make([]int, 0)
    for i := 1; i <= n; i++ {
        mp[i] = true
    }

    l, r := 1, pr 
    for l < r {
        ans = append(ans, l)
        ans = append(ans, r)
        mp[l] = false ; mp[r] = false
        l++
        r--
    }

    for i := 1; i <= n; i++ {
        if mp[i] {
            ans = append(ans, i)
        }
    }

    for _, value := range ans {
        Fprint(out, value, " ")
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