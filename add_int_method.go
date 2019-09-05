package main

import "fmt"

type MyInt int

func (mi MyInt) pow() MyInt {
    return mi * 2   
}

func main() {
    
    var i MyInt = 42
    
    fmt.Println(i.pow())
}
