# Graphs

A Graph, or Graph data structure, is a data structure that represents a pair of sets known as:

- `Vertices`
- `Edges`

Each `Edge` is associated with two `Vertices`, with the first being the starting point and the second being the endpoint of the edge.

If we have an edge `a` with a starting point `v` and an endpoint `w`, we say that `a` goes from `v` to `w`.

A practical example is that we can represent the structure of the World Wide Web (WWW) using a graph, where the `Vertices` are web pages and the `Edges` are links that direct to other pages.

Below, we have images of graphs for a better understanding:

![graph](https://user-images.githubusercontent.com/48635609/102935498-08912080-4485-11eb-9629-8f4bdeb6c93a.png)

![directed](https://user-images.githubusercontent.com/48635609/102935479-fd3df500-4484-11eb-9f17-7cca43e72749.png)

You can notice two graphs with different types of `Edges`. In the first graph, we have a symmetric relationship, meaning if vertex `v` is associated with `w`, then `w` is also associated with `v`. In the second graph, we have a `directed graph`, also called a `digraph`, where `w` can be associated with `v`, but `v` does not need to be associated with `w`.

Below, we'll see how to implement a graph in `Go`:

```Go
package main

import (
	"container/list"
	"fmt"
)

type Vertex struct {
	data        interface{}
	inputEdges  *list.List
	outputEdges *list.List
}

func (v *Vertex) create(value interface{}) {
	v.data = value
	v.inputEdges = list.New()
	v.outputEdges = list.New()
}

func (v *Vertex) addInputEdge(e *Edge) {
	v.inputEdges.PushBack(e)
}

func (v *Vertex) addOutputEdge(e *Edge) {
	v.outputEdges.PushBack(e)
}

type Edge struct {
	weight float64
	start  *Vertex
	end    *Vertex
}

func (e *Edge) create(w float64, start, end *Vertex) {
	e.weight = w
	e.start = start
	e.end = end
}

type Graph struct {
	vertex *list.List
	edges  *list.List
}

func (g *Graph) create() {
	g.vertex = list.New()
	g.edges = list.New()
}

func (g *Graph) addVertex(data interface{}) {
	var vertex *Vertex = &Vertex{}
	vertex.create(data)

	g.vertex.PushBack(vertex)
}

func (g *Graph) addEdge(weight float64, dataStart, dataEnd interface{}) {
	var start *Vertex = g.getVertex(dataStart)
	var end *Vertex = g.getVertex(dataEnd)

	var edge *Edge = &Edge{}
	edge.create(weight, start, end)

	start.addOutputEdge(edge)
	end.addInputEdge(edge)

	g.edges.PushBack(edge)
}

func (g *Graph) getVertex(data interface{}) *Vertex {
	var vertex *Vertex = nil

	for l := g.vertex.Front(); l != nil; l = l.Next() {
		v := l.Value.(*Vertex)

		if v.data == data {
			vertex = l.Value.(*Vertex)
			break
		}
	}

	return vertex
}

func (g *Graph) BFS() {
	visiteds := list.New()
	queue := list.New()

	element := g.vertex.Front()
	current := element.Value.(*Vertex)

	visiteds.PushBack(current)
	fmt.Println(current.data)

	queue.PushBack(current)

	for queue.Len() > 0 {
		var visited *Vertex = queue.Front().Value.(*Vertex)

		for e := visited.outputEdges.Front(); e != nil; e = e.Next() {
			var next *Vertex = e.Value.(*Edge).end

			if vis := isVisited(visiteds, next); !vis {
				visiteds.PushBack(next)
				fmt.Println(next.data)
				queue.PushBack(next)
			}

		}

		queue.Remove(queue.Front())
	}
}

func isVisited(l *list.List, e *Vertex) bool {
	for i := l.Front(); i != nil; i = i.Next() {
		var vertex *Vertex = i.Value.(*Vertex)
		if vertex == e {
			return true
		}
	}

	return false
}

func main() {
	var graph *Graph = &Graph{}
	graph.create()

	graph.addVertex("Guilherme")
	graph.addVertex("Yasmim")
	graph addVertex("João")
	graph addVertex("Clebin")
	graph addVertex("Rogerin")

	graph.addEdge(2, "Guilherme", "Yasmim")
	graph.addEdge(3, "Yasmim", "Clebin")
	graph.addEdge(1, "Clebin", "João")
	graph.addEdge(1, "Guilherme", "João")
	graph.addEdge(2, "Rogerin", "Yasmim")
	graph.addEdge(3, "Rogerin", "Guilherme")

	graph.BFS()
}
```

In the code above, we first have a `Vertex` struct that represents a vertex, containing the attributes:

- `data` Represents the data to be stored in the vertex
- `inputEdges` Is a list containing all input edges to the vertex
- `outputEdges` Is a list containing all output edges from the vertex

We have a `create` function that creates a vertex.

The `addInputEdge` function adds an input edge to the vertex.

The `addOutputEdge` function adds an output edge to the vertex.

Next, we have an `Edge` struct representing an edge, with the attributes:

- `weight` Weight of the edge
- `start` Start vertex
- `end` End vertex

We have a `create` function that creates an edge.

Next, we have a `Graph` struct that represents a graph, with the attributes:

- `vertex` A list containing all the vertices in the graph
- `edges` A list containing all the edges in the graph

We have a `create` function that creates a graph.

The `addVertex` function adds a new vertex to the graph. In it, we create a new vertex and then add it to the list of vertices in the graph.

The `addEdge` function adds an edge to the graph, which is responsible for associating the vertices. First, we search for the start and end vertices. Then, a new edge is created by passing the found start and end vertices. On the `start vertex`, we use the `addOutputEdge` function to add the created edge as an output edge, and on the `end vertex`, we use the `addInputEdge` function to add the created edge as an input edge. In other words, the edge goes from `start` to `end`.

The `getVertex` function searches for a vertex by its value, iterating through each vertex present in the graph, and returns it when found. If the vertex is not found, it returns nil.

The `BFS` function represents the Breadth-first search algorithm.

The `isVisited` function is part of the Bread