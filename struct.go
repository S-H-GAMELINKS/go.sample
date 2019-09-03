package main

import "fmt"

type Hoge struct {
    X int   
    Y int
}

func main() {
    hoge := Hoge{1, 3}
    
    fmt.Println(hoge.X, hoge.Y)
}
