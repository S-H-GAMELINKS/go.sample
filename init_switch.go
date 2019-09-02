package main

import "fmt"

func main() {
    switch i := 42; i {
        case 1:
            fmt.Println("1")
        case 2:
            fmt.Println("2")
        default:
            fmt.Println(i)
    }
}
