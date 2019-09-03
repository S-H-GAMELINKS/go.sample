package main

import "fmt"

func main() {
    var array [10]int
    
    for i := 0; i < 10; i++ {
        array[i] = i   
    }
    
    slice := array[0:3]
    
    slice[0] = 42
    
    fmt.Println(slice)
    fmt.Println(array)
}
