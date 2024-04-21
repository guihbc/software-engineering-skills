# Queues

Queue ou Fila é uma estrutura de dados que permite a adição e a remoção de novos elementos, porém há uma regra para que isso ocorra, a regra diz que o primeiro elemento da fila será sempre o que sairá primeiro e o novo elemento a ser adicionado será sempre o último da fila se assemelhando com uma fila de pessoas por exemplo, esta regra se denomina `FIFO`(first-in first-out) - O primeiro que entra é o primeiro que sai.

![queue](https://user-images.githubusercontent.com/48635609/102701712-87454e00-4238-11eb-8a5a-973837dad29e.gif)

No Gif a cima podemos ver como a fila funciona.

Exemplo:

```Go
package main

import (
	"fmt"
	"unsafe"
)

type Queue struct {
	capacity int
	data     []float64
	first    int
	last     int
	nItems   int
}

func (q *Queue) create(c int) {
	q.capacity = c
	q.data = make([]float64, c*int(unsafe.Sizeof(c)))
	q.first = 0
	q.last = -1
	q.nItems = 0
}

func (q *Queue) insert(v float64) {
	if q.last == q.capacity-1 {
		q.last = -1
	}

	q.last++
	q.data[q.last] = v
	q.nItems++
}

func (q *Queue) remove() {
	q.first++

	if q.first == q.capacity {
		q.first = 0
	}

	q.nItems--
}

func (q *Queue) show() {
	fmt.Println(q.data[q.first : q.last+1])
}

func main() {
	var q *Queue = &Queue{}
	q.create(10)
	q.insert(4)
	q.insert(3)
	q.show()
	q.remove()
	q.show()
	q.insert(10)
	q.show()
	q.remove()
	q.show()
}
```

No exemplo a cima temos a representação de uma fila circular em código escrito em `Go`.

Primeiramente podemos notar que foi criado um `struct` chamado `Queue` contendo alguns atributos. 

Em seguida temos uma função `create` que preenche os atributos da fila com os valores iniciais.

Abaixo temos a função `insert` que insere um novo elemento na fila, primeiramente ela verifica se a fila já está toda preenchida, caso esteja o atributo referente ao índice do último elemento passa a ser `-1` logo após o `if` incrementamos 1 no atributo referente ao índice do último elemento e inserimos o valor desejado nesta posição e incrementa 1 no atributo referente ao número de elementos presentes na fila.

Temos também a função `remove` que remove o elemento mais antigo da fila, nela incrementamos 1 no atributo referente ao índice do primeiro elemento verificamos se a posição do primeiro elemento é igual ao tamanho da capacidade total caso seja o índice do primeiro elemento agora é o 0, logo após o `if` decrementamos 1 no atributo referente a quantidade de elementos na fila.

Por últimos temos a função `show` que exibe a fila na tela, nela damos um print dos dados pegando da posição do primeiro elemento até a posição do último elemento.

`OBS: No arquivo queue.go o código está com comentários explicativos.`

Esta foi uma breve explicação a respeito das filas com exemplo em código escrito em `Go`.

### Referências

- https://www.cos.ufrj.br/~rfarias/cos121/filas.html#:~:text=S%C3%A3o%20estruturas%20de%20dados%20do,e%20remove-se%20do%20in%C3%ADcio.
