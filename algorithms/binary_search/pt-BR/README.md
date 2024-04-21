# Busca binária

Quando queremos fazer uma busca em um array, podemos utilizar o laço de repetição for/foreach para percorrê-lo e então quando encontrarmos o valor o retornamos, certo? Porém a quantidade de interações que o algoritmo terá será muito elevada considerando que contenha um número grande de elementos, este tipo de pesquisa é chamado de pesquisa linear, onde a função de complexidade da pesquisa cresce de forma linear.

Dito isto vejamos um exemplo deste tipo de algoritmo.

```Go
package main

import "fmt"

func main() {
	numbers := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println(search(numbers, 30))
}

func search(arr []int, element int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == element {
			return i
		}
	}

	return -1
}

```

Logo abaixo podemos ver um gráfico que mostra a complexidade do algoritmo crescendo linearmente, dado pela notação O(n).

![001_linear](https://user-images.githubusercontent.com/48635609/91091280-c36ef600-e62c-11ea-9f31-55708fb8196f.PNG)

Para contornar este problema existe um algoritmos denominado de busca binária, a busca binária é um algoritmo de busca, que serve para encontrar um determinado item em um vetor no qual o mesmo necessita estar ordenado.

Seu funcionamento consiste em dividir o vetor pela metade repetidas vezes (Paradigma denominado [divide and conquer](https://github.com/GuilhermehChaves/google-skills/tree/master/algoritmos/divide_and_conquer)) para dividir nosso problema maior em vários sub problemas e assim os resolver isoladamente até que o problema principal seja resolvido por completo, desta forma o número de interações com o array diminui e consequentemente diminui a função de complexidade do algoritmo.

Vejamos outro exemplo agora utilizando busca binária.

```go
package main

import "fmt"

func main() {
	numbers := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println(binarySearch(numbers, 30))
}

func binarySearch(arr []int, element int) int {
	var start, final, mid int
	start = 0
	final = len(arr) - 1

	for start <= final {
		mid = (start + final) / 2
		if arr[mid] == element {
			return mid
		}
		if arr[mid] < element {
			start = mid + 1
		} else {
			final = mid - 1
		}
	}

	return -1
}

```

Logo abaixo podemos ver um outro gráfico demonstrando a função de complexidade deste algoritmo que cresce de forma logarítmica, dado pela notação O(log n).

![002_log](https://user-images.githubusercontent.com/48635609/91092368-4a709e00-e62e-11ea-91b2-ddb17b872586.PNG)

Comparação da quantidade de interações entre a busca linear e a busca binária.

![004_compare](https://user-images.githubusercontent.com/48635609/91093901-73922e00-e630-11ea-98b0-a7370b856144.gif)

Analisando o gif vemos cada etapa que o algoritmo de busca binária realiza, podemos notar que de primeira ele divide o array em duas partes
e verifica se o elemento do meio é o que procuramos, caso seja já o achamos e é retornado sua posição, caso não seja verifica se o elemento ao meio é maior ou menor que o procurado,
sendo menor sabemos que o número que procuramos está na metade à direita do array, logo podemos ignorar a metade à esquerda, em seguida
o array é dividido novamente pela metade e o mesmo processo é realizado até que encontramos o valor desejado, esta técnica reduz a quantidade
de interações que o algoritmo realizará com nosso array como pode ser visto através da comparação feita no gif a cima.

![003_table](https://user-images.githubusercontent.com/48635609/91092411-5eb49b00-e62e-11ea-9466-1d7a393f26e0.PNG)

### Referências

- https://algoritmosempython.com.br/cursos/algoritmos-python/pesquisa-ordenacao/pesquisa-binaria

- https://pt.khanacademy.org/computing/computer-science/algorithms/binary-search/a/binary-search

- https://www.blogcyberini.com/2017/09/busca-binaria.html

### Materiais complementares

- [Como implementar BUSCA BINÁRIA? *Você deveria aprender isso!* | Algoritmos #10](https://www.youtube.com/watch?v=EgLE5HwRy_M) - vídeo

- [Busca Binária](https://www.youtube.com/watch?v=l6pxuyV3mKQ) - vídeo

- [43 - Busca Binária](https://www.youtube.com/watch?v=NxriUYWxoEU) - vídeo
