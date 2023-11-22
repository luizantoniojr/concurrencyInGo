package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()

	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	//tryDir allows a person to attempt to move in a direction and returns whether they were successful.
	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %v", dirName)

		//Increment the value of dir to indicate that a person is trying to move in a direction.
		//This will happen atomically, so that only one person can try to move in a direction at a time.
		atomic.AddInt32(dir, 1)

		//This simulates a constant cadence between people trying to move in a direction.
		takeStep()

		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprint(out, ". Success!")
			return true
		}
		takeStep()

		//Decrement the value of dir to indicate that a person is no longer trying to move in a direction.
		atomic.AddInt32(dir, -1)

		return false
	}

	var left, right int32

	tryLeft := func(out *bytes.Buffer) bool { return tryDir("left", &left, out) }

	tryRight := func(out *bytes.Buffer) bool { return tryDir("right", &right, out) }

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer

		defer func() {
			fmt.Println(out.String())
		}()

		defer walking.Done()
		fmt.Fprintf(&out, "%v is trying to scoot:", name)

		//This is an artificial limit to prevent the program from running forever.
		for i := 0; i < 5; i++ {

			//The person will attempt to step left and then right.
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}

		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
	}

	//This variable provides a way for the program to wait until
	//both people are either able to pass one another or give up.
	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()
}
