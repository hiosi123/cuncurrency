package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	call := make(chan bool, 4)

	evilNinjas := []string{"Tommny", "Johny", "Bobby", "Andy"}

	for _, evilNinja := range evilNinjas {
		go Attack(evilNinja, call)
	}

	fmt.Println(<-call)

}

func Attack(target string, call chan bool) {
	fmt.Println("Throwing ninja Sword at ", target)
	time.Sleep(time.Second)
	call <- true
}
