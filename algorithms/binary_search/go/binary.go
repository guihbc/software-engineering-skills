package main

func binarySearch(arr []int, element int) int {
	// Variáveis referente ao primeiro indice, último e meio do array
	var start, final, mid int
	start = 0
	final = len(arr) - 1

	// Enquanto o valor inicial for menor ou igual o final
	for start <= final {
		// Calculando o meio
		mid = (start + final) / 2
		// Se o elemento na posição do meio for o procurado o retorna
		if arr[mid] == element {
			return mid
		}
		// Se o elemento do meio for menor que o procurado
		// o indice do começo agora é a parte do array após o meio (a direita)
		if arr[mid] < element {
			start = mid + 1
			//Se o elemento do meio for maior que o elemento que procuramos
			// o indice do final agora é a parte do array antes do meio (a esquerda)
		} else { // arr[mid] > element
			final = mid - 1
		}
	}
	// Caso o valor não seja encontrado retorna -1
	return -1
}
