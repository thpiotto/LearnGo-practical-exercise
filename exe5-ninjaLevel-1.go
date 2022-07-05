package main

/*
 - Utilizando a solução do exercicio anterior:
	1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com um identificador "y".
	   O tipo dessa variável deve ser o tipo subjacente do tipo que você criou no exercicio anterior.
 	2. Na função main:
 		1. Isto já deve estar feito:
			1. Demonstre o valor da variável "x"
			2. Demonstre o tipo da variável "x"
 			3. Atribua à variavel "x" utilizando opoerado "="
			4. Demonstre o valor da variável "x"
		2. Agora faço isso:
			1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e,
			   utilizando o operador "=", atribua o valor de "x" a "y"
			2. Demonstre o valor da variável "y"
 			3. Demonstre o tipo da variável "y"

*/

import "fmt"

type gopher int

var x gopher
var y int

func main() {
	fmt.Printf("x: %v\n", x)
	fmt.Printf("Tipo de x: %T\n", x)

	x = 42
	fmt.Printf("x: %v\n", x)

	y = int(x)
	fmt.Printf("y: %v\n", y)
	fmt.Printf("Tipo de y: %T\n", y)
}
