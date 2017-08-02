// Using a WaitGroup to pause the
// execution of main until the
// goroutines finish running.
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	w sync.WaitGroup
)

func numbers() {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println(i)
	}
	w.Done()
}

func main() {
	w.Add(2)

	go numbers()
	go func() {
		for _, char := range "ABCDEFGHIJK" {
			time.Sleep(time.Second)
			fmt.Println(string(char))
		}
		w.Done()
	}()

	fmt.Println("Waiting for goroutines to finish.")
	w.Wait()
	fmt.Println("End main.")
}
