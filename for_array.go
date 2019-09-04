package main

import "fmt"

func main() {
    var container [10]int
    
    for i := 0; i < len(container); i++ {
        fmt.Println(container[i])   
    }
}
