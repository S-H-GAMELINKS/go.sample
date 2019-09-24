package main

import "fmt"

func pow(x int, c chan int) {
	c <- x * x
}

func main() {

	i := 42

	c := make(chan int)

	go pow(i, c)

	result := <-c

	fmt.Println(result)
}
