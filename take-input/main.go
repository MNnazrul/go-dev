package main

import (
	"fmt"
	"os"
)


func main() {
    file, err := os.Open("input.txt")

    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()


    var t, val, val1 int
    var str string 

    fmt.Fscanf(file, "%d", &t)

    fmt.Println(t)
    for ; t > 0; t-- {
        fmt.Fscanf(file, "%s %d", &str, &val)
        fmt.Println(str, val)

        for ; val > 0; val-- {
            fmt.Fscanf(file, "%d", &val1)
            fmt.Print(val1, " ")
        }
        fmt.Println()

    }


    // reader := bufio.NewReader(file)



    // for {
    //     line, err := reader.ReadString('\n')
    //     if err != nil {
    //         fmt.Println(err)
    //         return
    //     }

    //     results := strings.TrimSpace(line)
    //     fmt.Println("=> ", results[0])
    // }

    // data := make([]byte, 100)

    // _, err = reader.Read(data)

    // if err != nil {
    //     fmt.Println(err)
    //     return
    // }

    // fmt.Println(string(data))
}
