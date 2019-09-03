package main

import "fmt"

func main() {
    var array [3]int
    
    for i := 0; i < 3; i++ {
        array[i] = 0   
    }
    
    for a := range array {
        fmt.Println(a)   
    }
}
