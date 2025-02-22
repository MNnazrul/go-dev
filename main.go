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

func slice_tutorial() {
    slice1 := []int{}
    
    slice1 = append(slice1, 1)
    fmt.Println(cap(slice1))
    slice1 = append(slice1, 1)
    fmt.Println(cap(slice1))
    slice1 = append(slice1, 1)
    fmt.Println(cap(slice1))
}

func map_tutorial() {

    var a = make(map[string]string) // empty map
    var b map[string]string // empty map

    fmt.Println(a == nil)
    fmt.Println(b == nil)

    menu := map[string]float64 {
        "soup" : 4.99,
        "pie" : 3.4,
        "salad" : 6.99,
    }

    fmt.Println(menu)

    fmt.Println(menu["pie"])

    // looping maps
    for k, v := range menu {
        fmt.Println(k, "-", v)
    }

    //ints as key type

}

type Person struct {
   name string
   age int
   job string
   salary int
}

func print_person(person Person) {
    person.name = "Nazrul Islam"
    fmt.Println(person)
}

func class_tutorial() {
    var par1 Person
    par1.name = "nazrul"
    par1.age = 24
    // par1.job = "software engineer"
    // par1.salary = 40000
    print_person(par1)
    fmt.Println(par1.name)
}

func main() {
    
    class_tutorial()


}