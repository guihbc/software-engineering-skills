package main

import "fmt"

// Struct representing the Stack
type Stack struct {
	top      int       // top position
	capacity int       // capacity
	data     []float64 // stack data
	length   int       // current stack length
}

func (s *Stack) create(c int) { // init the stack properties given the stack size
	s.top = -1
	s.capacity = c
	s.data = make([]float64, c)
	s.length = 0
}

func (s *Stack) isEmpty() bool { // verify if the stack is empty
	return s.top == -1
}

func (s *Stack) isFull() bool { // verify if the stack is full
	return s.top == s.capacity-1
}

func (s *Stack) push(v float64) error { // stack up an element
	if s.isFull() {
		return fmt.Errorf("stack overflow")
	}

	s.top++           // increment 1 to the top position
	s.data[s.top] = v // add the element to the new top position
	s.length++        // increment 1 in the stack length

	return nil // return nil for the error
}

func (s *Stack) pop() (float64, error) { // unstack an element
	if s.isEmpty() {
		return 0, fmt.Errorf("Stack is empty") // return 0 and an error
	}

	element := s.data[s.top] // the element to be removed
	s.top--                  // decrease the top position
	s.length--               // decrease the stack length

	return element, nil // return the removed element and nil for the error
}
