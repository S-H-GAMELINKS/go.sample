package main

import "fmt"

func hello() {
    fmt.Println("Hello ")   
}

func fmtgo() {
    fmt.Println("Go")   
}

func main() {
    defer fmtgo()
    defer hello()
}
