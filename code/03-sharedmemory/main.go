package main

import (
	"fmt"
	"sync"
)

const workers = 2000

var counter int

func main() {
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			for i := 0; i < 2; i++ {
				counter++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final counter", counter)
}
