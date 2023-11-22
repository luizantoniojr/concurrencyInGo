package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	type value struct {
		mu    sync.Mutex
		value int
	}

	var wg sync.WaitGroup

	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)

	//Coffman Conditions
	//Mutual Exclusion - printSum does require exclusive rights to a and b.
	//Hold and Wait - printSum holds a lock and waits for another (second printSum call).
	//No Preemption - printSum cannot be stopped.
	//Circular Wait - printSum waits for a lock that is held by another printSum call.
	//If we prevent one of the conditions from happening, we can prevent a deadlock.
	go printSum(&a, &b)
	go printSum(&b, &a)

	wg.Wait()
}
