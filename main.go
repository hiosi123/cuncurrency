package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	channel := make(chan string)

	go throwingNinjaStar(channel)

	for message := range channel {
		fmt.Println(message)
	}
}

func throwingNinjaStar(channel chan string) {
	rand.Seed(time.Now().UnixNano())
	numRounds := 3
	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("you scored: ", score)
	}
	close(channel)
}
