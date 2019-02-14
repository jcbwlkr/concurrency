package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Program to generate 30 random numbers using 30 goroutines.
func main() {
	rand.Seed(time.Now().UnixNano())

	var nums []int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			n := rand.Intn(100)
			mu.Lock()
			nums = append(nums, n)
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Count:", len(nums))
	fmt.Println(nums)
}
