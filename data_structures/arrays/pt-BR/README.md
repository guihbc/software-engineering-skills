# Arrays

Para entendermos o que é um array primeiro precisamos entender, o que é uma variável?

Em programação uma `variável` é um espaço alocado na memória RAM do computador que utilizamos para armazenar um valor ou expressão durante a execução de um programa.

Exemplo:

```Go
variavel := 10
nome := "Guilherme"
```

No exemplo escrito em `Go` duas variáveis são declaradas ou seja alocamos dois espaços na memória RAM que armazenam os valores 10 e Guilherme respectivamente e podemos utilizalos da melhor maneira possível.

Entendendo o que é uma variável podemos então entender, o que é um array?

Em programação um `array` ou arranjo, chamado também de variável composta é um conjunto de elementos de mesmo tipo e é considerado uma estrutura de dados onde cada elemento pode ser identificado com o seu índice nesta estrutura.

Exemplo:

```Go
numeros := [5]int{1,2,3,4,5}
num := numeros[1]
```

No exemplo a cima escrito em `Go` temos um array do tipo `int`(números inteiros) denominado `numeros` onde temos um conjunto de elementos, para podermos acessar um elemento específico em `números`, basta especificarmos seu índice no array como podemos ver na variável `num` que armazena o valor de `numeros` na posição 1 ou seja o elemento `2`, em geral na maioria das linguagens de programação os índices dos arrays começam na posição 0, portanto índice 1 é referente ao elemento `2`.
