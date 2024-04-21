package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	numbers := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}
	index := binarySearch(numbers, 30)

	if index != 3 {
		t.Fail()
	}
}
