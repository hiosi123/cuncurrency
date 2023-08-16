package main

import "fmt"

func main() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	go func() {
		anotherChannel <- "mew"
	}()

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannle := <-anotherChannel:
		fmt.Println(msgFromAnotherChannle)
	}
}
