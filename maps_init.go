package main

import "fmt"

type Hoge struct {
    x, y int   
}

func main() {
    
    m := map[string]Hoge{
        "hoge": Hoge{
            42, 42,
        },
        "fuga": Hoge{
            21, 21,
        },
    }
    
    fmt.Println(m)
}
