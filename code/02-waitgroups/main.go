package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		sing("Jacob", 5)
		wg.Done()
	}()

	sing("Anna", 3)

	wg.Wait()
}

func sing(singer string, n int) {
	for i := 0; i < n; i++ {
		for _, line := range lyrics {
			fmt.Printf("[%5s %02d]: %s\n", singer, i+1, line)
			time.Sleep(300 * time.Millisecond)
		}
	}
}

var lyrics = []string{
	"Row, row, row your boat,",
	"Gently down the stream.",
	"Merrily, merrily, merrily, merrily,",
	"Life is but a dream.",
}
