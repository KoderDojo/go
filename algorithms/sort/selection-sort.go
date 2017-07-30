package practice

func SelectionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	for i := 0; i < len(arr)-1; i++ {
		minIndex := i

		j := i + 1
		for j < len(arr) {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}

			j++
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

---

package practice_test

import (
	"practice"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		Values   []int
		Expected []int
	}{
		{Values: nil, Expected: nil},
		{Values: []int{}, Expected: []int{}},
		{Values: []int{1}, Expected: []int{1}},
		{Values: []int{2, 1}, Expected: []int{1, 2}},
		{Values: []int{3, 2, 1}, Expected: []int{1, 2, 3}},
		{Values: []int{4, 3, 2, 1, 0, -1}, Expected: []int{-1, 0, 1, 2, 3, 4}},
	}

	for _, test := range tests {
		practice.SelectionSort(test.Values)
		for i := 0; i < len(test.Values); i++ {
			if test.Values[i] != test.Expected[i] {
				t.Errorf("Expected: %v, Actual: %v", test.Expected, test.Values)
			}
		}
	}
}
