package main

import (
	"fmt"
	"github.com/sumitsj/channel-synchronizer/async"
	"github.com/sumitsj/channel-synchronizer/printer"
)

func main() {
	fmt.Println("Application started.")
	defer fmt.Println("Application ended.")

	printer.Run()
	async.Run()
	async.RunWithContext()
	async.RunWithTomb()
	async.RunWithChannel()
}
