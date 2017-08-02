// One can see that the 2 goroutines
// run concurrently. The numbers and
// letters are interspersed. The use
// of Scanln is a hack to halt execution
// of main function.
package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println(i)
	}
}

func main() {
	go numbers()
	go func() {
		for _, char := range "ABCDEFGHIJK" {
			time.Sleep(time.Second)
			fmt.Println(string(char))
		}
	}()

	// Pause execution
	fmt.Println("Waiting for input.")
	fmt.Scanln()
	fmt.Println("End main.")
}
