package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	work()
}

func work() {
	ch := make(chan string)

	go func() {
		delay := time.Duration(rand.Intn(200)) * time.Millisecond
		time.Sleep(delay)
		ch <- "some value"
	}()

	timeout := time.After(100 * time.Millisecond)

	select {
	case val := <-ch:
		fmt.Println("Worker gave me:", val)
	case <-timeout:
		fmt.Println("Took too long! Moving on!")
	}
}
