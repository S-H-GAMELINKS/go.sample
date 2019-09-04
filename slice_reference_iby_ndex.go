package main

import "fmt"

func main() {
    var container [10]int
    
    for i := 0; i < 10; i++ {
        container[i] = i   
    }
    
    slice := container[:]
    
    fmt.Println(slice)
    
    slice = container[5:]
    
    fmt.Println(slice)
    
    slice = container[:9]
    
    fmt.Println(slice)
}
