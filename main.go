package main

import "fmt"

// 1st for select loop pattern
func main() {
	charChannel := make(chan string, 3)
	//this is buffered channel async
	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		select {
		case charChannel <- s:
		}
	}

	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}

	//if unbuffered channel sync

}
