package main

import (
	"fmt"
	"time"
)

// 1st for select loop pattern
// done channel

// what is goroutine leak? unintentionally leaving goroutine.
func main() {
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 3)

	close(done)
}

func doWork(done <-chan bool) {
	//when parent wants to close the child goroutine, it is the main reason to use for select
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("doing work")
		}
	}
}
