# Queues

A queue is a data structure that allows the addition and removal of elements with a specific rule: the first element to be added is the first one to be removed. This rule follows the **FIFO (First-In, First-Out)** order, which means that the first item to enter the queue is also the first to leave. Queues are similar to a real-life queue of people, where the person who arrives first is the one to be served first.

Here's a simple explanation with a visual representation:

![queue](https://user-images.githubusercontent.com/48635609/102701712-87454e00-4238-11eb-8a5a-973837dad29e.gif)

In the example above, you can see how the queue works. New elements are added at the back of the queue, and when an element is removed, it's always the first one (at the front of the queue) that gets removed.

Now, let's look at a Go code example that represents a circular queue:

```Go
package main

import (
	"fmt"
	"unsafe"
)

type Queue struct {
	capacity int
	data     []float64
	first    int
	last     int
	nItems   int
}

func (q *Queue) create(c int) {
	q.capacity = c
	q.data = make([]float64, c*int(unsafe.Sizeof(c)))
	q.first = 0
	q.last = -1
	q.nItems = 0
}

func (q *Queue) insert(v float64) {
	if q.last == q.capacity-1 {
		q.last = -1
	}

	q.last++
	q.data[q.last] = v
	q.nItems++
}

func (q *Queue) remove() {
	q.first++

	if q.first == q.capacity {
		q.first = 0
	}

	q.nItems--
}

func (q *Queue) show() {
	fmt.Println(q.data[q.first : q.last+1])
}

func main() {
	var q *Queue = &Queue{}
	q.create(10)
	q.insert(4)
	q.insert(3)
	q.show()
	q.remove()
	q.show()
	q.insert(10)
	q.show()
	q.remove()
	q.show()
}
```

In this code, we have implemented a circular queue in Go:

1. We define a `Queue` struct with attributes such as `capacity` (maximum number of items), `data` (the data store), `first` (index of the first element in the queue), `last` (index of the last element in the queue), and `nItems` (number of items currently in the queue).

2. The `create` method initializes the attributes of the queue with initial values.

3. The `insert` method adds a new element to the queue. If the `last` index reaches the end of the capacity, it wraps around to the beginning of the queue. This behavior creates a circular queue.

4. The `remove` method removes the oldest element from the queue. It advances the `first` index and wraps around to the beginning of the queue if necessary.

5. The `show` method displays the current contents of the queue.

In the `main` function, we demonstrate the use of this queue with insertions, removals, and display.

Comments in the code provide further explanations for each part.

This code provides a simple example of a queue data structure implemented in Go.