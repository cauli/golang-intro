package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	counter := 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			doSlowTask(&counter)
			wg.Done()
		}()
	}

	fmt.Println("Finished?")

	wg.Wait()

	fmt.Println("Really finished!")

	loopForever()
}

func doSlowTask(counter *int) {
	fmt.Println("Starting slow task...")
	time.Sleep(1 * time.Second)

	*counter++
	fmt.Println("Done! Count: ", *counter)
}

func loopForever() {
	// Infinite loop :)
	for {
	}
}
