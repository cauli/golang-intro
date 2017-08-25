package buffered

import (
	"fmt"
	"time"
)

// Stolen from https://gobyexample.com/channel-buffering
func Run() {
	done := make(chan bool, 1)

	go worker(done)

	<-done

	fmt.Println("Channel received a value")
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}
