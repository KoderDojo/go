/* ***********************************************************************
 *
 *                           Bubble Sort
 *
 *  Except for the classic 'hello world!' application, this is the first
 *  application I have written in Go.
 *
 *  Bubble sort is not useful for production purposes, but it is a useful
 *  example in the classroom for teaching sorting algorithms.
 *
 *  Bubble sort makes multiple passes through the array, bubbling the
 *  next highest value to the next top-most position by continually
 *  comparing neigboring values in the array.
 *
 *  Runs in O(n^2).
 *
 * ******************************************************************** */

package main

import "fmt"

func bubbleSort(array []int) {
	j := len(array) - 1

	for j > 0 {
		swap := false

		for i := 0; i < j; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				swap = true
			}
		}

		if !swap {
			break
		}

		j--
	}
}

func main() {
	// This is technically a slice.
	var array = []int{7, 2, 1, 0, 10, 8, 6, 5, 3, 4, 9}

	bubbleSort(array)

	for i := 0; i < len(array); i++ {
		fmt.Printf("% d", array[i])
	}
}
