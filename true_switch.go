package main

import "fmt"

func main() {
    
    i := 42
    
    switch {
        case 1 == i:
            fmt.Println("1")
        case 2 == i:
            fmt.Println("2")
        default:
            fmt.Println(i)
    }
}
