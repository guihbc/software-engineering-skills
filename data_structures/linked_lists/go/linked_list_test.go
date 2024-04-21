package main

import (
	"testing"
)

func TestLinkedListInsert(t *testing.T) {
	var list *LinkedList = NewLikedList()
	list.insert(1)
	list.insert(2)
	list.insert(3)

	first := list.first
	if first.element != 3 {
		t.Fail()
	}

	second := first.next
	if second.element != 2 {
		t.Fail()
	}

	third := second.next
	if third.element != 1 {
		t.Fail()
	}
}

func TestLinkedListRemove(t *testing.T) {
	var list *LinkedList = NewLikedList()
	list.insert(1)
	list.remove(1)

	if !list.first.isEmpty() {
		t.Fail()
	}
}
