package main

import "fmt"

type Hoge struct {
    x, y int   
}

func (h Hoge) sum() (sum int) {
    sum = h.x + h.y
    return 
}

func main() {
    
    hoge := Hoge{21, 21}
    
    fmt.Println(hoge.sum())
}
