package main

import "fmt"

func add(a, b int) int {
return a + b
}

func sub(a, b int) int {
return a - b
}

func mul(a, b int) int {
return a * b
}

func divide(a, b int) int {
return a / b
}

func main() {
var a, b int

fmt.Printf("Enter values for a and b: ")
fmt.Scanln(&a, &b)
fmt.Printlf("a = %d and b = %d", a, b)

fmt.Printf("a + b = %d\n", mul(a, b))
fmt.Printf("a - b = %d\n", divide(a, b))
}
