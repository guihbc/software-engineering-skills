package main

import (
	"reflect"
	"testing"
)

func TestAStar(t *testing.T) {
	var graph *Graph = NewGraph()

	a := graph.addVertex("A", 30)
	b := graph.addVertex("B", 26)
	c := graph.addVertex("C", 21)
	d := graph.addVertex("D", 7)
	e := graph.addVertex("E", 22)
	f := graph.addVertex("F", 36)
	g := graph.addVertex("G", 0)

	graph.addEdge(12, a, b)
	graph.addEdge(14, a, c)
	graph.addEdge(9, b, c)
	graph.addEdge(38, b, d)
	graph.addEdge(24, c, d)
	graph.addEdge(7, c, e)
	graph.addEdge(9, d, g)
	graph.addEdge(13, e, d)
	graph.addEdge(29, e, g)
	graph.addEdge(9, e, f)

	path := graph.aStar(a, g)
	expectedPath := []*Vertex{a, c, b, c, e, d, g}

	if !reflect.DeepEqual(path, expectedPath) {
		t.Fail()
	}
}
