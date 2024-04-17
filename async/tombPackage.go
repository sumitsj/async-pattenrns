package async

import (
	"fmt"
	"gopkg.in/tomb.v1"
	"time"
)

func infiniteLoop3(tomb *tomb.Tomb) {
	defer tomb.Done()
	for {
		select {
		case <-tomb.Dying():
			fmt.Println("Tomb is dying")
			return
		default:
			fmt.Println("Running...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func RunWithTomb() {
	t := tomb.Tomb{}
	go infiniteLoop3(&t)
	fmt.Println("Go Routine started")
	time.Sleep(5 * time.Second)
	t.Kill(fmt.Errorf("terminate go routine"))
	err := t.Wait()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(1 * time.Second)
}
