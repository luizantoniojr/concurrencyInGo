package main

import (
	"fmt"
	"sync"
)

func main() {
	//Here we add a variable that will allo our code to synchronize access to the data variable.
	var memoryAccess sync.Mutex
	var data int

	go func() {
		//Here we declare that until we declare otherwise our
		//goroutine should have exclusive access to this memory.
		memoryAccess.Lock()
		data++
		//Here we declare that we are done with our exclusive access to this memory.
		memoryAccess.Unlock()
	}()

	//Here we declare that until we declare otherwise our
	//goroutine should have exclusive access to this memory.
	memoryAccess.Lock()
	if data == 0 {
		fmt.Printf("The value is %v.\n", data)
	} else {
		fmt.Printf("The value is %v.\n", data)
	}
	//Here we declare that we are done with our exclusive access to this memory.
	memoryAccess.Unlock()

	//When we use memory access synchronization we have some problems about performance.
	//The program will run slower because we have to wait for the goroutine to finish.
}
