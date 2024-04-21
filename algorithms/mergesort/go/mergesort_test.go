package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	array := []int{5, 3, 2, 1, 4}
	expectedReturn := []int{1, 2, 3, 4, 5}
	sortedArray := mergeSort(array)

	if !reflect.DeepEqual(sortedArray, expectedReturn) {
		t.Fail()
	}
}
