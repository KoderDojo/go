// use sleep to pause execution of main so
// other functions have a chance to run.
// sleep is a hack just for demonstration
// purposes.
package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello World!")
}

func main() {
	go hello()
	go func() {
		fmt.Println("Anonymous function.")
	}()

	// Pause execution using sleep
	fmt.Println("Start sleep.")
	time.Sleep(2 * time.Second)
	fmt.Println("End main.")
}
