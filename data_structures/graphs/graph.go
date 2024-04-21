package main

import (
	"container/list"
	"fmt"
)

type Vertex struct { // struct representando um vértice
	data        interface{} // dado que iremos adidionar no vértice
	inputEdges  *list.List  // arestas de entrada
	outputEdges *list.List  // arestas de saída
}

func (v *Vertex) create(value interface{}) { // função que cria um novo vértice
	v.data = value             // setando o dado do vértice
	v.inputEdges = list.New()  // inicia como uma lista vazia
	v.outputEdges = list.New() // inicia como uma lista vazia
}

func (v *Vertex) addInputEdge(e *Edge) { // função que adidiona uma aresta de entrada ao grafo
	v.inputEdges.PushBack(e)
}

func (v *Vertex) addOutputEdge(e *Edge) { // função que adiciona uma aresta de saída ao grafo
	v.outputEdges.PushBack(e)
}

type Edge struct { // struct que representa uma aresta
	weight float64 // peso da aresta
	start  *Vertex // vétice de inicio
	end    *Vertex // vértice de final
}

func (e *Edge) create(w float64, start, end *Vertex) { // função que cria um novo vértice
	e.weight = w
	e.start = start
	e.end = end
}

type Graph struct { // struct que representa um grafo
	vertex *list.List // vértices do grafo
	edges  *list.List // arestas do grafo
}

func (g *Graph) create() { // função que cria um grafo
	g.vertex = list.New() // começa como uma lista vazia
	g.edges = list.New()  // começa como uma lista vazia
}

func (g *Graph) addVertex(data interface{}) { // função que adiciona um vértice ao grafo
	var vertex *Vertex = &Vertex{} // define um grafo
	vertex.create(data)            // cria um grafo passando o dado
	g.vertex.PushBack(vertex)      // adiciona o vértice na lista de vértices do grafo
}

func (g *Graph) addEdge(weight float64, dataStart, dataEnd interface{}) { // função que adiciona uma aresta entre os grafos
	var start *Vertex = g.getVertex(dataStart) // Pesquisa o vértice através do dado - inicio
	var end *Vertex = g.getVertex(dataEnd)     // Pesquisa o vértice através do dado - final

	var edge *Edge = &Edge{}        // define uma nova aresta
	edge.create(weight, start, end) // cria uma nova aresta, passando os vértices de início e final

	start.addOutputEdge(edge) // adiciona a aresta de saida no vértice de ínicio
	end.addInputEdge(edge)    // adiciona a aresta de entrada no vértice de saída
	g.edges.PushBack(edge)    // adiciona a aresta na lista de arestas do grafo
}

func (g *Graph) getVertex(data interface{}) *Vertex { // função que pesquisa um vértice
	var vertex *Vertex = nil // vértice

	for l := g.vertex.Front(); l != nil; l = l.Next() { // loop que percorre cada vértice do grafo
		v := l.Value.(*Vertex) // passando do tipo interface{} para *Vertex

		if v.data == data { // verifica se o dado do vértice atual é igual ao procurado
			vertex = v // vertex é igual ao v
			break      // para o loop
		}
	}

	return vertex // retorna o vértice caso seja encontrado ou nil
}

func (g *Graph) BFS() { // Algoritmo Breadth-first search
	visiteds := list.New()
	queue := list.New()

	element := g.vertex.Front()
	current := element.Value.(*Vertex)

	// current := g.getVertex("Rogerin")

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
	var graph *Graph = &Graph{} // define um grafo
	graph.create()              // cria o grafo com os valores iniciais

	graph.addVertex("Guilherme") // adiciona um novo vértice com o valor Guilherme
	graph.addVertex("Yasmim")    // adiciona um novo vértice com o valor Yasmim
	graph.addVertex("João")      // adiciona um novo vértice com o valor João
	graph.addVertex("Clebin")    // adiciona um novo vértice com o valor Clebin
	graph.addVertex("Rogerin")   // adiciona um novo vértice com o valor Rogerin

	graph.addEdge(2, "Guilherme", "Yasmim")  // cria uma aresta entre Guilherme e Yasmim
	graph.addEdge(3, "Yasmim", "Clebin")     // cria uma aresta entre Yasmim e Clebin
	graph.addEdge(1, "Clebin", "João")       // cria uma aresta entre Clebin e João
	graph.addEdge(1, "Guilherme", "João")    // cria uma aresta entre Guilherme e João
	graph.addEdge(2, "Rogerin", "Yasmim")    // cria uma aresta entre Rogerin e Yasmim
	graph.addEdge(3, "Rogerin", "Guilherme") // cria uma aresta entre Rogerin e Guilherme

	graph.BFS()
}
