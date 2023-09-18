package main

import "sync"

func main() {
	// var beeper sync.WaitGroup
	// var evilNinjas = []string{"Tommny", "John", "Bob"}
	// beeper.Add(len(evilNinjas))

	// for _, evilNinja := range evilNinjas {
	// 	go attack(evilNinja, &beeper)
	// }

	// beeper.Wait()
	// fmt.Println("Mission Completed")

	var beeper sync.WaitGroup
	beeper.Add(2)
	beeper.Done()
	beeper.Done()
	beeper.Wait()
}

// func attack(evilNinja string, beeper *sync.WaitGroup) {
// 	fmt.Println("Attacked evil ninja: ", evilNinja)
// 	beeper.Done()
// }
