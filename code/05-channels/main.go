package main

import (
	"fmt"
)

func main() {
	input := make(chan int)
	output := make(chan int)

	go func() {
		n := <-input
		n = n * 3
		output <- n
	}()

	input <- 14
	answer := <-output

	fmt.Println("Final answer", answer)
}
