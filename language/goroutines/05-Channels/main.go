// Using channels to block execution
// of main function.
package main

import (
	"fmt"
	"time"
)

// Specifying a write only channel
func numbers(done chan<- bool) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println(i)
	}

	done <- true
}

func main() {
	// Channels are reference types like
	// slices and maps. Use make.
	done := make(chan bool)

	go numbers(done)
	go func() {
		for _, char := range "ABCDEFGHIJK" {
			time.Sleep(time.Second)
			fmt.Println(string(char))
		}

		done <- true // closure
	}()

	fmt.Println("Waiting for goroutines to finish.")

	// This is blocking. Waiting for 2
	// items to be added to channel.
	// Here we are throwing away the
	// values added to channel.
	<-done
	<-done

	fmt.Println("End main.")
}
