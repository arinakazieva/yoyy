package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		id := i
		go func() {
			fmt.Printf("Goroutine %d says hello\n", id)
		}()
	}

	time.Sleep(500 * time.Millisecond)

	fmt.Println("\n--- 10000 goroutines ---")

	for i := 0; i < 10000; i++ {
		id := i
		go func() {
			fmt.Printf("Goroutine %d says hello\n", id)
		}()
	}

	time.Sleep(1 * time.Second)
}
