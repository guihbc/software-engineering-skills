package main

import (
	"math/rand"
)

func quicksort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	// Indice do primeiro e último elemento
	var start, end int = 0, len(a) - 1

	// Setando número aleatório ao pivot
	pivot := rand.Int() % len(a)

	// Realizando a troca do último elemento com o pivot
	a[pivot], a[end] = a[end], a[pivot]

	for i, _ := range a {
		// Se o elemento na posição i for menor que o elemento na posição final
		if a[i] < a[end] {
			// Troca as posição do primeiro elemento com o elemento na posição i
			a[start], a[i] = a[i], a[start]
			// Somando + 1 no start (neste caso a variável start representa o i no exemplo)
			start++
		}
	}

	// Realiza a troca do pivô com a próxima posição a ser trocada
	a[start], a[end] = a[end], a[start]
	// Chama a função recursivamente, pegando da posição 0 até o start
	quicksort(a[:start])
	// Chama a função recursivamente, começando do start + 1 até o final
	quicksort(a[start+1:])

	return a
}
