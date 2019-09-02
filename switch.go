package main

import "fmt"

func main() {
    
    i := 42
    
    switch i {
        case 1:
            fmt.Println("1")
        case 2:
            fmt.Println("2")
        default:
            fmt.Println(i)
    }
}
