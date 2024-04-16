package main

import (
	"fmt"
	"time"
)

func printCharacter(c string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(c)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Application started.")
	go printCharacter("A", 10)
	go printCharacter("B", 10)
	go printCharacter("C", 10)
	time.Sleep(2 * time.Second)
	defer fmt.Println("Application ended.")
}
