package printer

import (
	"fmt"
	"sync"
	"time"
)

const messageCount = 5

func printCharacterA(wg *sync.WaitGroup, signalA chan int, signalC chan int) {
	defer wg.Done()
	for i := 0; i < messageCount; i++ {
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
	for i := 0; i < messageCount; i++ {
		<-signalA
		fmt.Println("B")
		time.Sleep(100 * time.Millisecond)
		signalB <- i
	}
}
func printCharacterC(wg *sync.WaitGroup, signalB chan int, signalC chan int) {
	defer wg.Done()
	for i := 0; i < messageCount; i++ {
		<-signalB
		fmt.Println("C")
		time.Sleep(100 * time.Millisecond)
		if i < messageCount-1 {
			signalC <- i
		}
	}
}

func Run() {
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
