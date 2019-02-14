package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Number of goroutines: %d\n\n", runtime.NumGoroutine())

	numberGenLeaky()
	numberGenLeaky()
	numberGenLeaky()

	// Give the goroutines a chance to finish before reporting the new number.
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("\nNumber of goroutines: %d\n", runtime.NumGoroutine())
}

// numberGenLeaky gives a random number from 1 to 200 that is divisible by 13. It
// uses unnecessary concurrency to accomplish this goal. It leaks a goroutine
// every time it is called.
func numberGenLeaky() int {
	nums := make(chan int)

	// Spawn a goroutine to continuously generate numbers from 1 to 200. Send it
	// on the channel if possible or abort if the context is canceled.
	go func() {
		for {
			nums <- rand.Intn(200) + 1
		}
	}()

	// Read from our channel until we find a suitable number.
	for n := range nums {
		if n%13 == 0 {
			fmt.Printf("%d is divisible by 13!\n", n)
			return n
		}
	}

	return 0
}

// numberGen is an improved version of numberGenLeaky. It closes its spawned
// goroutine as it returns.
func numberGen() int {
	nums := make(chan int)
	cancel := make(chan struct{})

	// Spawn a goroutine to continuously generate numbers from 1 to 200. Send it
	// on the channel if possible or abort if the context is canceled.
	go func() {
		for {
			n := rand.Intn(200) + 1
			select {
			case <-cancel:
				// Close the goroutine.
				return
			case nums <- n:
				// Send the value.
			}
		}
	}()

	// Read from our channel until we find a suitable number.
	for n := range nums {
		if n%13 == 0 {
			fmt.Printf("%d is divisible by 13!\n", n)
			close(cancel)
			return n
		}
	}

	return 0
}

func workWithTimeout() {
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
