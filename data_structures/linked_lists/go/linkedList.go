package main

import (
	"fmt"
)

type Node struct { // struct que representa um nó
	element interface{} // dado a ser inserido no nó
	next    *Node       // endereço do próximo nó
}

func (n *Node) isEmpty() bool { // função que verifica se o nó está vazio
	return n == nil
}

type LinkedList struct { // struct que representa uma lista
	first *Node // endereço do primeiro nó
}

func NewLikedList() *LinkedList {
	return &LinkedList{
		first: nil,
	}
}

func (l *LinkedList) insert(e interface{}) { // função que insere um novo dado na lista
	var node *Node = &Node{} // cria um novo nó
	node.element = e         // o elemento do nó é igual ao elemento passado
	node.next = l.first      // o próximo nó agora é igual ao primeiro
	l.first = node           // o primeiro nó da lista agora é igual ao nó criado a cima
}

func (l *LinkedList) isEmpty() bool { // função que verifica se a lista está vazia
	return l.first == nil
}

func (l *LinkedList) search(e interface{}) *Node { // função que procura um elemento na lista
	for n := l.first; n != nil; n = n.next {
		if n.element == e { // se o elemento no nó for igual ao passado
			return n // retorna o
		}
	}

	return nil // retorna nil
}

func (l *LinkedList) remove(e interface{}) int { // função que remove um elemento do nó
	var previous *Node = nil // nó anterior
	var p *Node = l.first    // primeiro nó da lista

	for p != nil && p.element != e {
		previous = p // nó anterior é igual ao primeiro
		p = p.next   // primeiro nó é igual ao próximo
	}

	if p == nil { // se o p for nil
		return -1 // o elemento não foi encontrado
	}

	if previous == nil { // se o anterior for nil
		l.first = p.next // primeiro nó da lista é igual ao nó depois do primeiro
	} else { // se não
		previous.next = p.next // o nó após o anterior é igual ao nó após o primeiro
	}

	return 0
}

func (l *LinkedList) free() { // função que esvazia a lista
	for l.first != nil { // enquanto o primeiro nó for diferente de nil
		var temp *Node = l.first.next // nó depois do primeiro
		l.first = nil                 // primeiro nó da lista é igual a nil
		l.first = temp                // primeiro nó da lista é igual ao nó seguinte
	}
}

func (l *LinkedList) show() { // função que exibe a lista na tela
	fmt.Print("[ ")
	print(l.first)
	fmt.Println("]")
}

func print(n *Node) { // função recursiva que imprime os elementos da lista
	if !n.isEmpty() { // verifica se o nó é diferente de vazio
		print(n.next)        // chama a si mesma passando o próximo nó
		fmt.Print(n.element) //printa o elemento na tela
		fmt.Print(" ")       // printa espaço vazios
	}
}
