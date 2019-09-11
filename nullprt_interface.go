package main

import "fmt"

type MyInterface interface {
    hello()   
}

type Hoge struct {
    str string
}

func (h *Hoge) hello() {
    if h == nil {
        fmt.Println("nil!")   
        return
    }
    
    fmt.Println("Hello", h.str)   
}

func main() {
    hoge := Hoge{"go"} 
    
    hoge.hello()
    
    var h *Hoge
    
    h.hello()
}
