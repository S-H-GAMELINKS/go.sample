package main

import "fmt"

func main() {
    
    var inter interface{} 
    
    inter = 42
    
    fmt.Println(inter.(int))
}
