package main

import "fmt"

func myFunction(x int, y string) (result int, txt1 string) {
    result = x + x
    txt1 = y + " world"
    return
}

func recursiveFunction(x int) int {
    if x == 11 {
        return 0
    }
    fmt.Println(x)
    return recursiveFunction(x + 1)
}

func main() {
    // var a complex128 = complex(5, 2)
    
    // fmt.Println(a)
    // _pile , name := myFunction(2, "hello")
    // fmt.Println( name, _pile)

    //---------------
    recursiveFunction(0)

}