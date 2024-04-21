# Quicksort

Quicksort is a sorting algorithm that employs the concept of sorting by exchange, which arranges elements in ascending or descending order, or any other desired order.

This algorithm utilizes the [divide and conquer](https://github.com/GuilhermehChaves/google-skills/tree/master/algoritmos/divide_and_conquer) paradigm, which involves recursively breaking down the input to divide the problem into smaller subproblems.

#### Partition

In the partition process, a pivot variable is defined. All elements smaller than the pivot are placed to its left. The pivot can be either the last or first element and can be chosen randomly.

If we consider the sequence: [5, 3, 2, 1, 4], the pivot will be the number 4. We traverse the entire array, and when a value is smaller than the pivot, we move it to the left. To keep track of the elements that have been swapped, we use a variable `i`. At the end of the traversal, we replace the pivot with the value at position `i`. This results in all values less than the pivot to the left and values greater to the right. However, this doesn't mean that they are yet correctly sorted. Subsequently, we use recursion to call the function again to sort the left and right parts of the array, with each function call having a new pivot.

Let's see step by step how this sorting would look: [5, 3, 2, 1, 4]

![quicksort](https://user-images.githubusercontent.com/48635609/91211629-7a2ead00-e6e5-11ea-84f4-d632ddc35841.png)

Now, let's look at a code example:

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := []int{5, 3, 2, 1, 4}

	fmt.Println("\n --- Unsorted ---\n\n", array, "\n")
	quicksort(array)
	fmt.Println("\n --- Sorted ---\n\n", array, "\n")
}

func quicksort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	var start, end int = 0, len(a) - 1
	pivot := rand.Int() % len(a)

	a[pivot], a[end] = a[end], a[pivot]

	for i, _ := range a {
		if a[i] < a[end] {
			a[start], a[i] = a[i], a[start]
			start++
		}
	}

	a[start], a[end] = a[end], a[start]

	quicksort(a[:start])
	quicksort(a[start+1:])

	return a
}
```

At the beginning of the `quicksort` function, the algorithm checks if the array's size is less than or equal to 1. If so, it returns the array as it is. The function then declares two variables, `start` and `end`, representing the indices of the first and last elements of the array, respectively. The pivot is chosen as a random value from within the array and is placed at the end.

Following this process, the array is traversed, comparing each element at position `i` with the pivot. If the element at position `i` is smaller, a swap is performed with the element at the `start` position. The variable `start` is used to keep track of which elements have been swapped, incrementing with each swap.

At the end of the loop, the pivot is swapped with the element at the `start` position, as explained in the algorithm's description.

Finally, the function `quicksort` is called recursively, first with the left part of the array `a[:start]` (which returns the array from position 0 up to the `start` position) and then with the right part `a[start+1:]` (which returns the array from position `start+1` to the end). This division of the array into multiple parts and the sorting of each part separately continues until the entire array is sorted.

### References

- [Golang Programs - Quick Sort Implementation](https://www.golangprograms.com/golang-program-for-implementation-of-quick-sort.html)

- [Pantuza Blog - The QuickSort Sorting Algorithm](https://blog.pantuza.com/artigos/o-algoritmo-de-ordenacao-quicksort)

### Supplementary Materials

- [QUICKSORT | Algoritmos #8](https://www.youtube.com/watch?v=wx5juM9bbFo) - video