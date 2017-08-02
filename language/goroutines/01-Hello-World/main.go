// goroutines are lightweight threads of execution.
// main is also a goroutine. It finishes before the
// other 2 functions get a chance to run, because all
// goroutines run concurrently.
package main

import "fmt"

func hello() {
	fmt.Println("Hello World!")
}

func main() {
	go hello()
	go func() {
		fmt.Println("Anonymous function.")
	}()
}
