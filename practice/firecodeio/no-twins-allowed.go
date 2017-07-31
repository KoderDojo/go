package challenge

func AreAllCharactersUnique(str string) int {
	chars := make(map[rune]struct{})
	var q struct{}

	for _, char := range str {
		if _, ok := chars[char]; ok {
			return 0
		}

		chars[char] = q
	}

	return 1
}

---

package challenge_test

import (
	"challenge"
	"testing"
)

func TestAreAllCharactersUnique(t *testing.T) {
	tests := []struct {
		value    string
		expected int
	}{
		{"hello", 0},
		{"print", 1},
		{"", 1},
		{"mississippi", 0},
	}

	for _, test := range tests {
		expected := test.expected
		actual := challenge.AreAllCharactersUnique(test.value)

		if expected != actual {
			t.Errorf("'%v' -> Expected: %v, Actual: %v", test.value, expected, actual)
		}
	}
}
