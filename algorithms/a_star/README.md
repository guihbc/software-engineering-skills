# A*

The A* algorithm is used for pathfinding in a [graph](https://github.com/GuilhermehChaves/google-skills/tree/develop/estrutura_de_dados/graphs). It seeks the best path between an initial vertex and a final vertex.

This algorithm is widely used in various applications, including games and route finding in GPS systems.

## How Does it Work?

The basic idea behind the algorithm is that starting from the current vertex (the initial vertex), it examines all its neighbors and applies the evaluation function `F(n) = G(n) + H(n)`, where:

- G(n) = Cost of reaching the vertex.
- H(n) = Heuristic value of the vertex, indicating the distance from the vertex to the final vertex.

The vertex with the lowest F(n) value becomes the current vertex, and this process continues until the current vertex is the final vertex.

## Example

![astar](https://user-images.githubusercontent.com/48635609/104031436-7c346c80-51ab-11eb-8c5f-6e432b5b1146.PNG)

Considering the directed graph above, we want to go from point A to G. We already have heuristic values defined as follows:

| Vertex | H |
| ------ | --- |
| A | 30 |
| B | 26 |
| C | 21 |
| D | 7 |
| E | 22 |
| F | 36 |
| G | 0 |

Keep in mind that the heuristic function can be calculated within the algorithm, but in this example, we already have it defined.

Let's take a look at the following Go code:

```Go
type F struct {
	value     float64
	vertex    *Vertex
	previousG float64
}

func (g *Graph) aStar(initialVertex, finalVertex *Vertex) {
	closed := list.New()
	opened := list.New()

	closed.PushBack(initialVertex)

	var father *Vertex = initialVertex
	fmt.Print(father.data.(string))

	for father != finalVertex {
		for i := father.outputEdges.Front(); i != nil; i = i.Next() {
			var edge *Edge = i.Value.(*Edge)
			var children *Vertex = edge.end
			children.previousG = father.previousG + edge.weight

			opened.PushBack(&F{
				value:     children.previousG + children.h,
				vertex:    children,
				previousG: children.previousG,
			})
		}

		element := getMinF(father, opened)
		father = element.Value.(*F).vertex
		father.previousG = element.Value.(*F).previousG

		fmt.Print(" -> " + father.data.(string))

		closed.PushBack(father)
		opened.Remove(element)
	}
}

func getMinF(old *Vertex, l *list.List) *list.Element {
	var f *F = nil
	var min float64 = 999
	var minF *list.Element = nil

	for i := l.Front(); i != nil; i = i.Next() {
		f = i.Value.(*F)

		if f.value <= min {
			min = f.value
			minF = i
		}
	}

	return minF
}

func main() {
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

	graph.aStar(a, g)

}
```

In the code above, we have an implementation of the A* algorithm in Go. Some key points to note are:

### First, there is a struct called F.

```Go
type F struct {
	value     float64
	vertex    *Vertex
	previousG float64
}
```

This struct is used to store information about F(n) for the vertices. It's separated from the Vertex struct to avoid conflicts. It includes:

- value = The value of F(n) calculated as `F(n) = G(n) + H(n)`.
- vertex = The address of the vertex associated with this F(n) value.
- previousG = The sum of G(n) values of the preceding vertices.

### Next, there's a function called getMinF.

```Go
func getMinF(l *list.List) *list.Element {
	func getMinF(l *list.List) *list.Element {
		var f *F = nil
		var min float64 = 999
		var minF *list.Element = nil

		for i := l.Front(); i != nil; i = i.Next() {
			f = i.Value.(*F)

			if f.value <= min {
				min = f.value
				minF = i
			}
		}

		return minF
    }
}
```

This function returns the element with the minimum F(n) value from a list of elements. It loops through the list and finds the element with the smallest F value.

### Finally, we have the main A* algorithm function called aStar.

```Go
func (g *Graph) aStar(initialVertex, finalVertex *Vertex) {
	closed := list.New()
	opened := list.New()

	closed.PushBack(initialVertex)

	var father *Vertex = initialVertex
	fmt.Print(father.data.(string))

	for father != finalVertex {
		for i := father.outputEdges.Front(); i != nil; i = i.Next() {
			var edge *Edge = i.Value.(*Edge)
			var children *Vertex = edge.end
			children.previousG = father.previousG + edge.weight

			opened.PushBack(&F{
				value:     children.previousG + children.h,
				vertex:    children,
				previousG: children.previousG,
			})
		}

		element := getMinF(opened)
		father = element.Value.(*F).vertex
		father.previousG = element.Value.(*F).previousG

		fmt.Print(" -> " + father.data.(string))

		closed.PushBack(father)
		opened.Remove(element)
	}
}
```

This function is the core of the A* algorithm. It operates as follows:

- It initializes two lists: `closed`, which stores closed vertices, and `opened`, which stores open vertices.
- It starts with the initial vertex as the current vertex and prints its value.
- It enters a loop that continues until the current vertex equals the final vertex.
- Within the loop, it iterates through the outgoing edges of the current vertex and calculates the values for the child vertices.
- These child vertices are added to the list of open vertices.
- The algorithm then selects the vertex with the minimum F(n) value from the open vertices.
- The selected vertex becomes the new current vertex, and the loop continues.
- The chosen vertex is printed, added to the closed vertices, and removed from the open vertices.

The algorithm repeats these steps until the current vertex is the final vertex.

In conclusion, the A* algorithm starts at the initial vertex, evaluates its child vertices, and selects the one with the lowest F(n) value. It continues this process until it reaches the final vertex.