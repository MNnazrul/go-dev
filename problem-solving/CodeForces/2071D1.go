package main

import (
    ."fmt"
    "os"
    "bufio"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func chk(ar []int, pos int64) int64 {
    one := 0
    n := int64(len(ar) - 1)

    for _, value := range ar {
        one += value
    }
    Fprintln(out, "pos", pos)
    if pos <= n {
        ans := int64(0)
        for idx, value := range ar {
            if(int64(idx) <= pos) {
                ans += int64(value)
            }
        }
        return ans
    }
    bt, ans := 0, int64(one)
    for idx, value := range ar {
        
        bt ^= value      

        if int64(2 * idx + 1) <= n {
            continue
        }

        if bt == 0 {
            continue
        }

        Fprintln(out, "idx", idx)
        
        l := int64(idx)
        for 2 * l <= pos {
            if 2 * l > n {
                ans++
            }
            if 2 * l + 1 <= pos {
                ans++
            }
            Fprintln(out, "l=", l, ", a=", ans)
            
            if 2 * l > n {
                l = 2 * l 
            } else {
                break
            }
        } 
    }

    return ans
}

func solve() {
    var n, l, r int64
    Fscan(in, &n, &l, &r)
    ar := make([]int, 2 * n + 1)
    for i := int64(1); i <= n; i++ {
        Fscan(in, &ar[i])
    }
    i1, bt := 0, 0
    for i := n + 1; i <= 2 * n; i++ {
        for int64(i1) <= i / 2 {
            bt ^= ar[i1]
            i1++
        }
        ar[i] = bt
    }
    Fprintln(out, ar)
    // Fprintln(out, chk(ar, l - 1))
    Fprintln(out, chk(ar, 30))
    Fprintln(out, "=======----------=========")
    // Fprintln(out, chk(ar, r))
    // Fprintln(out, chk(ar, r) - chk(ar, l - 1))
}

func main() {
    defer out.Flush()

    var T int 
    for Fscan(in, &T); T > 0; T-- {
        solve()
    }
}

/*
[0 1 0 1 0 1 0 0 0 0 1 1 1 1 1 1 1 1 1 1 0]
[0 1 0 1 0 1 0 0 0 0 1 1 1 1 1 1 1 1 1 1 0 0 1 1 0 0 1 1 0 0 1 1 0 0 1 1 0 0 1 1 1 1 1 1 0 0 1 1 1 1 1]
*/