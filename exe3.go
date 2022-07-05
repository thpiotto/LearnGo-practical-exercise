package main

/*
 - Utilizando a solução do exercício anterior:
 1. Em package-level scope, atribua os seguintes valores às variáveis:
	1. para "x", atribua 42
 	2. para "y", atribua James Bond
	3. para "z", atribua true

 - Na função main:
 	1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável.
 	   Faça essa atribuição de tipo string a uma variável de indentificador "s" utilizando  operador curto de declaração.
 	2. Demonstre a variável "s"
*/

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("x: %v, y: %v, z: %v", x, y, z)
	fmt.Print(s)
}
