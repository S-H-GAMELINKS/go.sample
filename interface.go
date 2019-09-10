package main

import "fmt"

type RubyFunc interface {
    hello()
}

type Ruby struct {}

func (r *Ruby) hello() {
    fmt.Println("Hello Ruby")   
}

func main() {
    var ruby Ruby
    
    ruby.hello()    
}
