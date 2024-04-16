package main

import (
	"fmt"
	"sync"
	"time"
)

func printCharacterA(wg *sync.WaitGroup, signalA chan int, signalC chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		if i > 0 {
			<-signalC
		}
		fmt.Println("A")
		time.Sleep(100 * time.Millisecond)
		signalA <- i
	}
}
func printCharacterB(wg *sync.WaitGroup, signalA chan int, signalB chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-signalA
		fmt.Println("B")
		time.Sleep(100 * time.Millisecond)
		signalB <- i
	}
}
func printCharacterC(wg *sync.WaitGroup, signalB chan int, signalC chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-signalB
		fmt.Println("C")
		time.Sleep(100 * time.Millisecond)
		if i < 9 {
			signalC <- i
		}
	}
}

func main() {
	fmt.Println("Application started.")
	defer fmt.Println("Application ended.")

	var wg sync.WaitGroup
	wg.Add(3)

	chanA := make(chan int)
	chanB := make(chan int)
	chanC := make(chan int)

	go printCharacterA(&wg, chanA, chanC)
	go printCharacterB(&wg, chanA, chanB)
	go printCharacterC(&wg, chanB, chanC)

	wg.Wait()
}
