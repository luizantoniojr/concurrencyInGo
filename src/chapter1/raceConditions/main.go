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
}
