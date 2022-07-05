package main

/*
 - Crie um tipo. O tipo subjacente deve ser int
 - Crie uma variável para este tipo, com o indetificador "x", utilizando a palavra var.
 - Na função main:
 	1. Demonstre o valor da variável "x"
	2. Demonstre o tipo da variável "x"
 	3. Atribua à variavel "x" utilizando opoerado "="
	4. Demonstre o valor da variável "x"
*/

import "fmt"

type gopher int

var x gopher

func main() {
	fmt.Printf("x: %v\n", x)
	fmt.Printf("Tipo de x: %T\n", x)
	x = 42
	fmt.Printf("x: %v", x)
}
