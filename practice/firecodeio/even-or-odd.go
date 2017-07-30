package practice

type ListNode struct {
	Value int
	Next *ListNode
}

func IsListEven(head *ListNode) bool {
	isEven := true

	currentNode := head
	for currentNode != nil {
		isEven = !isEven
		currentNode = currentNode.Next
	}

	return isEven
}

---

package practice_test

import (
	"testing"
	"practice"
)

func TestReturnsTrueWhenHeadIsNil(t *testing.T) {
	expected := true
	actual := practice.IsListEven(nil)

	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestReturnsFalseWithSingleListNode(t *testing.T) {
	head := practice.ListNode{ Value:6, Next:nil}

	expected := false
	actual := practice.IsListEven(&head)

	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
