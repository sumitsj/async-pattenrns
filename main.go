package main

import (
	"fmt"
	"sync"
	"time"
)

func printCharacter(c string, times int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < times; i++ {
		fmt.Println(c)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Application started.")
	defer fmt.Println("Application ended.")

	var wg sync.WaitGroup
	wg.Add(3)

	go printCharacter("A", 10, &wg)
	go printCharacter("B", 10, &wg)
	go printCharacter("C", 10, &wg)

	wg.Wait()
}
