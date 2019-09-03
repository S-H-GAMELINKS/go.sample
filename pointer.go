package main

import "fmt"

func main() {
    i := 42
    ip := &i
    fmt.Println(*ip)
    *ip = 21
    fmt.Println(*ip)
}
