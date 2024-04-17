package main

import (
	"fmt"
	"github.com/sumitsj/channel-synchronizer/printer"
	"github.com/sumitsj/channel-synchronizer/terminator"
)

func main() {
	fmt.Println("Application started.")
	defer fmt.Println("Application ended.")

	printer.Run()
	terminator.Run()
	terminator.RunWithContext()
	terminator.RunWithTomb()
}
