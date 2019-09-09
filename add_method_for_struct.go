package main

import "fmt"

type Hoge struct {
    x, y int   
}


func (hoge Hoge) sum() (sum int) {
    sum = hoge.x + hoge.y
    return
}

func main() {
    hoge := Hoge{21, 21}
    
    fmt.Println(hoge.sum())
}
