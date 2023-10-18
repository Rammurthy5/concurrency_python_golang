package main

import (
	"fmt"
	"sync"
	"time"
)

// Function that simulates a task
func task(number int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	time.Sleep(2 * time.Second) // Simulate a time-consuming task (2 seconds)
	results <- fmt.Sprintf("Task %d result", number)
}

func main() {
	startTime := time.Now()
        var wg sync.WaitGroup
	results := make(chan string, 10000000)

	// Run 1000 tasks concurrently
	for i := 1; i <= 10000000; i++ {
		wg.Add(1)
		go task(i, &wg, results)
	}

	// Wait for all tasks to complete
	wg.Wait()
	close(results)

        elapsedTime := time.Since(startTime) // Calculate the total execution time
	fmt.Printf("Total execution time: %s\n", elapsedTime)
	fmt.Println("All tasks completed.")
}

