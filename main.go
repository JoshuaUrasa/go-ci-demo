package main

import "fmt"

func Add(a, b int) int {
    return a + b
}

func main() {
    fmt.Println("Automatic build triggered from Git")
    fmt.Println("2 + 3 =", Add(2, 3))
}
