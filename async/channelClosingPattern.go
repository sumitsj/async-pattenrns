package async

import (
	"fmt"
	"time"
)

func infiniteLoop4(dataChan *chan string) {
	for {
		data, ok := <-*dataChan
		if !ok {
			fmt.Println("Channel closed...")
			return
		}

		fmt.Println("Running...", data)
		time.Sleep(500 * time.Millisecond)
	}
}

func RunWithChannel() {
	dataChan := make(chan string)
	go infiniteLoop4(&dataChan)
	fmt.Println("Go Routine started")
	dataChan <- "A"
	dataChan <- "B"
	dataChan <- "C"
	dataChan <- "D"
	dataChan <- "E"
	close(dataChan)
}
