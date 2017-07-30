package practice


func ReverseString(str string) string {
	newString := ""

	for _, char := range str {
		newString = string(char) + newString
	}

	return newString
}

---

package practice_test

import (
	"testing"
	"practice"
)


func TestReverseString(t *testing.T) {
	tests := []struct {
		str string
		expected string
	}{
		{ "Hello", "olleH" },
		{ "abcd", "dcba" },
		{ "a", "a" },
	}

	for _, test := range tests {
		actual := practice.ReverseString(test.str)
		if actual != test.expected {
			t.Errorf("Expected %v, Actual %v", test.expected, actual)
		}
	}
}
