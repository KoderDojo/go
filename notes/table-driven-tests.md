# Table-Driven Tests in Go

I recently learned how to use the *testing* package in Go to unit test my code. In many instances I want to run the same test using different input values, but I don't want to create multiple tests and write a lot of code.

I found a nice article by Dave Cheney that discusses [Writing table driven tests in Go](https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go). This is exactly what I was looking for.

## Example

Let's say I have a type `Square` and I want to verify it is calculating the correct area.

```
package practice

type Square struct {
	Side int
}

func (z Square) Area() int {
	return z.Side * z.Side
}
```

Here is an example of writing a table-driven test. Notice the use of an anonymous *struct* that holds any inputs as well as the expected value.

I loop through the various tests one at a time, comparing the expected results with the actual results.

```
package practice_test

import (
	"practice"
	"testing"
)

func TestSquare(t *testing.T) {
	tests := []struct{
		side int
		areaExpected int
	}{
		{5, 25},
		{2, 4},
	}

	for _, test := range tests {
		s := practice.Square{test.side}
		if s.Area() != test.areaExpected {
			t.Errorf("Expected %v, Actual %v", test.areaExpected, s.Area())
		}
	}
}
```