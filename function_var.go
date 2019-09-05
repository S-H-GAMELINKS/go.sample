package main

import "fmt"

func main() {
    
    f := func(x, y int) (sum int) {
        sum = x + y
        return
    }
    
    fmt.Println(f(21, 21))
}
