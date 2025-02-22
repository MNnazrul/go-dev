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

func sayHello(name string) {
    fmt.Println("Welcome to the golang course. ")
}

func application() {
    var name string
    fmt.Println("Enter your name - ")
    fmt.Scanln(&name)
    
    fmt.Println("---------" + name)
    
    var num1 int
    var num2 int
    
    fmt.Println("Enter first number - ")
    fmt.Scanln(&num1)
    fmt.Println("Enter second number - ")
    fmt.Scanln(&num2)
    
    sum := num1 + num2
    
    fmt.Println("Hello, ", name)
    fmt.Println("Summation = ", sum)
    
    fmt.Println("Thank you for using the application")
}

func main() {
    
    slice1 := []int{}

    slice1 = append(slice1, 1)
    fmt.Println(cap(slice1))
    slice1 = append(slice1, 1)
    fmt.Println(cap(slice1))
    slice1 = append(slice1, 1)
    fmt.Println(cap(slice1))


}