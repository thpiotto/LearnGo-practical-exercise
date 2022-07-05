package main

/*
 - Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis.
   Utilize os seguintes identificadores e tipos para estas variáveis:
 	1. Identificador "x" e deverá ser tipo int
 	2. Identificador "y" e deverá ser tipo string
	3. Identificador "z" e deverá ser tipo bool

 - Na função main:
 	1. Demonstre os valores de cada identidicador
 	2. O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?

 Resposta: Valor zero!
*/

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("x: %v, y: %v, z: %v", x, y, z)
}
