package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Program to generate 30 random numbers using a set of worker goroutines.
// A dedicated goroutine will feed work to the workers to perform.
// The main goroutine will fan the results back in.
func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	// work is how we signal to the workers to generate more values.
	work := make(chan struct{})

	// results is where the workers send their generated values.
	results := make(chan int)

	// Kick off 8 goroutines to generate numbers.
	const workers = 8
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			for range work {
				n := rand.Intn(100)
				results <- n
			}
			wg.Done()
		}()
	}

	// This goroutine will signal the workers to perform their work. It closes
	// the channel when it's done so they know it's quitting time.
	go func() {
		for i := 0; i < 30; i++ {
			work <- struct{}{}
		}
		close(work)
	}()

	// This goroutine waits for all of the workers to finish their jobs so we
	// can close the results channel. This will tell the main goroutine below
	// that there will be no mroe results.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Here in the main goroutine we will receive from the results channel
	// until there are no more results coming.
	var nums []int
	for n := range results {
		nums = append(nums, n)
	}

	fmt.Println("Count:", len(nums))
	fmt.Println(nums)
}
