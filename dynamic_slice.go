package main

import "fmt"

func main() {
    
    var slice []int
    
    for i := 0; i < 20; i++ {
        slice = append(slice, i)
    }
    
    fmt.Println(slice)
}
