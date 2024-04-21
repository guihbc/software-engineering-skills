# Mergesort

Assim como o Quicksort o Mergesort é um algoritmo de ordenação que utiliza do paradigma [divide and conquer](https://github.com/guihcodes/software-engineering-skills/tree/develop/algorithms/divide_and_conquer/pt-BR),
que consiste em ramificar a entrada várias vezes para assim separar os problemas em vários outros menores.

### Merge

Este algoritmo divide a lista em duas partes e utiliza recursão para executar a função em ambas elas ou seja divide cada parte já divida em duas
até que não tenha mais jeito e reste várias listas com apenas uma posição referente a cada elemento da lista original. Em seguida compara os valores 
entre um e outro e verifica qual é menor, o menor valor é colocado à esquerda e em seguida realiza o merge destas partes, agora juntando a lista novamente
formando listas maiores e realizando o mesmo processo de comparação e merge em ambas as listas formadas até o fim onde a lista estará ordenada por completo.

Observando a imagem podemos entender melhor como este algoritmo funciona.

![002_mergesort](https://user-images.githubusercontent.com/48635609/91216255-0643d300-e6ec-11ea-82ba-2b8af5fe5b42.gif)

Com o gif podemos entender mais claramente o processo, a lista começa "de cima para baixo" realizando a separação dos elementos até que tenha
várias listas com apenas uma posição referente a cada elemento da lista principal, ao chegar ao final a lista vem realizando o processo "de baixo para cima", desta vez 
realizando as comparações dos valores entre as várias partes da lista as colocando em ordem e em seguida realiza o merge, ordenando as sub listas até que chegue 
ao seu tamanho original e a ordenação seja realizada por completo.

![001_mergesort](https://user-images.githubusercontent.com/48635609/91215813-638b5480-e6eb-11ea-82ca-fa414d5a5d2f.png)

Abaixo podemos ver a implementação do algoritmo `mergesort` em Go:

```go
package main

import (
	"fmt"
)

func main() {
	array := []int{5, 3, 2, 1, 4}
	fmt.Println("\n--- Desordenado --- \n\n", array)
	fmt.Println("\n--- Ordenado ---\n\n", mergeSort(array), "\n")
}

func mergeSort(a []int) []int {
	var num = len(a)

	if num == 1 {
		return a
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)

	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = a[i]
		} else {
			right[i-middle] = a[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

```

`OBS: No arquivo "mergesort.go" o código acima está com comentários explicativos em cada linha`

### Referências

- https://www.golangprograms.com/golang-program-for-implementation-of-mergesort.html

- https://www.geeksforgeeks.org/merge-sort/

### Materiais complementares

- [MERGE SORT | Algoritmos #7](https://www.youtube.com/watch?v=5prE6Mz8Vh0) - vídeo

- [Algoritmo de ordenação - MergeSort Video Aula](https://www.youtube.com/watch?v=PKCMMSXQyJE) - vídeo
