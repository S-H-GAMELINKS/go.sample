package main

import "fmt"

func deferprint() {
    defer fmt.Println("Hello")   
    fmt.Println("Go!")
}

func main() {
    deferprint()
}
