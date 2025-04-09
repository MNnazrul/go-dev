package main

import "fmt"

/*
- First order function cannot take any function as a input.
- Higher order function take fucntion as input
- Here , add is a first order order function
    on the other hand processOperation is a Higher order function.


Note : one thing to notice here that, Not only taking inputs.

Higher order function can take function as input or it can return function also.

Example :
    call() is also a higher order function.


Higher order function:
    1. parameter -> function
    2. return function
    3. both


In the processOperation() , op is a function that is also a parameter.
    Here, op is callback function.

** First class function => HOF



*/

func call() func (a int, b int) {
    return add
}

func processOperation(a int, b int, op func(x int, y int)) {
    op(a, b)
}

func add(a int, b int) {
    z := a + b
    fmt.Println(z)
}

func main() {
    processOperation(2, 5, add)

    sum := call() // function expression.

    sum(3, 5)
}

/*

1. Parameter vs argument
2. First order Function (works with object, property and relation)
  i. standard function or named function
  ii. anonymous function
  iii. IIFE
  iv. function expression

3. Higher order function or First Class Function (workds with object, property, relation and function)
4. Callback function  (function that is passed as parameter of HOF)
5. First class citizer => (variable assign data)

*/