package main

import "fmt"

func main() {
    array := make([]int, 10)
    
    fmt.Println(array, len(array), cap(array))
}
