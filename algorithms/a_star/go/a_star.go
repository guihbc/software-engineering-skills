package main

import (
	"container/list"
	"fmt"
)

type Vertex struct { // struct representando um vértice
	data        interface{} // dado que iremos adidionar no vértice
	inputEdges  *list.List  // arestas de entrada
	outputEdges *list.List  // arestas de saída
	h           float64
	previousG   float64
}

func (v *Vertex) create(value interface{}, h float64) { // função que cria um novo vértice
	v.data = value             // setando o dado do vértice
	v.inputEdges = list.New()  // inicia como uma lista vazia
	v.outputEdges = list.New() // inicia como uma lista vazia
	v.h = h
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

func NewGraph() *Graph {
	return &Graph{
		vertex: list.New(),
		edges:  list.New(),
	}
}

func (g *Graph) addVertex(data interface{}, h float64) *Vertex { // função que adiciona um vértice ao grafo
	var vertex *Vertex = &Vertex{} // define um grafo
	vertex.create(data, h)         // cria um grafo passando o dado
	g.vertex.PushBack(vertex)      // adiciona o vértice na lista de vértices do grafo

	return vertex
}

func (g *Graph) addEdge(weight float64, start, end *Vertex) { // função que adiciona uma aresta entre os grafos
	// var start *Vertex = g.getVertex(dataStart) // Pesquisa o vértice através do dado - inicio
	// var end *Vertex = g.getVertex(dataEnd)     // Pesquisa o vértice através do dado - final

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

func isVisited(l *list.List, e *Vertex) bool {
	for i := l.Front(); i != nil; i = i.Next() {
		var vertex *Vertex = i.Value.(*Vertex)
		if vertex == e {
			return true
		}
	}

	return false
}

// --------------------------------- algoritmo A* ----------------------------------------------

type F struct { // Struct que representa o F(n)
	value     float64 // Valor do cálculo F(n) = G(n) + H(n)
	vertex    *Vertex // Endereço do vértice referente a este valor de F
	previousG float64 // Valor da soma de G(n) dos vértices antecessores
}

func (g *Graph) aStar(initialVertex, finalVertex *Vertex) []*Vertex { // Função do algoritmo A*, recebe o vértice inicial e o vértice final
	closed := list.New() // Variável referente a lista dos vértices fechados
	opened := list.New() // Variável referente a lista dos vértices abertos
	var result = []*Vertex{}

	closed.PushBack(initialVertex) // Adiciona o vértice inicial na lista dos vértices fechados

	var father *Vertex = initialVertex // Vértice pai = vértice inicial
	result = append(result, father)

	fmt.Println(father.data.(string))

	for father != finalVertex { // Enquanto o vértice pai for diferente do vértice final, continua o loop
		for i := father.outputEdges.Front(); i != nil; i = i.Next() { // Percorre cada aresta do vértice pai
			var edge *Edge = i.Value.(*Edge)                    // Variável referente a aresta
			var children *Vertex = edge.end                     // Recupera o filho deste vértice pai através da aresta
			children.previousG = father.previousG + edge.weight // previousG do vértice filho = previousG do vértice pai + seu custo

			opened.PushBack(&F{ // Adiciona um Struct de F na lista dos vértices abertos contendo as informações do vértice filho, incluindo seu valor de f
				value:     children.previousG + children.h,
				vertex:    children,
				previousG: children.previousG,
			})
		}

		element := getMinF(opened)                      // Retorna o elemento com o menor valor de F
		father = element.Value.(*F).vertex              // O vértice pai agora passa a ser o elemento retornado
		father.previousG = element.Value.(*F).previousG // Por garamtia o previousG do novo pai é o previousG presente no elemento retornado

		result = append(result, father)

		closed.PushBack(father) // Adiciona o novo pai a lista dos vértices fechados
		opened.Remove(element)  // Remove o elemento da lista dos vértices aberto
	}

	return result
}

func getMinF(l *list.List) *list.Element { // Função que retorna o menor valor de F(n)
	var f *F = nil               // Variável que armazena um endereço de F
	var min float64 = 999        // Variável min que armazena o valor mínimo de F
	var minF *list.Element = nil // Variável minF que armazena um endereço de *list.Element referente ao elemento com o valor mínimo, que será o elemento retornado

	for i := l.Front(); i != nil; i = i.Next() { // Percorre cada elemento da lista
		f = i.Value.(*F) // Variável f = endereço de F presente na lista

		if f.value <= min { // Caso o value de f seja menor ou igual ao min
			min = f.value // Variável min passar a ser o value de f
			minF = i      // Variável minf passa a ser o elemento
		}
	}

	return minF // Retorna o elemento da lista do menor valor de f
}
