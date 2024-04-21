# A*

O algoritmo A* é um algoritmo usado para busca de caminho em um [grafo](https://github.com/GuilhermehChaves/google-skills/tree/develop/estrutura_de_dados/graphs). Ele busca o melhor caminho entre um vértice inicial e um vértice final.

Este algoritmo é muito utilizado em jogos, encontrar rotas como em GPS, entre outros...

## Como funciona?

O que o algoritmo faz basicamente é, a partir do vértice atual (Vértice inicial) ele pega todos os seus vizinhos e aplica a função de avaliação `F(n) = G(n) + H(n)` onde: 

- G(n) = Custo de chegada até o vértice
- H(n) = Valor heurístico do vértice - Indica a distância do vértice até o vértice final

O vértice que tiver como retorno o menor valor de F(n) passará a ser o vértice atual, este procedimento se repete até que o vértice atual seja o vértice final.

## Exemplo

![astar](https://user-images.githubusercontent.com/48635609/104031436-7c346c80-51ab-11eb-8c5f-6e432b5b1146.PNG)

Levando em consideração o digrafo acima, Queremos partir do ponto A ao G, onde já temos os valores de heurítica definidos como:

| Vétice | H |
| ------ | --- |
| A | 30 |
| B | 26 |
| C | 21 |
| D | 7 |
| E | 22 |
| F | 36 |
| G | 0 |

Lembrando que a função de heurística pode ser calculada no próprio algoritmo, mas neste exemplo já teremos ela definida.

Vejamos o código abaixo escrito e `Go`:

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

No código acima temos uma representação do algoritmo A* implementado em `Go` nele podemos notar:

### Primeiramente temos um struct denominado F.

```Go
type F struct { // Struct que representa o F(n)
	value     float64 // Valor do cálculo F(n) = G(n) + H(n)
	vertex    *Vertex // Endereço do vértice referente a este valor de F
	previousG float64 // Valor da soma de G(n) dos vértices antecessores
}
```

Este struct foi criado para armazenar as informações de `F` dos nossos vértices, colocando-o como um struct separado de `Vertex` para não haver conflitos, nele temos:

- value = Valor de F(n) dado por `F(n) = G(n) + H(n)`
- vertex = Endereço do vértice que contém este valor de F(n)
- previousG = Valor de G(n) dos vértices antecessores a este

### Temos também a função denominada getMinF.

```Go
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
```

Essa função retornará uma variável do tipo `*list.Element` (Tipo armazenado na lista do Go) que armazena dados do tipo `F` que foi comentado acima. Analisando-a notamos que:

- Temos uma variável chamada `f` que armazenará um endereço de `F` começanco com `nil`
- Temos uma variável chamada `min` que armazenará o valor do F com o menor `value`
- Temos uma variável chamada `minf` que armazenará o elemento da lista com o menor valor de `value`

- Dentro de um loop percorremos cada elemento da lista
- Define `f` como sendo o elemento atual do loop
- Se o `value` deste elemento for menor ou igual ao `min` agora o `min` é igual ao `value` de `f` e o `minF` é igual ao elemento atual

- Fora do loop retornamos o valor de `minF`, assim temos o valor do elemento com o menor valor de F(n)

### Por último temos a função do algotimo A* denominada aStar

```Go
func (g *Graph) aStar(initialVertex, finalVertex *Vertex) { // Função do algoritmo A*, recebe o vértice inicial e o vértice final
	closed := list.New() // Variável referente a lista dos vértices fechados
	opened := list.New() // Variável referente a lista dos vértices abertos

	closed.PushBack(initialVertex) // Adiciona o vértice inicial na lista dos vértices fechados

	var father *Vertex = initialVertex // Vértice pai = vértice inicial
	fmt.Print(father.data.(string))

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

		element := getMinF(opened)              // Retorna o elemento com o menor valor de F
		father = element.Value.(*F).vertex              // O vértice pai agora passa a ser o elemento retornado
		father.previousG = element.Value.(*F).previousG // Por garamtia o previousG do novo pai é o previousG presente no elemento retornado

		fmt.Print(" -> " + father.data.(string)) // Printa o valor do vértice de menor F ou seja as decisões que o algoritmo tomou

		closed.PushBack(father) // Adiciona o novo pai a lista dos vértices fechados
		opened.Remove(element)  // Remove o elemento da lista dos vértices aberto
	}
}
```

Esta é a fução principal do algoritmo que encontra a melhor possibilidade de caminho, oque ela faz é:

- Define uma variável denominada `closed` que é uma lista, que armazenará os vértices fechados
- Define uma variável denominada `opened` que é uma lista, que armazenará os vértices abertos

`Vértices abertos` são os vértices encontrados porém que ainda não foram visitados (Não passou por esse caminho)
`Vértices fechados` são os vértices que já foram visitados (Já passou por este caminho e não pode ir nele novamente)

- Adiciona o vértice inical na lista dos vértices fechados, pois ele é o primeiro e já estamos nele
- Define uma variável chamada `father` que representa o vértice pai e o inicial como sendo o vértice inicial
- Da um print no valor do vértice pai

- Temos um loop que é percorrido enquanto o valor de `father` for diferente do vértice final
- Dentro do loop anterior temos outro loop que percorre cada aresta do vértice pai
- Recuperamos o valor da aresta
- A partir do valor da aresta temos uma variável denominada `children` que representa o vértice filho
- Cada vértice tem uma propriedade chamada `previousG` que armazena o valor de G(n) de todos os vértices antecessores a ele e o setamos como o previousG do seu pai + se próprio valor de G(n) assim sempre teremos a soma de todos os G(n) de seus antecessores + o próprio G(n) do vértice
- Adicionamos na lista dos vértices abertos o endereço de `F` onde armazena o `value`, `vertex`, `previousG` do `children`, agora este vértice esta suscetível a ser um caminho possível

- Fora do segundo loop chamamos a função `getMinF` mencionada acima que retorna o elemento com o menor valor de F
- Recuperamos o valor deste elemento que armazena nosso `F` como uma `interface{}` o convertemos para o tipo `F` de fato e recuperamos a propriedade `vertex` esse `vertex` agora é o novo pai
- Setamos o `previousG` do novo pai como o `previousG` presente em `F` por garantia para que não alteremos o valor do `previousG` presente no `vertex` acidentalmente e venha o valor errado

- Da um print no valor do menor vértice escolhido
- Adiciona o novo pai a lista dos vértices fechados
- Remove o elemento da lista dos vértices aberto (Por isso na função `getMinF` é retornado um *list.Element, para poder removelo da lista)

`Ao executar o algoritmo é possível notar que, apesar de o print ser nos melhores caminhos encontrados as vezes ele pode ir e depois acabar voltando novamente para um vértice que que já foi passado, porém por outro caminho e alguns casos podemos ignorar algumas vértices visitados para selecionar o menor caminho e no momento o algoritmo ainda não faz isso.`

Por exemplo: Temos como print os valores (`A -> C -> B -> C -> E -> D -> G`) notamos que do A ele sai para o C, em seguida ela pega o caminho para o B e do B para o C novamente, nesta grafo que é o grafo da imagem apresentada anteriormente podemos notar que, podemos ignorar o caminho B -> C então temos como menor caminho o seguinte `A -> C -> E -> D -> G`

![path](https://user-images.githubusercontent.com/48635609/104066381-144c4900-51e0-11eb-973b-e4aa7ffb17d8.PNG)

### Conclusão

Para facilitar o entendimeto o algoritmo basicamente começa do vértice inicial, checa todos os seus vértices filhos e calcula o valor de F(n) de cada filho, e os adiciona na lista de vértices abertos, a partir dos elementos presente na lista dos vértices abertos ele retorna o vértice com o menor valor de F(n) o adiona agora lista de vértices fechados e o remove da lista dos vértices abertos pois não poderá mais ser visitado por esta caminho, este vértice com o menor valor de F(n) passa agora a ser o novo pai, o mesmo procedimento se repete até que o vértice pai seja igual ao vértice final.

### Referências

- https://www.youtube.com/watch?v=FU6JQaRMMDM&t=106s

- https://pt.stackoverflow.com/questions/328048/como-%C3%A9-o-funcionamento-b%C3%A1sico-do-algoritmo-a
