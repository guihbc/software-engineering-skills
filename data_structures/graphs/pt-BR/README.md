# Graphs

Um Graph ou Grafo é uma estrutura de dados que representa um par de conjuntos denominados:

- `Vértices` (Vertex)
- `Arestas` (Edges)

Cada `Aresta` está associoada a dois `Vértices` sendo o primeiro o ponto inicial e o segundo o ponto final da aresta.

Se tivermos uma aresta `a` e esta tiver um ponto inicial `v` e um ponto final `w` dizemos que `a` vai de `v` até `w`.

Um exemplo na prática é que podemos representar a estrutura da rede WWW usando grafo onde os `vértices` são as páginas web e as `arestas` são os links que direcionam para as outras páginas. 

Abaixo temos imagens de grafos para um melhor entendimento.

![graph](https://user-images.githubusercontent.com/48635609/102935498-08912080-4485-11eb-9629-8f4bdeb6c93a.png)

![directed](https://user-images.githubusercontent.com/48635609/102935479-fd3df500-4484-11eb-9f17-7cca43e72749.png)

Podemos notar dois grafos com `arestas` diferentes, no primeiro grafo temos uma relação simétrica ou seja se o vértice `v` está associoado a `w` então `w` está associado a `v`, já no segundo temos um `grafo direcionado` ou também chamado de `digrafo` onde `w` pode está associado a `v` mas `v` não precisa está associado a `w`. 

Abaixo veremos como implementar um grafo em `Go`:

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
	graph.addVertex("João")
	graph.addVertex("Clebin")
	graph.addVertex("Rogerin")

	graph.addEdge(2, "Guilherme", "Yasmim")
	graph.addEdge(3, "Yasmim", "Clebin")
	graph.addEdge(1, "Clebin", "João")
	graph.addEdge(1, "Guilherme", "João")
	graph.addEdge(2, "Rogerin", "Yasmim")
	graph.addEdge(3, "Rogerin", "Guilherme")

	graph.BFS()
}
```

No código a cima primeiramente temos um struct denominado `Vertex` que representa um vértice, contendo os atributos:

- `data` Representa o dado a ser armazenado no vértice
- `inputEdges` É uma lista contendo todas as arestas de entrada do vértice
- `outputEdges` É uma lista contendo todas as arestas de saída do vértice

Temos uma função `create` que cria um vértice;

A função `addInputEdge` adiciona uma aresta de entrada ao vértice;

A função `addOutputEdge` adiciona uma aresta de saída ao vértice;

Em seguida temo um struct denominado `Edge` que representa uma aresta, contendo os atributos:

- `weight` Peso da aresta
- `start` Vértice de início
- `end` Vértice de saída

Temos uma função `create` que cria uma aresta;

Em seguida temos um struct denominado `Graph` que representa um grafo, contendo os atributos:

- `vertex` Lista contendo todos os vértices do grafo
- `edges` Lista contendo todas as arestas do grafo

Temos uma função `create` que cria um grafo;

A função `addVertex` adiciona um novo vértice ao grafo, nela criamos um novo vértice e em seguida o adicionamos na lista de vértices do grafo;

A função `addEdge` adiciona uma aresta ao grafo, ela é responsável por associar os vértices. Nela primeiramente fazemos uma busca pelos vértices de início e final, em seguida uma nova aresta é criada passando os vértices de início e final encontrados, no `vértice de início` usamos a função `addOutputEdge` para adicionar a aresta criada como sendo uma aresta de saída dele e usamos a função `addInputEdge` no `vértice de final` para adicionar a aresta criada como uma aresta de entrada nele, ou seja `a aresta sai de início e entra em final` início -> final;

A função `getVertex` procura um vértice através do seu valor, percorrendo cada vértice presente no grafo, quando encontrado o retorna, caso não encontre o vértice retorna nil;

A função `BFS` representa o algoritmo `Breadth-first search`;

A função `isVisited` faz parte do algoritmo `Breadth-first search`;

`OBS: O arquivo graph.go contém comentários explicativos`

### Referências

- https://www.youtube.com/watch?v=jq0N1LDOTlw&t=558s.

- https://www.youtube.com/watch?v=fWuqx5JWSQA

- https://www.ime.usp.br/~pf/algoritmos_para_grafos/aulas/graphs.html

- https://www.ime.usp.br/~pf/algoritmos_em_grafos/aulas/grafos.html

- https://www.inf.ufsc.br/grafos/definicoes/definicao.html
