# Quicksort

O Quicksort é um algoritmo de ordenação que utiliza o conceito de ordenação por troca, ordenação é quando colocamos algo em ordem podendo ser crescente, decrescente, alfabética, etc.

Este algoritmo utiliza do paradigma [divide and conquer](https://github.com/GuilhermehChaves/google-skills/tree/master/algoritmos/divide_and_conquer), que consiste em ramificar a entrada várias vezes para assim separar o problemas em vários outros menores.

#### Partition

No partition definimos uma variável chamada de pivô, e sem seguida colocamos todos os elementos menores do que o pivô a sua esquerda o pivô pode ser o último ou o primeiro elemento e pode ser escolhido de maneira aleatória.

Se considerarmos a sequência: [5,3,2,1,4]

O pivô será o número 4, percorremos todo o array e quando um valor for menor que o pivô o colocaremos a sua esquerda, para não nos perdermos armazenamos em uma variável `i` os elementos já trocados, no final do percurso pegamos nosso pivô e o substituímos pela posição de `i`, assim temos à esquerda todos os valores menores que o pivô e à direita os valores maiores, mas não quer dizer que ainda estão ordenados corretamente, logo em seguida usamos recursão para chamarmos a função novamente para ordenar a parte à esquerda e à direita do array e ordená las, cada vez que a função é chamada teremos um novo pivô.

Vejamos passo a passo como ficaria esta ordenação: [5,3,2,1,4]

![quicksort](https://user-images.githubusercontent.com/48635609/91211629-7a2ead00-e6e5-11ea-84f4-d632ddc35841.png)

Vejamos agora um exemplo em código:

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := []int{5, 3, 2, 1, 4}

	fmt.Println("\n --- Desordenado ---\n\n", array, "\n")
	quicksort(array)
	fmt.Println("\n --- Ordenado ---\n\n", array, "\n")
}

func quicksort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	var start, end int = 0, len(a) - 1
	pivot := rand.Int() % len(a)

	a[pivot], a[end] = a[end], a[pivot]


	for i, _ := range a {
		if a[i] < a[end] {
			a[start], a[i] = a[i], a[start]
			start++
		}
    }
    
	a[start], a[end] = a[end], a[start]

	quicksort(a[:start])
	quicksort(a[start+1:])

	return a
}

```

No início da função quicksort o algoritmo verifica se o tamanho do array é menor ou igual a 1, caso seja retorna o proprio array,
em seguida declaramos duas variáveis start e end, referentes ao indice do primeiro elemento e do último,
logo em seguida declaramos o pivô como um valor aleatório do nosso array e o colocamos na última posição.

Ao final deste processo entramos na parte onde percorremos nosso array comparando elementos na posição `i` com o nosso pivô e
caso o nosso elemento na posição `i` for menor realizamos a troca do mesmo pela posição `start`, podemos notar também que 
estamo utilzando a variável `start` para armazenar quais elementos já foram trocados utilizando `start++`, assim sempre que
um elemento trocar de posição saberemos qual é a posição do próximo elemento a ser trocado.

No final da nossa repetição realizamos a troca do nosso pivô com posição do `start` (Posição que contém qual o próximo elemento a ser trocado)
como dito na explicação do algoritmo.

Por fim chamamos recursivamente a função `quicksort` passando a parte à esquerda `a[:start]` (Retorna o array da posição 0 até a posição de start) do array e
chamamos-a novamente passando a parte à direita `a[start+1:]` (Retorna o array da posição start+1 até o final) do array,
assim divide o array em várias partes e ordena cada parte separadamente até que sua ordenação esteja completa.

### Referências

- https://www.golangprograms.com/golang-program-for-implementation-of-quick-sort.html

- https://blog.pantuza.com/artigos/o-algoritmo-de-ordenacao-quicksort

### Materiais complementares

- [QUICKSORT | Algoritmos #8](https://www.youtube.com/watch?v=wx5juM9bbFo) - vídeo
