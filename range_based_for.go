package main

import "fmt"

func main() {
    
    var slice [10]int
    
    fmt.Println(slice)
    
    for i, v := range slice {
        fmt.Println(i, v)
    }    
}
