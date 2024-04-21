package main

import "errors"

type Queue struct { // struct da fila
	capacity int
	data     []float64
	first    int
	last     int
	nItems   int
}

func (q *Queue) create(c int) { // função que cria uma fila
	q.capacity = c              // capacidade da fila
	q.data = make([]float64, c) // dados da fila iniciado vazio inicialmente
	q.first = 0                 // posição do primeiro elemento
	q.last = -1                 // posição do último elemento
	q.nItems = 0                // número de elementos
}

func (q *Queue) insert(v float64) error { // função que insere um elemento à fila
	if q.last == q.capacity-1 {
		return errors.New("queue overflow")
	}

	q.last++           // incrementamos mais 1 no dado referente ao índice da última posição da fila
	q.data[q.last] = v // inserimos o elemento na última posição da fila
	q.nItems++         // incrementamos mais 1 no número de elementos da fila
	return nil
}

func (q *Queue) remove() { // função que remove o elemento da fila
	q.first++ // incrementamos mais 1 no dado referente ao índice da primeira posição da fila

	if q.first == q.capacity { // se o índice do primeiro elemento for o tamanho da capacidade
		q.first = 0 // o índice do primeiro elemento agora é 0
	}

	q.nItems-- // decrementamos 1 no número de elementos da fila
}
