package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Program to generate 30 random numbers using a 30 goroutines. There is no
// shared memory so no need for a mutex.
func main() {
	rand.Seed(time.Now().UnixNano())

	// results is where the goroutines will send random numbers
	results := make(chan int)

	const grs = 30

	// Kick off 30 goroutines to generate our numbers.
	for i := 0; i < grs; i++ {
		go func() {
			results <- rand.Intn(100)
		}()
	}

	// Receive from the channel 30 times.
	var nums []int
	for i := 0; i < grs; i++ {
		n := <-results
		nums = append(nums, n)
	}

	fmt.Println("Count:", len(nums))
	fmt.Println(nums)
}
