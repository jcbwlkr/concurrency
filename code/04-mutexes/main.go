package main

import (
	"fmt"
	"sync"
)

const workers = 20000

var counter int
var mu sync.Mutex

func main() {
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			for i := 0; i < 2; i++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final counter", counter)
}
