package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0

	doSlowTask(&counter)

	// for i := 0; i < 1000; i++ {
	// 	doSlowTask(&counter)
	// }

	// go doSlowTask(&counter)

	// for i := 0; i < 1000; i++ {
	// 	go doSlowTask(&counter)
	// }

	fmt.Println("Finished?")

	// loopForever()
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
