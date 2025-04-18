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
    cnt1, cnt2 := 0, 0
    odd, even := 0, 0
    mx := 0
    for i := 0; i < n; i++ {
        Fscan(in, &ar[i])
        if ar[i] > mx {
            mx = ar[i]
        }
        if i < n - 1 {
            if ar[i] % 2 == 1 {
                odd += ar[i]
                cnt2++
            } else {
                even += ar[i]
                cnt1++
            }
        }
    }

    if ar[n-1] % 2 == 1 && cnt1 > 0 {
        ar[n-1] += even
        if cnt2 > 0 {
            ar[n-1]--
            ar[n-1] += odd - (cnt2 - 1)
        }
    } else if ar[n-1] % 2 == 0 && cnt2 > 0 {
        ar[n-1] += odd  - (cnt2 - 1)
        if cnt1 > 0 {
            ar[n-1] += even
        }
    }

    if mx > ar[n-1] {
        ar[n-1] = mx
    }

    Fprintln(out, ar[n-1])
}

func main() {

    defer out.Flush()

    var T int
    for Fscan(in, &T); T > 0; T-- {
        solve()
    }
}