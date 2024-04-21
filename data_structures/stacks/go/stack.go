package main

import "fmt"

// struct representando a pilha
type Stack struct {
	top      int       // posição do top
	capacity int       // tamanho da capacidade
	data     []float64 // dados da pilha
	length   int
}

func (s *Stack) create(c int) { // função que cria uma pilha com os valores iniciais
	s.top = -1
	s.capacity = c
	s.data = make([]float64, c)
}

func (s *Stack) isEmpty() bool { // função que verifica se a pilha está vazia
	return s.top == -1
}

func (s *Stack) isFull() bool { // função que verifica se a pilha está cheia
	return s.top == s.capacity-1
}

func (s *Stack) push(v float64) error { // função que empilha um elemento
	if s.isFull() { // verifica se a pilha está cheia
		return fmt.Errorf("stack overflow") // retorna um erro
	}

	s.top++           // incrementa 1 no atributo referente a posição do topo
	s.data[s.top] = v // adiciona o elemento na posição do novo topo
	s.length++

	return nil // retorna nil para o erro
}

func (s *Stack) pop() (float64, error) { // função que desempilha um elemento
	if s.isEmpty() { // verifica se a pilha está vazia
		fmt.Println("Stack is empty")
		return 0, fmt.Errorf("Stack is empty") // retorna 0 e um erro
	}

	element := s.data[s.top] // elemento que será removido
	s.top--                  // decrementa 1 no atributo referente a posição do topo
	s.length--

	return element, nil // retorna o elemento removido e nil para o erro
}

func (s *Stack) getTop() float64 { // função que retona o elemento no topo
	return s.data[s.top]
}

func (s *Stack) show() { // função que exibe a pilha na tela
	fmt.Println(s.data[0 : s.top+1])
}

func main() {
	var s *Stack = &Stack{} // define uma pilha chamada s
	s.create(10)            // cria uma pilha com os valores iniciais
	s.push(2)               // empilha o valor 2
	s.push(7)               // empilha o valor 7
	s.show()                // exibe a pilha na tela
	s.pop()                 // desempilha um valor
	s.show()                // exibe a pilha na tela
}
