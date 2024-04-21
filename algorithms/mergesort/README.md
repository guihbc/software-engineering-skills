# Mergesort

Just like Quicksort, Mergesort is a sorting algorithm that utilizes the [divide and conquer](https://github.com/guihcodes/software-engineering-skills/tree/develop/algorithms/divide_and_conquer) paradigm. This approach involves repeatedly breaking down the input to divide the problem into smaller subproblems.

### Merge

Mergesort divides the list into two parts and uses recursion to perform the same function on both halves. This process continues until each part consists of only one element from the original list. It then compares the values between each part, placing the smaller value on the left. Subsequently, it merges these parts, reassembling the list into larger ones and repeating the comparison and merge process on both formed lists until the entire list is ordered.

To better understand how this algorithm works, refer to the image below.

![002_mergesort](https://user-images.githubusercontent.com/48635609/91216255-0643d300-e6ec-11ea-82ba-2b8af5fe5b42.gif)

The gif illustrates the process. The list initially undergoes a "top-down" separation of elements until there are multiple lists, each with a single element from the main list. When the process reaches the end, the list reverses direction, this time performing comparisons between the various parts of the list to arrange them in order and then merging them, sorting the sublists until they reach their original size and the entire list is sorted.

![001_mergesort](https://user-images.githubusercontent.com/48635609/91215813-638b5480-e6eb-11ea-82ca-fa414d5a5d2f.png)

Below is the implementation of the `mergesort` algorithm in Go:

```go
package main

import (
	"fmt"
)

func main() {
	array := []int{5, 3, 2, 1, 4}
	fmt.Println("\n--- Unsorted --- \n\n", array)
	fmt.Println("\n--- Sorted ---\n\n", mergeSort(array), "\n")
}

func mergeSort(a []int) []int {
	var num = len(a)

	if num == 1 {
		return a
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)

	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = a[i]
		} else {
			right[i-middle] = a[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
```

*Note: The code above in the "mergesort.go" file includes explanatory comments for each line.*

### References

- [Golang Programs - Mergesort Implementation](https://www.golangprograms.com/golang-program-for-implementation-of-mergesort.html)

- [GeeksforGeeks - Merge Sort](https://www.geeksforgeeks.org/merge-sort/)

### Supplementary Materials

- [MERGE SORT | Algoritmos #7](https://www.youtube.com/watch?v=5prE6Mz8Vh0) - video

- [Sorting Algorithm - MergeSort Video Lecture](https://www.youtube.com/watch?v=PKCMMSXQyJE) - video