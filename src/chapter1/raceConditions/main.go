package main

import (
	"fmt"
)

func main() {
	var data int

	go func() {
		data++
	}()

	//There is no guarantee that the goroutine will run before the main function
	if data == 0 {
		fmt.Printf("The value is %v.\n", data)
	}

	//Possible outcomes:
	//1. Nothing is printed. In this case, line 11 is executed before line 16.
	//2. The value is 0. In this case, lines 15 and 16 is executed before line 11.
	//3. The value is 1. In this case, line 15 is executed before line 11, but the goroutine is executed before line 16.
}
