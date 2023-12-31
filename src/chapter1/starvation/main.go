package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second

	greedyWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {

			// The greedy worker will lock the sharedLock and hold it for 3 nanoseconds.
			sharedLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}

		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	}

	politeWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {

			//The polite worker will lock the sharedLock only when it needs to.
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			count++
		}

		fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
	}

	//This setup indicate that there are two goroutines to wait for.
	wg.Add(2)

	go greedyWorker()
	go politeWorker()

	wg.Wait()

	//One of the ways to detect and solve starvation is by logging whe work is accomplished.
	//Starvation cal also apply to CPU, memory, file handles, database connections
	//and other resources, any shared resource can be a source of starvation.
}
