# Stacks

Stack ou pilha, é uma estrutura de dados que permite a adição e a remoção de novos elementos, porém há uma regra para que isso ocorra, a regra diz que o último elemento a ser inserido será sempre o  primeiro a ser retirado da pilha, se assemelhando com uma pilha de pratos por exemplo, esta regra se denomina `LIFO`(last-in first-out) - O último que entra é o primeiro que sai.

![stack](https://user-images.githubusercontent.com/48635609/102821122-4757a580-43b5-11eb-8e79-63c915d24378.gif)

No gif a cima podemos entender melhor como uma pilha funciona.

Exemplo `código em Go`:

```Go
package main

import "fmt"

type Stack struct {
	top      int
	capacity int
	data     []float64
}

func (s *Stack) create(c int) {
	s.top = -1
	s.capacity = c
	s.data = make([]float64, c)
}

func (s *Stack) isEmpty() bool {
	return s.top == -1
}

func (s *Stack) isFull() bool {
	return s.top == s.capacity
}

func (s *Stack) push(v float64) error {
	if s.isFull() {
		fmt.Println("Stack is full")
		return fmt.Errorf("Stack is full")
	}

	s.top++
	s.data[s.top] = v

	return nil
}

func (s *Stack) pop() (float64, error) {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return 0, fmt.Errorf("Stack is empty")
	}

	element := s.data[s.top]
	s.top--

	return element, nil
}

func (s *Stack) getTop() float64 {
	return s.data[s.top]
}

func (s *Stack) show() {
	fmt.Println(s.data[0 : s.top+1])
}

func main() {
	var s *Stack = &Stack{}
	s.create(10)
	s.push(2)
	s.push(7)
	s.show()
	s.pop()
	s.show()
}
```

No código a cima podemos ver a implementação da pilha em código.

Primeiramente temos um struct denominado `Stack` que representa nossa pilha com os atributos:

- `top` - Posição do elemento no topo

- `capacity` - Capacidade total da pilha

- `data` - Dados da pilha

Abaixo temos a função `create` que inicializa a pilha com os valores iniciais.

A função `isEmpty` verifica se a pilha está vazia checando se o valor do `top` é igual a `-1`.

A função `isFull` verifica se a pilha está cheia checando se o valor do `top` é igual a `capacity` (capacidade total da pilha).

A função `push` é a responsável por empilhar um elemento, nela primeiramente verificamos se a pilha já está cheia, caso esteja um erro é retornado, caso não esteja incrementamos 1 no valor de `top` e em seguida adicionamos o elemento na posição do novo valor de `top` e retornamos `nil` para o erro.

Já a função `pop` é a responsável por desempilhar um elemento, nela primeiramente verificamos se a pilha está vazia, caso esteja é retornado 0 e um erro, caso não esteja armazenamos o elemento a ser removido em uma variável em seguida decrementamos 1 do `top`, pois como um elemento foi removido o `top` agora passa a ser o elemento anterior e então retornamos o valor do elemento que foi removido e `nil` para o erro.

Por fim temos as funções `getTop` e `show` que retorna o elemento do topo e exibe a pilha na tela respectivamente.

`OBS: No arquivo stack.go o código está com comentários explicativos`.

### Referências

- https://www.cos.ufrj.br/~rfarias/cos121/pilhas.html.
