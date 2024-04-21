package main

import "testing"

func TestQueueInsert(t *testing.T) {
	var q *Queue = &Queue{}
	q.create(1)
	q.insert(1)

	if q.nItems != 1 {
		t.Errorf("expected %d but got %d", 1, q.nItems)
	}
}

func TestQueueRemove(t *testing.T) {
	var q *Queue = &Queue{}
	q.create(3)

	q.insert(1)
	q.insert(5)
	q.insert(10)

	q.remove()
	first := q.data[q.first]

	if q.nItems != 2 {
		t.Errorf("expected %d but got %d", 2, q.nItems)
	}

	if first != 5 {
		t.Errorf("expected %d but got %f", 5, first)
	}
}

func TestQueueCapacity(t *testing.T) {
	var q *Queue = &Queue{}
	q.create(3)

	q.insert(1)
	q.insert(5)
	q.insert(10)
	err := q.insert(4)

	if err == nil {
		t.Errorf("expected error but nothing returned")
	}

	if err.Error() != "queue overflow" {
		t.Errorf("expected %s but got %s", "queue overflow", err.Error())
	}
}
