// Concurrency is composition of independently executing processes while
// paralelism is the simultaneous execution of computations.
// Goroutine: abstraction layer over Threads in Go (i.e. does not use directly threads for concurrency).
// Goroutines are lighter weight and managed by the language, and have less switching between threads
// Multiple goroutines can be assigned to one single thread!
// Go's concurrency model: actor model, using CSP (Communicating Sequential Processes through channels).
// Goroutines work as management service, avoiding the burden of switching in threads

package main

import (
	"fmt"
	"runtime" // Add parallel programming
	"sync"
	"time"
)

func main() {
	// Adding 2 threads as max processes to the main program, allowing parallel programming
	// runtime.GOMAXPROCS(2)

	// Go as goroutine needs syncronization so that the main function wait for concurrent functions to be finished
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		defer waitGroup.Done()
		fmt.Println("Pluralsight")
	}()

	waitGroup.Wait()

	// Channels: async (buffered) vs sync (un-buffered)
	myChannelAsync := make(chan int, 5)
	myChannelSync := make(chan int)
}
