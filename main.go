package main

import (
	"fmt"
	"github.com/sumitsj/async-pattenrns/async"
	"github.com/sumitsj/async-pattenrns/printer"
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
