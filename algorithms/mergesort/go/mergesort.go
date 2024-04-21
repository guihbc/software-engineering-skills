package main

func mergeSort(a []int) []int {
	// Número de elementos
	var num = len(a)

	// Se o número de elementos for 1, retorna o próprio array
	if num == 1 {
		return a
	}

	// Metade do array
	middle := int(num / 2)

	// Cria um array de tamanho "middle" e um de tamanho "num-middle"
	// ou seja um array terá o número de posições de 0 até o meio
	// e um array que terá o número de posições de 0 até quantidade de elementos - meio
	// representando a parte esquerda e dereita do array principal
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)

	// loop que executa a quantidade de vezes do tamanho do array original
	for i := 0; i < num; i++ {
		// Se o i for menor que o meio (parte à esquerda)
		if i < middle {
			// O array que geramos chamado left na posição i, recebe o valor do array naquela posição
			left[i] = a[i]
		} else { // i > middle
			// Caso o i seja maior que o meio (parte à direita)
			// O array que geramos chamado right na posuição i-middle, recebe o valor do array naquela posição
			right[i-middle] = a[i]
		}
	}

	// Ao final teremos dois arrays com os elemendos à esquerda e à direita do nosso array principal

	// Retornamos a função merge passando como argumento a função mergeSort(left) e mergeSort(right)
	// utilizando recursão, o mesmo processo já feito é realizado em cada parte do array
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	// Cria um array com o tamanho do array principal len(left) + len(right)
	result = make([]int, len(left)+len(right))

	i := 0
	// Enquanto o tamnho do array à esquerda for maior que 0 e
	// o tamanho do array a direita for maior que 0
	for len(left) > 0 && len(right) > 0 {
		// Se o primeiro elemento do array referente a parte a esquerda for
		// menor que o primeiro elemento do array referente a parte a direita
		if left[0] < right[0] {
			// result na posição i = left[0] (primeiro elemento do array referente a parte esquerda)
			result[i] = left[0]
			// left = ele mesmo a partir da posição 1 até o final
			// ou seja, dimunui o seu tamanho
			left = left[1:]
		} else { // left[0] > right[0]
			// Caso o primeiro elemento referente a parte à esquerda for maior
			// result na posição i = right[0] (primeiro elemento do array referente a parte direita)
			result[i] = right[0]
			// right = ele mesmo a partir da posição 1 até o final
			// ou seja, diminui seu tamanho
			right = right[1:]
		}
		// Incrementamos 1 no valor de i
		i++
	}

	// Loop que executa a quantidade de vezes do tamnho do array referente a parte à esquerda
	for j := 0; j < len(left); j++ {
		// result na posição i = left na posição j
		result[i] = left[j]
		// Incrementamos 1 no valor de i
		i++
	}
	// Loop que executa a quantidade de vezes do tamnho do array referente a parte à direita
	for j := 0; j < len(right); j++ {
		// result na posição i = right na posição j
		result[i] = right[j]
		// Incrementamos 1 no valor de i
		i++
	}

	return
}
