package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	tasks := make(chan func(), 64)

	// spawn four worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for cmd := range tasks {
				cmd()
			}
			wg.Done()
		}()
	}

	// generate some tasks
	for i := 0; i < 10; i++ {
		tasks <- func() { fmt.Println("'Hello from iteration n." + strconv.Itoa(i) + "'") }
	}
	close(tasks)

	// wait for the workers to finish
	wg.Wait()
}
