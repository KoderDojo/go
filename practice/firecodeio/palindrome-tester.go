package challenge

func IsStringPalindrome(str string) int {
	j := len(str) - 1

	for i := 0; i < j; i++ {
		if str[i] != str[j] {
			return 0
		}

		j--
	}

	return 1
}

---

package challenge_test

import (
	"challenge"
	"testing"
)

func TestIsStringPalindrome(t *testing.T) {
	tests := []struct {
		value    string
		expected int
	}{
		{"racecar", 1},
		{"M", 1},
		{"", 1},
		{"not a palindrome", 0},
		{"abca", 0},
		{"abba", 1},
	}

	for _, test := range tests {
		expected := test.expected
		actual := challenge.IsStringPalindrome(test.value)

		if expected != actual {
			t.Errorf("'%v' -> Expected: %v, Actual: %v", test.value, expected, actual)
		}
	}
}
