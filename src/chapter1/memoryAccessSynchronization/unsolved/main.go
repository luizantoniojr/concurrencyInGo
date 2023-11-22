package main

import (
	"fmt"
)

func main() {
	//There is a name for a section of the program that needs
	//exclusive access to a shared resource: critical section.
	//We have three critical sections in this program:
	var data int

	go func() {
		//1. The goroutine needs exclusive access to the data variable.
		data++
	}()

	//2. The if statement needs exclusive access to the data variable.
	if data == 0 {
		//3. The fmt.Printf function needs exclusive access to the data variable.
		fmt.Printf("The value is %v.\n", data)
	} else {
		//3. The fmt.Printf function needs exclusive access to the data variable.
		fmt.Printf("The value is %v.\n", data)
	}
}
