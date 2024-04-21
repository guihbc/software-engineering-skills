package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	array := []int{5, 3, 2, 1, 4}
	expectedReturn := []int{1, 2, 3, 4, 5}
	sortedArray := quicksort(array)

	if !reflect.DeepEqual(sortedArray, expectedReturn) {
		t.Fail()
	}
}
