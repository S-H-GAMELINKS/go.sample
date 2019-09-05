package main

import "fmt"

type Hoge struct {
    x, y int   
}

func main() {
    
    m := make(map[string]Hoge)
    
    m["hoge"] = Hoge{42, 42}
    
    fmt.Println(m["hoge"])
}
