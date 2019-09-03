package main

import "fmt"

func main() {
    var array [10]int
    
    for i := 0; i < 10; i++ {
        array[i] = i   
    }
    
    slice := array[0:3]
    
    fmt.Println(slice)
}
