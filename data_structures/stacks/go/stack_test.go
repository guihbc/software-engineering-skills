package main

import "testing"

func TestStackPush(t *testing.T) {
	var s *Stack = &Stack{}
	s.create(1)
	s.push(10)

	if !s.isFull() {
		t.Errorf("stack should be full")
	}

	if s.length != 1 {
		t.Errorf("expected length 1 but got %d", s.length)
	}
}

func TestStackPop(t *testing.T) {
	var s *Stack = &Stack{}
	s.create(3)
	s.push(10)
	s.push(20)
	s.push(30)

	s.pop()

	last := s.data[s.top]

	if last != 20 {
		t.Errorf("expected last 20 but got %f", last)
	}
}

func TestStackOverFlow(t *testing.T) {
	var s *Stack = &Stack{}
	s.create(3)
	s.push(10)
	s.push(20)
	s.push(30)
	err := s.push(40)

	if err == nil {
		t.Errorf("error should not be nil")
	}

	if err.Error() != "stack overflow" {
		t.Errorf("expected error %s but got %s", "stack overflow", err.Error())
	}
}
