# Linked Lists

Linked list ou Lista encadeada é uma estrutura de dados na qual representa uma sequência de dados de mesmo tipo, cada dado é armazenado em um nó que contém o elemento a ser armazenado e o endereço do próximo nó. No exemplo a seguir os elementos armazenados em cada nó serão do tipo `interface{}`

Vale lembrar que uma lista encadeada não tem um limite de tamanho específico, estes tipos de estruturas de dados conhecidas também como listas dinâmicas `aumentam ou diminuem de tamanho de acordo com sua necessidade`.

![linked_list](https://user-images.githubusercontent.com/48635609/102827421-4c225680-43c1-11eb-8f9d-67812d386563.gif)

Analisando a imagem a cima podemos entender melhor o funcionamento da lista encadeada, podemos observar cada nó "ligado" no nó seguinte por isso o nome lista encadeada, cada nó contém um objeto e o endereço do próximo nó tendo assim uma lista com diversos nós.

Exemplo `código em Go`:

```Go
package main

import (
	"fmt"
)

type Node struct {
	element interface{}
	next    *Node 
}

func (n *Node) isEmpty() bool {
	return n == nil
}

type LinkedList struct {
	first *Node
}

func (l *LinkedList) create() {
	l.first = nil
}

func (l *LinkedList) insert(e interface{}) {
	var node *Node = &Node{}
	node.element = e
	node.next = l.first
	l.first = node
}

func (l *LinkedList) isEmpty() bool {
	return l.first == nil
}

func (l *LinkedList) search(e interface{}) *Node {
	for n := l.first; n != nil; n = n.next {
		if n.element == e {
			return n
		}
	}

	return nil
}

func (l *LinkedList) remove(e interface{}) int {
	var previous *Node = nil
	var p *Node = l.first

	for p != nil && p.element != e {
		previous = p
		p = p.next
	}

	if p == nil {
		return -1
	}

	if previous == nil {
		l.first = p.next
	} else {
		previous.next = p.next
	}

	return 0
}

func (l *LinkedList) free() {
	for l.first != nil {
		var temp *Node = l.first.next
		l.first = nil
		l.first = temp
	}
}

func (l *LinkedList) show() {
	fmt.Print("[ ")
	print(l.first)
	fmt.Println("]")
}

func print(n *Node) {
	if !n.isEmpty() {
		print(n.next)
		fmt.Print(n.element)
		fmt.Print(" ")
	}
}

func main() {
	var list *LinkedList = &LinkedList{}
	list.create()
	list.insert(1)
	list.insert(2)
	list.insert(3)
	list.insert(4)
	list.insert(5)
	list.insert(6)
	list.insert(7)
	list.show()
	list.remove(2)
	list.show()
}
```

No código a cima temos a implementação de uma lista encadeada em código, onde as principais partes são:

Primeiramente temos um `struct` denominado `Node` que representa os nós da lista contendo os atributos:

- `element` - Dado que adicionamos na lista

- `next` - Endereço do próximo da lista

Temos também um `struct` denominado `LinkedList` que representa a lista encadeada, nele temos o atributo `first` que representa o endereço do primeiro nó da lista

Em seguida temos a função `create` que cria uma nova lista com os valores iniciais.

A função `insert` é a função que insere um novo elemento na lista, nela primeiramente:

- Criamos um novo nó;

- O atributo `element` do novo nó é igual ao elemento passado como parâmentro;

- O próximo elemento(`next`) do novo nó é igual ao primeiro nó da lista;

- O primeiro nó da lista é agora passa a ser o novo nó criado;

A função `search` faz uma pesquisa e verifica se um elemento existe na lista, nela primeiramente:

- Executa um `for` onde `n` é igual ao primeiro nó da lista, enquanto `n` for diferente de nil(Enquanto ouver nós na lista) ao final do loop `n` é igual ao próximo nó - `for n := l.first; n != nil; n = n.next`;

- Dentro do `for` verificamos se o elemento do nó é igual ao elemento passado como paramêtro, caso seja o retornamos, caso não encontre retorna nil;

A função `remove` remove um elemento da lista, nela primeiramente:

- Cria variáveis referente ao nó anterior(Anterior ao que será removido) e ao primeiro nó da lista denominadas `previous` e `p` respectivamente;

- Executa um `for` onde verifica se o elemento presente em `p` é diferente do elemento que estamos procurando;

- Dentro do for define o `previous` como sendo o `p` agora e o `p` sendo `p.next`(Próximo nó);

- Fora do for verificamos se o `p` é igual a `nil`, caso seja o `for` percorreu por todos os elementos e não encontrou o que procurava;

- Em seguida verificamos se o `previous` é igual a `nil`, caso seja significa que estamos removendo o primeiro elemento do nó e o endereço do primeiro nó da lista passa a ser o seguinte;

- Caso o `previous` seja diferente de `nil` o valor de `previous.next`(Endereço do próximo nó do nó anterior ao que queremos remover) é igual ao endereço do nó seguinte, ou seja para remover um nó da lista basta apontar no endereço do nó anterior a ele o nó seguinte ao que iremos remover;

`OBS: O arquivo linkedList.go contém comentários explicativos.`

### Referências

- http://fbarth.net.br/materiais/docs/estruturas/aListaEncadeada.pdf.
