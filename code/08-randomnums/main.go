package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Serial program to generate 30 random numbers from 0-99.
func main() {
	rand.Seed(time.Now().UnixNano())
	var nums []int

	for i := 0; i < 30; i++ {
		n := rand.Intn(100)
		nums = append(nums, n)
	}

	fmt.Println("Count:", len(nums))
	fmt.Println(nums)
}
