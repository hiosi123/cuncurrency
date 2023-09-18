package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock   sync.Mutex
	rwLock sync.RWMutex
	count  int
)

func main() {
	// basics()
	readAndWrite()
}

func basics() {
	iterations := 1000
	for i := 0; i < iterations; i++ {
		go increament()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("resulted count is:", count)
}

func readAndWrite() {
	go read()
	go read()
	go write()

	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}

func read() {
	rwLock.RLock()
	defer rwLock.RUnlock()

	fmt.Println("read locking")
	time.Sleep(1 * time.Second)
	fmt.Println("reading unlocking")
}

func write() {
	rwLock.Lock()
	defer rwLock.Unlock()

	fmt.Println("write locking")
	time.Sleep(1 * time.Second)
	fmt.Println("write unlocking")
}

func increament() {
	lock.Lock()
	count++
	lock.Unlock()
}
