package main

import "fmt"

func main() {
    container := []int{0, 1, 2, 3, 4, 5}
    
    fmt.Println(container)
    
    container[0] = 42
    
    fmt.Println(container)
}
