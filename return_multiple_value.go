package main

import "fmt"

func add(x, y int) (int, int) {
    return x, y
}

func main() {
    
    x, y := add(1, 41)
    
    fmt.Println(x, y)
}
