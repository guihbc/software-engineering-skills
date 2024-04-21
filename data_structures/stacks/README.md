# Stacks

A stack is a data structure that allows the addition and removal of elements with a specific rule: the last element to be added is the first one to be removed. This rule follows the **LIFO (Last-In, First-Out)** order, which means that the last item to enter the stack is the first to leave. Stacks are similar to a real-life stack of plates, where the plate placed on top of the stack is the first to be removed.

Here's a simple explanation with a visual representation:

![stack](https://user-images.githubusercontent.com/48635609/102821122-4757a580-43b5-11eb-8e79-63c915d24378.gif)

In the example above, you can see how the stack works. New elements are added on top of the stack, and when an element is removed, it's always the last one (at the top of the stack) that gets removed.

Now, let's look at a Go code example that represents a stack:

```Go
package main

import "fmt"

type Stack struct {
	top      int
	capacity int
	data     []float64
}

func (s *Stack) create(c int) {
	s.top = -1
	s.capacity = c
	s.data = make([]float64, c)
}

func (s *Stack) isEmpty() bool {
	return s.top == -1
}

func (s *Stack) isFull() bool {
	return s.top == s.capacity
}

func (s *Stack) push(v float64) error {
	if s.isFull() {
		fmt.Println("Stack is full")
		return fmt.Errorf("Stack is full")
	}

	s.top++
	s.data[s.top] = v

	return nil
}

func (s *Stack) pop() (float64, error) {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return 0, fmt.Errorf("Stack is empty")
	}

	element := s.data[s.top]
	s.top--

	return element, nil
}

func (s *Stack) getTop() float64 {
	return s.data[s.top]
}

func (s *Stack) show() {
	fmt.Println(s.data[0 : s.top+1])
}

func main() {
	var s *Stack = &Stack{}
	s.create(10)
	s.push(2)
	s.push(7)
	s.show()
	s.pop()
	s.show()
}
```

In this code, we have implemented a stack in Go:

1. We define a `Stack` struct with attributes such as `top` (the current top index), `capacity` (maximum stack size), and `data` (the array to store the stack elements).

2. The `create` method initializes the attributes of the stack with initial values.

3. The `isEmpty` method checks whether the stack is empty based on the value of `top`.

4. The `isFull` method checks whether the stack is full based on the value of `top` compared to the capacity.

5. The `push` method is responsible for pushing a new element onto the stack. It first checks if the stack is full. If it is, an error is returned. Otherwise, it increments the `top`, adds the element to the top position in the data array, and returns no error.

6. The `pop` method is responsible for popping (removing) the top element from the stack. It first checks if the stack is empty. If it is, an error is returned. Otherwise, it retrieves the element at the current top position, decrements the `top`, and returns the element along with no error.

7. The `getTop` method simply returns the element at the top of the stack without removing it.

8. The `show` method displays the current contents of the stack from the bottom to the top.

In the `main` function, we demonstrate the use of this stack with pushes, pops, and display of the stack's contents.

Comments in the code provide further explanations for each part.

This code provides a simple example of a stack data structure implemented in Go.