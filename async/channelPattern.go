package async

import (
	"fmt"
	"time"
)

func infiniteLoop(quit *chan bool) {
	for {
		select {
		case <-*quit:
			fmt.Println("Channel closed...")
			return
		default:
			fmt.Println("Running...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func Run() {
	quit := make(chan bool)
	go infiniteLoop(&quit)
	fmt.Println("Go Routine started")
	time.Sleep(5 * time.Second)
	quit <- true
	time.Sleep(1 * time.Second)
}
